package projectOperateModel

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
)

var onceProjectOperate sync.Once

type ProjectOperateModel struct {
	*infrastructure.Infrastructure
}

var projectOperateModel *ProjectOperateModel

func ProduceProjectOperateModel(i *infrastructure.Infrastructure) *ProjectOperateModel  {
	onceProjectOperate.Do(func() {
		projectOperateModel = &ProjectOperateModel{
			Infrastructure:i,
		}
	})
	return projectOperateModel
}