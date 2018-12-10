
  
      package model_chencanxin
      
      import (
          "OutsourcingPlatform/ospf/infrastructure"
          "sync"
      )
      
      var onceValidSmsVerificationCodeModel sync.Once
      
      type ValidSmsVerificationCodeModel struct {
        *infrastructure.Infrastructure
      }
      
      var validSmsVerificationCodeModel *ValidSmsVerificationCodeModel
      
      func GetValidSmsVerificationCodeModel(i *infrastructure.Infrastructure) *ValidSmsVerificationCodeModel {
        onceValidSmsVerificationCodeModel.Do(func() {
          validSmsVerificationCodeModel = &ValidSmsVerificationCodeModel{
            Infrastructure: i,
          }
        })
        return validSmsVerificationCodeModel
      }
      
    