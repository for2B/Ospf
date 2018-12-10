package ospf

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/infrastructure"
	"OutsourcingPlatform/ospf/model"
	"OutsourcingPlatform/ospf/infrastructure/option"
	"OutsourcingPlatform/ospf/view"
	"OutsourcingPlatform/utils"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

//remember to delete me
type OSPF struct {
	httpListener net.Listener
	sync.RWMutex
	IsLoading int32		//isLoading 0->loading 1->load finish
	startTime time.Time
	exitChan  chan struct{}
	waitGroup utils.WaitGroupWrapper
	infrastructure *infrastructure.Infrastructure
}

func New(opts *option.Options) *OSPF {
	p := &OSPF{
		startTime:      time.Now(),
		exitChan:       make(chan struct{}),
		infrastructure: infrastructure.NewInfrastructure(opts),
	}
	return p
}

func (p *OSPF) Main() {
	opts := p.infrastructure.GetOpts()
	mm := model.GetModel(p.infrastructure)
	mm.InitAllTable()
	 //mm.InsertDataToAllTable()  // 暂时不要删除，数据库模拟数据插入
	httpListener, err := net.Listen("tcp", ":"+opts.ServerPort)
	if err != nil {
		p.infrastructure.Logger.Fatal(fmt.Sprintf("listen (%s) failed - %s", opts.ServerPort, err))
		os.Exit(1)
	}
	p.Lock()
	p.httpListener = httpListener
	p.Unlock()
	view.Init(opts.FrontDir)
	server := newHTTPServer(
		&context{
			ospf: p,
		})
	p.waitGroup.Wrap(func() {
		httpapi.Serve(httpListener, server, p.infrastructure.Logger)
	})
}

func (p *OSPF) Exit() {
	if p.httpListener != nil {
		p.httpListener.Close()
	}
	p.WriteCurData()
	//TODO: ADD MORE
	close(p.exitChan)
	p.waitGroup.Wait()
	p.infrastructure.Logger.Sync()
	if p.infrastructure.DB != nil {
		p.infrastructure.DB.Close()
	}
}

func (p *OSPF) LoadLastData() {
	//todo:load last data

}

func (p *OSPF) WriteCurData() {
	//todo:write data before exit
}
