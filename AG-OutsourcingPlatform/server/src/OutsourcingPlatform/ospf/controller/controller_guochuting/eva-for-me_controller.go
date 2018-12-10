package controller_guochuting
//对我的评价
import (
	"net/http"
	"OutsourcingPlatform/ospf/model/model_guochuting"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
	"fmt"
)

type GetEvaluationForMeController struct {
	*model_guochuting.GetEvaluationModel
}

type GetEvaluationForMe struct {
	KeyUserName int	`json:"key_user_name"`
	CurrentOage int	`json:"current_page"`
	ShowNums int		`json:"show_nums"`
}

func (g *GetEvaluationForMeController) GetEvaluationForMeController(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != "GET" {
		return nil, httpapi.NewErr(constant.GLOBAL_HTTP_METHOD_ERR, "Invalid request method", nil)
	}
	fmt.Println("GetEvaluationForMeController")
	vars := r.URL.Query();
	key_user_name := vars.Get("key_user_name")
	current_page := vars.Get("current_page")
	show_nums := vars.Get("show_nums")
	return g.GetEvaForMeModel(key_user_name, current_page, show_nums);
}
