
  
      package model_chencanxin
      
      import (
          "OutsourcingPlatform/ospf/infrastructure"
          "sync"
      )
      
      var onceSelectTableDateModel sync.Once
      
      type SelectTableDateModel struct {
        *infrastructure.Infrastructure
      }
      
      var selectTableDateModel *SelectTableDateModel
      
      func GetSelectTableDateModel(i *infrastructure.Infrastructure) *SelectTableDateModel {
        onceSelectTableDateModel.Do(func() {
          selectTableDateModel = &SelectTableDateModel{
            Infrastructure: i,
          }
        })
        return selectTableDateModel
      }
      
    