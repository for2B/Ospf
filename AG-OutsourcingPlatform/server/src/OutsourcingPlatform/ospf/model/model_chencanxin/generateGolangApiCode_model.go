package model_chencanxin

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
)


var onceGenerateGolangApiCodeModel sync.Once

type GenerateGolangApiCodeModel struct {
	*infrastructure.Infrastructure
}

var generategolangapicodeModel *GenerateGolangApiCodeModel

func GetGenerateGolangApiCodeModel(i *infrastructure.Infrastructure) *GenerateGolangApiCodeModel {
	onceGenerateGolangApiCodeModel.Do(func() {
		generategolangapicodeModel = &GenerateGolangApiCodeModel{
			Infrastructure: i,
		}
	})
	return generategolangapicodeModel
}

