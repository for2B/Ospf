package getProjectModel

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
)

var onceGetProject sync.Once

type GetProjectModel struct {
	*infrastructure.Infrastructure
}

type Project struct {
	ProjectName string  		`json:"project_name"`
	ProjectId int64	 			`json:"project_id"`
	PublisherId int64	 		`json:"publisher_id"`
	ProjectPublicTime string	`json:"project_publish_time"`
	ProjectCircle int		`json:"project_circle"`
	ProjectDetail string		`json:"project_detail"`
	ProjectMoney float64			`json:"project_money"`
	ProjectStatus int			`json:"project_status"`
	CompetitorsNums int 		`json:"competitors_nums"`
	*infrastructure.Infrastructure 	`json:"-"`
}

type ProjectList struct {
	Total int					`json:"total"`
	Projects []Project          `json:"project_list"`
}

type NumberOfEachStatus struct {
	Total 		int `json:"total"`
	StatusOne 	int `json:"status_one"`
	StatusTwo 	int `json:"status_two"`
	StatusThree int `json:"status_three"`
	StatusFour 	int `json:"status_four"`
	StatusFive 	int `json:"status_five"`
}

var getProjectModel *GetProjectModel

func ProduceGetProjectModel (i *infrastructure.Infrastructure) *GetProjectModel{
	onceGetProject.Do(func() {
		getProjectModel = &GetProjectModel{
			Infrastructure:i,
		}
	})
	return getProjectModel
}

