package model_guochuting

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
	"database/sql"
	"time"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

var oncePostEvaluateModel sync.Once

type PostEvaluateModel struct {
	*infrastructure.Infrastructure
}


var postEvaluateModel *PostEvaluateModel

func EvaluateModel(i *infrastructure.Infrastructure) *PostEvaluateModel {
	oncePostEvaluateModel.Do(func() {
		postEvaluateModel = &PostEvaluateModel{
			Infrastructure: i,
		}
	})
	return postEvaluateModel
}

func (m *PostEvaluateModel) Postevaluatemodel(EstimatorId, EstimatoredId int, EvaluateTime , EvaluateDetail string, ProjectSsociatedId int, IsPrivate bool, db *sql.DB) (interface{}, error){
	//fmt.Println("PostEvaluateMoel")
	//fmt.Println("data=", EstimatorId, EstimatoredId, EvaluateTime)
	sqlstatement := `INSERT
					INTO evaluate(estimator_id, estimatored_id, evaluate_time, evaluate_detail, project_ssociated_id, is_privacy)
					VALUES ($1, $2, $3, $4, $5, $6);
	`
	stmt, err := db.Prepare(sqlstatement)
	defer stmt.Close()
	if err != nil{
		//fmt.Println("err25=", err)
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "tx.Prepare func invalid", err)
	}
	EvaluateTime = time.Now().Format("2006-01-02")
	//fmt.Println("Time=", EvaluateTime)
	_, err = stmt.Exec(EstimatorId, EstimatoredId,EvaluateTime, EvaluateDetail, ProjectSsociatedId, IsPrivate)

	if err != nil {
		//fmt.Println("err3=", err)
		return nil, httpapi.NewErr(constant.GLOBAL_PARAM_ERR, "db.Query func failed", err)
	}
	return nil, nil
}
