package controller_guochuting


//邀请的评价
import (
	"net/http"
	"OutsourcingPlatform/ospf/model/model_guochuting"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"

	"strconv"
)

type GetEvaluationController struct {
	*model_guochuting.GetEvaluationModel
}

type GetEvaluationInv struct {
	UserId int		`json:"user_id"`
	KeyUserName string	`json:"key_user_name"`
	CurrentPage int	`json:"current_page"`
	ShowNums int	`json:"show_nums"`
}

func (g *GetEvaluationController) GetEvaluationController(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != "GET" {
		return nil, httpapi.NewErr(constant.GLOBAL_HTTP_METHOD_ERR, "Invalid request method", nil)
	}


	//ses := r.Context().Value("session");
	//fmt.Println("ses=", ses)
	//fmt.Println("evaluation_controller")
	body := r.Form
	//fmt.Println("r.Form=", body)
	name := body.Get("key_user_name")
	currentPage := body.Get("current_page")
	showNums := body.Get("show_nums")
	current_age , err:= strconv.Atoi(currentPage)
	if err != nil{
		return nil, httpapi.Err{Code: constant.GLOBAL_PARAM_ERR, Msg: "row param invalid: must be number"}
	}
	show_nums , err:= strconv.Atoi(showNums)
	if err != nil{
		return nil, httpapi.Err{Code: constant.GLOBAL_PARAM_ERR, Msg: "row param invalid: must be number"}
	}
	//fmt.Println("num", current_age, show_nums)
	return g.GetEvaModel(name, current_age, show_nums, g.DB)

}
