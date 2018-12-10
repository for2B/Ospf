
  
      package model_chencanxin
      
      import (
          "OutsourcingPlatform/ospf/infrastructure"
          "sync"
      )
      
      var onceGetSmsVerificationCodeModel sync.Once
      
      type GetSmsVerificationCodeModel struct {
        *infrastructure.Infrastructure
      }
      
      var getSmsVerificationCodeModel *GetSmsVerificationCodeModel
      
      func GetGetSmsVerificationCodeModel(i *infrastructure.Infrastructure) *GetSmsVerificationCodeModel {
        onceGetSmsVerificationCodeModel.Do(func() {
          getSmsVerificationCodeModel = &GetSmsVerificationCodeModel{
            Infrastructure: i,
          }
        })
        return getSmsVerificationCodeModel
      }
      
    