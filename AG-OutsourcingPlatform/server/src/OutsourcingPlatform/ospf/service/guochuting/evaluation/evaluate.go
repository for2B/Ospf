package evaluation

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"database/sql"
	"fmt"
)

type Evaluates struct {
	HeadImg string  					`json:"head_img"`
	EstimatorName int64	 				`json:"estimator_name"`
	EvaluateTime int64	 				`json:"evaluate_time"`
	EvaluateDetail string				`json:"evaluate_detail"`
	ProjectSsociatedName string			`json:"project_ssociated_name"`
	*infrastructure.Infrastructure 		`json:"-"`
}

func Evaluate(i *infrastructure.Infrastructure) *Evaluates {
	evaluate := &Evaluates{
		Infrastructure:i,
	}
	return evaluate
}
//type PostEvaluates struct {
//	EvaulateId int		`json:"evaluate_id"`
//	EstimatorId int     `json:"estimator_id"`
//	EstimatoredId int   `json:"estimatored_id"`
//	EvaluateTime string `json:"evaluate_time"`
//	EvaluateDetail string	`json:"evaluate_detail"`
//	ProjectSsociatedId  int  `json:"project_ssociated_id"`
//	IsPrivate bool	`json:"is_private"`
//}

//根据请求参数返回平台项目列表
func PostEvaluate(EvaluateId, EstimatorId, EstimatoredId int, EvaluateTime, EvaluateDetail string, ProjectSsociatedId int, IsPrivate bool,db *sql.DB) ([]string,error) {
	fmt.Println("PostEvaluate")
	sqlstatement := `INSERT
					INTO evaluate(evaluate_id, estimator_id, estimatored_id, evaluate_time, evaluate_time, evaluate_detail, project_ssociated_id, is_privacy)
					VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	stmt, err := db.Prepare(sqlstatement)
	defer stmt.Close()
	if err != nil{
		fmt.Println("err1 = ", err)
		return nil, err
	}
	_, err = stmt.Exec(EvaluateId, EstimatorId, EstimatoredId , EvaluateTime, EvaluateDetail , ProjectSsociatedId , IsPrivate)
	if err != nil{
		return nil, err
	}
	fmt.Println("stmt =", stmt)
	return nil, nil
}


