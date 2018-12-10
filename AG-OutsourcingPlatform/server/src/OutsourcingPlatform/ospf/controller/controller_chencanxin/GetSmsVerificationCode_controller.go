
    package controller_chencanxin
    
    import (
		"OutsourcingPlatform/ospf/model/model_chencanxin"
		"net/http"
	  )
    
    type GetSmsVerificationCodeController struct {
      *model_chencanxin.GetSmsVerificationCodeModel
    }
    
    type ApiParamOfGetSmsVerificationCode struct {
      
    }
    
    type ApiResponseOfGetSmsVerificationCode struct {
      
    }
    
    func (c *GetSmsVerificationCodeController) ToGetSmsVerificationCode(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    
    
        //var apiParam ApiParamOfGetSmsVerificationCode
        //var apiResponse ApiResponseOfGetSmsVerificationCode
    	
    	
      	msg:="GetSmsVerificationCodeController is work!"
      	return msg,nil
      	
    }
    