package model_chencanxin

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
	"OutsourcingPlatform/ospf/constant"
)

var onceLogModel sync.Once

type LogModel struct {
	*infrastructure.Infrastructure
}

var logModel *LogModel

func GetLogModel(i *infrastructure.Infrastructure) *LogModel {
	onceLogModel.Do(func() {
		logModel = &LogModel{
			Infrastructure: i,
		}
	})
	return logModel
}

func (m *LogModel) GetLog() ([]string, error) {
	var ans []string
	rows, err := m.DB.Query("select log from syslog order by id desc limit 10;")
	if err != nil {
		return nil, httpapi.NewErr(constant.GLOBAL_DB_ERR,"DB.Query fail",nil)
	}


	for rows.Next() {
		description := ""
		rows.Scan(&description)

		ans = append(ans, description)
	}

	return ans, nil
}
