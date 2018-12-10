package model_guozhenzhen

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
)

type GetUserStateModel struct {
	*infrastructure.Infrastructure
	UserState int64
}

var onceGetUserStateModel sync.Once
var getuserstate *GetUserStateModel

func GetUserstateModel(i *infrastructure.Infrastructure) *GetUserStateModel {
	onceGetUserStateModel.Do(func() {
		getuserstate = &GetUserStateModel{
			Infrastructure: i,
		}
	})
	return getuserstate
}

func (m *GetUserStateModel)GetuserState(user string)(){

}
