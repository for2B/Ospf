package projectRelatedPersonModel

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
)

var onceProjectRelatedPerson sync.Once

type ProjectRelatedPersonModel struct {
	*infrastructure.Infrastructure
}

type Bidder struct {
	UserId             int    `json:"user_id"`
	Name 			   string `json:"name"`
	HeadImg            string `json:"head_img"`
	ExpertiseField     string `json:"expertise_field"`
	Address            string `json:"address"`
	FinishProjectCount int    `json:"finish_project_count"`
	Identity           int    `json:"identity"`
}

type BidderList struct {
	Total int 			`json:"total"`
	Bidders []Bidder 	`json:"bidding_list"`
}

type ProjectIdentity struct {
	Identity int `json:"identity"`
}

type Teammate struct {
	Uid 	 int  `json:"Uid"`
	Identity int  `json:"Identity"`
	IsMe 	 int  `jsonL:"IsMe"`
}


var projectRelatedPersonModel *ProjectRelatedPersonModel

func ProduceProjectRelatedPersonModel(i *infrastructure.Infrastructure) *ProjectRelatedPersonModel {
	onceProjectRelatedPerson.Do(func() {
		projectRelatedPersonModel = &ProjectRelatedPersonModel{
			Infrastructure:i,
		}
	})
	return projectRelatedPersonModel
}