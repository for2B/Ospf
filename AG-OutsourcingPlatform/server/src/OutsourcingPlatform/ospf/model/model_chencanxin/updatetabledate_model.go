
  
      package model_chencanxin
      
      import (
          "OutsourcingPlatform/ospf/infrastructure"
          "sync"
      )
      
      var onceUpdateTableDateModel sync.Once
      
      type UpdateTableDateModel struct {
        *infrastructure.Infrastructure
      }
      
      var updateTableDateModel *UpdateTableDateModel
      
      func GetUpdateTableDateModel(i *infrastructure.Infrastructure) *UpdateTableDateModel {
        onceUpdateTableDateModel.Do(func() {
          updateTableDateModel = &UpdateTableDateModel{
            Infrastructure: i,
          }
        })
        return updateTableDateModel
      }
      
    