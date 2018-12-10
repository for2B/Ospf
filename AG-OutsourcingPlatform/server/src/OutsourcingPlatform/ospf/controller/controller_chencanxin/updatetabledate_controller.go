
    package controller_chencanxin
    
    import (
		"OutsourcingPlatform/ospf/model/model_chencanxin"
		"net/http"
	  )
    
    type UpdateTableDateController struct {
      *model_chencanxin.UpdateTableDateModel
    }
    
    type ApiParamOfUpdateTableDate struct {
      
    }
    
    type ApiResponseOfUpdateTableDate struct {
      
    }
    
    func (c *UpdateTableDateController) ToUpdateTableDate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    
    
        //var apiParam ApiParamOfUpdateTableDate
        //var apiResponse ApiResponseOfUpdateTableDate
    	
    	
      	msg:="UpdateTableDateController is work!"
      	return msg,nil
      	
    }
    