package controller_guochuting

import (
	"net/http"
	"OutsourcingPlatform/ospf/model/model_guochuting"
	"fmt"
	"encoding/json"
	"OutsourcingPlatform/internal/httpparse"
)


type PostEvaluateController struct {
	*model_guochuting.PostEvaluateModel
}


func (g *PostEvaluateController) PostEvaluateController(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fmt.Println("EvaluateController")
	if r.Method == "POST"{
		getbody ,err := httpparse.GetBody(r)
		//fmt.Println("getbody=", getbody)
		var eva model_guochuting.PostEvaluate
		err = json.Unmarshal([]byte(getbody), &eva)

		if err != nil{
			//fmt.Println("err=", err)
			return nil, err
		}
		//fmt.Println("eva=", eva)
		return g.Postevaluatemodel(eva.EstimatorId, eva.EstimatoredId, eva.EvaluateTime, eva.EvaluateDetail, eva.ProjectSsociatedId, eva.IsPrivate, g.DB)
	}

	return nil, nil
}
