

package model_chencanxin

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
)

var onceResetPasswordModel sync.Once

type ResetPasswordModel struct {
	*infrastructure.Infrastructure
}

var resetpasswordModel *ResetPasswordModel

func GetResetPasswordModel(i *infrastructure.Infrastructure) *ResetPasswordModel {
	onceResetPasswordModel.Do(func() {
		resetpasswordModel = &ResetPasswordModel{
			Infrastructure: i,
		}
	})
	return resetpasswordModel
}
      
    