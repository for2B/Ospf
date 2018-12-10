package model_chencanxin

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
)

var onceWcAuthModel sync.Once

type WcAuthModel struct {
	*infrastructure.Infrastructure
}

var wcAuthModel *WcAuthModel

func GetWcAuthModel(i *infrastructure.Infrastructure) *WcAuthModel {
	onceWcAuthModel.Do(func() {
		wcAuthModel = &WcAuthModel{
			Infrastructure: i,
		}
	})
	return wcAuthModel
}


