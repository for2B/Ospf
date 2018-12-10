
    package controller_chencanxin
    
    import (
		"OutsourcingPlatform/ospf/model/model_chencanxin"
		"net/http"
	  )
    
    type ValidSmsVerificationCodeController struct {
      *model_chencanxin.ValidSmsVerificationCodeModel
    }
    
    type ApiParamOfValidSmsVerificationCode struct {
      
    }
    
    type ApiResponseOfValidSmsVerificationCode struct {
      
    }
    
    func (c *ValidSmsVerificationCodeController) ToValidSmsVerificationCode(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    
    
        //var apiParam ApiParamOfValidSmsVerificationCode
        //var apiResponse ApiResponseOfValidSmsVerificationCode
    	
    	
      	msg:="ValidSmsVerificationCodeController is work!"
      	return msg,nil
      	
    }
    