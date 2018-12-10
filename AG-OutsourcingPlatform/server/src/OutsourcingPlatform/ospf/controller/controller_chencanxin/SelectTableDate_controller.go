
    package controller_chencanxin
    
    import (
		"OutsourcingPlatform/ospf/model/model_chencanxin"
		"net/http"
	  )
    
    type SelectTableDateController struct {
      *model_chencanxin.SelectTableDateModel
    }
    
    type ApiParamOfSelectTableDate struct {
      
    }
    
    type ApiResponseOfSelectTableDate struct {
      
    }
    
    func (c *SelectTableDateController) ToSelectTableDate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    
    
        //var apiParam ApiParamOfSelectTableDate
        //var apiResponse ApiResponseOfSelectTableDate
    	
    	
      	msg:="SelectTableDateController is work!"
      	return msg,nil
      	
    }
    