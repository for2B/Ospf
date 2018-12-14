package main

import (
	"fmt"
	"log"
	"syscall"
	"github.com/judwhite/go-svc/svc"

	"OutsourcingPlatform/ospf/infrastructure/option"
	"OutsourcingPlatform/utils"
	"flag"
	"sync/atomic"
	"OutsourcingPlatform/ospf"
)

type program struct {
	Ospf *ospf.OSPF
}

var BUILDVER string
var BUILDTYPE string

func (p *program) Init(env svc.Environment) error {
	return nil
}

func (p *program) Start() error {

	//load options from terminal
	proDir, err := utils.GetProDir()
	fmt.Println(proDir)
	if err != nil {
		log.Panicln("get prodir err: " + err.Error())
	}
	fmt.Println("----------- Now serve is start -------------", )

	opts := option.NewOpts(proDir, BUILDTYPE)
	fmt.Println(fmt.Sprintf("dev server address : http://localhost:%s" ,opts.ServerPort ))

	flag.StringVar(&opts.BuildType, "BuildType", opts.BuildType, "BuildType")
	flag.StringVar(&opts.BuildVer, "BuildVer", opts.BuildVer, "BuildVer")
	flag.StringVar(&opts.AppName, "AppName", opts.AppName, "AppName")
	flag.StringVar(&opts.FrontDir, "FrontDir", opts.FrontDir, "FrontDir")
	flag.StringVar(&opts.DBDriver, "DBDriver", opts.DBDriver, "DBDriver")
	flag.StringVar(&opts.DBHost, "DBHost", opts.DBHost, "DBHost")
	flag.StringVar(&opts.DBPort, "DBPort", opts.DBPort, "DBPort")
	flag.StringVar(&opts.DBUser, "DBUser", opts.DBUser, "DBUser")
	flag.StringVar(&opts.DBPassword, "DBPassword", opts.DBPassword, "DBPassword")
	flag.StringVar(&opts.DBName, "DBName", opts.DBName, "DBName")
	flag.StringVar(&opts.DBTestName, "DBTestName", opts.DBTestName, "DBTestName")
	flag.StringVar(&opts.RedisHost, "RedisHost", opts.RedisHost, "RedisHost")
	flag.StringVar(&opts.RedisPort, "RedisPort", opts.RedisPort, "RedisPort")
	flag.StringVar(&opts.RedisPwd, "RedisPwd", opts.RedisPwd, "RedisPwd")
	flag.IntVar(&opts.RedisDB, "RedisDB", opts.RedisDB, "RedisDB")
	flag.StringVar(&opts.ServerPort, "ServerPort", opts.ServerPort, "ServerPort")
	flag.StringVar(&opts.ConsoleOutPutPath, "OutPutPath", opts.ConsoleOutPutPath, "OutPutPath")
	flag.StringVar(&opts.JsonOutPutPath, "JsonOutPutPath", opts.JsonOutPutPath, "JsonOutPutPath")
	flag.StringVar(&opts.ErrOutPutPath, "ErrOutPutPath", opts.ErrOutPutPath, "ErrOutPutPath")
	flag.IntVar(&opts.LoggerBufferCap, "LoggerBuffer", opts.LoggerBufferCap, "LoggerBuffer")
	flag.Int64Var(&opts.LoggerBufDuration, "LoggerBufDuration", opts.LoggerBufDuration, "LoggerBufDuration")
	flag.StringVar(&opts.SessionKey, "SessionKey", opts.SessionKey, "SessionKey")
	flag.IntVar(&opts.MaxLifeTime, "MaxLifeTime", opts.MaxLifeTime, "MaxLifeTime")
	flag.StringVar(&opts.Path, "Path", opts.Path, "Path")
	flag.BoolVar(&opts.HTTPOnly, "HTTPOnly", opts.HTTPOnly, "HTTPOnly")
	flag.IntVar(&opts.MaxAge, "MaxAge", opts.MaxAge, "MaxAge")
	flag.Parse()
	if opts.BuildType!=""||opts.BuildVer!=""{
		fmt.Println("current build type:" + opts.BuildType)
		fmt.Println("current build ver:" + opts.BuildVer)
	}

	p.Ospf = ospf.New(opts)

	//init system
	atomic.StoreInt32(&p.Ospf.IsLoading, 0)
	p.Ospf.LoadLastData()
	atomic.StoreInt32(&p.Ospf.IsLoading, 1)

	p.Ospf.Main()
	return nil
}

func (p *program) Stop() error {
	if p != nil {
		p.Ospf.Exit()
	}
	return nil
}

func main() {
	proDir, err := utils.GetProDir()
	if err != nil {
		log.Panicln("get prodir err: " + err.Error())
	}
	defer utils.HandlePanic(proDir)

	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}
