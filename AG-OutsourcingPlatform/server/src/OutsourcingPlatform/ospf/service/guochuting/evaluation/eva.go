package evaluation

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"database/sql"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"fmt"
)

type Evaluation struct {
	HeadImg string  					`json:"head_img"`
	EstimatorName int64	 				`json:"estimator_name"`
	EvaluateTime int64	 				`json:"evaluate_time"`
	EvaluateDetail string				`json:"evaluate_detail"`
	ProjectSsociatedName string			`json:"project_ssociated_name"`
	*infrastructure.Infrastructure 		`json:"-"`
}

type Data struct {
	TotalPage string `json:"totalPage"`
	UserArray []Evaluation `json:"userArray"`
}

func NewEvaluation(i *infrastructure.Infrastructure) *Evaluation {
	evaluation := &Evaluation{
		Infrastructure:i,
	}
	return evaluation
}

func GetEvaluation(user_id int,key_user_name string,current_page int,show_nums int,db *sql.DB) ([]string,error) {

	var ans []string
	rows, err := db.Query("select user_name from users")
	if err != nil {
		return nil, httpapi.NewErr(constant.GLOBAL_DB_ERR,"DB.Query fail",nil)
	}

	for rows.Next() {
		description := ""
		rows.Scan(&description)
		ans = append(ans, description)
	}
	fmt.Println(ans)
	//return ans, nil
	return ans, nil


}




//func checkGPLPparams(user_id int ,key_user_name string,current_page int,show_nums int)(err error){
//	if current_page<0 {
//		err =httpapi.NewErr(constant.CHECK_PARAMS_ERR,"param err",nil)
//		return
//	}
//
//	if show_nums<0{
//		err =httpapi.NewErr(constant.CHECK_PARAMS_ERR,"param err",nil)
//		return
//	}
//
//
//	return nil
//}
