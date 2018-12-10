package model_guochuting

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
	"database/sql"
	"strconv"
	"math"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

var onceGetEvaluationModel sync.Once

type GetEvaluationModel struct {
	*infrastructure.Infrastructure
}

type PostEvaluate struct {
	EvaluateId int			    	`json:"evaluate_id"`
	EstimatorId int    				 `json:"estimator_id"`
	EstimatoredId int  		 		`json:"estimatored_id"`
	EvaluateTime string 			`json:"evaluate_time"`
	EvaluateDetail string			`json:"evaluate_detail"`
	ProjectSsociatedId  int 		 `json:"project_ssociated_id"`
	IsPrivate bool					`json:"is_private"`
}

type Evaluate struct {
	HeadImg string                  `json:"head_img"`
	Estimator string   				 `json:"estimator"`
	EvaluateTime string 			`json:"evaluate_time"`
	EvaluateDetail string			`json:"evaluate_detail"`
	ProjectSsociated  string 		 `json:"project_ssociated"`
	IsPrivate bool					`json:"is_private"`
}
type Data struct {
	TotalPage string `json:"totalPage"`
	UserArray []Evaluate `json:"userArray"`
}

var getEvaluationModel *GetEvaluationModel

func EvaluationModel(i *infrastructure.Infrastructure) *GetEvaluationModel {
	onceGetEvaluationModel.Do(func() {
		getEvaluationModel = &GetEvaluationModel{
			Infrastructure: i,
		}
	})
	return getEvaluationModel
}

func (m *GetEvaluationModel) GetEvaModel(key_user_name string, current_page, show_nums int, db *sql.DB) (interface{}, error){
	//验证参数合法性
	//fmt.Println("evaluation.go")
	sqlstatement := `SELECT users.head_img, users.user_name, 
		evaluate.evaluate_time, evaluate_detail, project.project_name, evaluate.is_privacy,
		(
			SELECT count (*)
			FROM evaluate
			WHERE  evaluate.estimator_id = users.user_id 
			AND project.project_id = evaluate.project_ssociated_id
			AND evaluate.is_privacy = false
			AND users.user_name like $1
		)
		FROM users, project, evaluate
		WHERE  evaluate.estimator_id = users.user_id 
		AND project.project_id = evaluate.project_ssociated_id
		AND evaluate.is_privacy = false
		AND users.user_name like $1
		ORDER BY evaluate_time DESC
		LIMIT $3 offset $2;
	`
	rows, err := db.Query(sqlstatement, key_user_name, current_page, show_nums)
	if err != nil{
		//fmt.Println("err != nil")
		return nil, httpapi.NewErr(constant.GLOBAL_PARAM_ERR, "db.Query func failed", err)

	}
	//fmt.Println("row=", rows)
	//defer rows.Close()
	//var row *sql.Rows

	count := 0
	var evaluation Data
	//var eva Evaluate
	for rows.Next(){
		var aeva Evaluate
		rows.Scan(
			&aeva.HeadImg,
			&aeva.Estimator,
			&aeva.EvaluateTime,
			&aeva.EvaluateDetail,
			&aeva.ProjectSsociated,
			&aeva.IsPrivate,
			&count,
		)
		evaluation.UserArray = append(evaluation.UserArray, aeva)
	}
	evaluation.TotalPage = strconv.Itoa(int(math.Ceil(float64(count)/float64(show_nums))))
	int_totalPage, err := strconv.Atoi(evaluation.TotalPage)
	if current_page > int_totalPage {
		return nil, httpapi.Err{Code: constant.GLOBAL_PARAM_ERR, Msg: "currentPage param invalid: can't larger than totalPage"}
	}
	//fmt.Println("totalpage=", int_totalPage)

	//fmt.Println("eva=", evaluation)
	return evaluation, nil;
}



func validParam(currentPage, row int) (int_currentPage, int_row int, err error) {
	// 检查参数合法性，int_currentPage最小值为1 ，show_nums最大值20,
	if int_currentPage <= 0 {
		err = httpapi.Err{Code: constant.GLOBAL_PARAM_ERR, Msg: "currentPage param invalid: can't smaller than 0 or equal 0"}
		return
	}
	if int_row <= 0 {
		err = httpapi.Err{Code: constant.GLOBAL_PARAM_ERR, Msg: "row param invalid: can't smaller than 0 or equal 0"}
		return
	}
	if int_row > 20 {
		err = httpapi.Err{Code: constant.GLOBAL_PARAM_ERR, Msg: "row param invalid: can't bigger than 20"}
		return
	}
	return
	return
}