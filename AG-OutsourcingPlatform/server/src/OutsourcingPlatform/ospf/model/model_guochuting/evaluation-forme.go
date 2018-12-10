package model_guochuting

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
	"fmt"
)


var onceGetEvaluationForMeModel sync.Once

type GetEvaluationForMeModel struct {
	*infrastructure.Infrastructure
}

// type User struct {
// 	Name           string         `json:"name"`
// 	Id             string         `json:"id"`
// 	Level          string         `json:"level"`
// 	Account        string         `json:"account"`
// 	Phone          string         `json:"phone"`
// 	Department     pq.StringArray `json:"department"`
// 	Identity       string         `json:"identity"`
// 	Progress       string         `json:"progress,omitempty"`
// 	DetailProgress string         `json:"detailprogress,omitempty"`
// 	Count          int            `json:"-"`
// }

var getEvaluationForMeModel *GetEvaluationModel

func EvaluationForMeModel(i *infrastructure.Infrastructure) *GetEvaluationModel {
	onceGetEvaluationModel.Do(func() {
		getEvaluationModel = &GetEvaluationModel{
			Infrastructure: i,
		}
	})
	return getEvaluationModel
}

func (m *GetEvaluationModel) GetEvaForMeModel(key_user_name, current_page, show_nums string) (interface{}, error){
	fmt.Println("GetevaluationModal")
	return nil, nil;
}

