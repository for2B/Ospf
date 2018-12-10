package projectRelatedPersonModel

import (
	"strconv"
	"database/sql"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)



func(m *ProjectRelatedPersonModel)GetTeammate(StrProjectId,OpenId string) (interface{},error) {
	//Checking if parameter are empty
	if StrProjectId == "" {
		m.Logger.Error("StrProjectId is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectId is metpy",nil)
	}

	if OpenId == "" {
		m.Logger.Error("OpenId is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"OpenId is empty",nil)
	}

	//Converting parameter type
	IntProject , err := strconv.Atoi(StrProjectId)
	if err!= nil {
		m.Logger.Error("Fail to convert StrProject to int")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProject to int",nil)
	}

	//Checking if parameter is invalid
	if IntProject < 0 {
		m.Logger.Error("IntProject is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProject is invalid",nil)
	}

	//get UserId
	var UserId int
	err = m.DB.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			m.Logger.Error("Could not find the user according to OpenId")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		m.Logger.Error("Fail to QueryRow for searching user")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}

	//get teammate
	var Teammates []Teammate
	row , err := m.DB.Query(`SELECT user_id , identity_id FROM teammate WHERE project_id = $1 AND identity_id!=2;`,StrProjectId)
	defer row.Close()
	if err != nil {
		m.Logger.Error("Fail to select teammate")
		return nil , httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to select teammate",err)
	}

	for row.Next(){
		var teammate Teammate
		err = row.Scan(
			&teammate.Uid,
			&teammate.Identity)
		if err!=nil {
			m.Logger.Error("Fail to scan row for teammate")
			return nil , httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to sacn row for teammate",err)
		}

		//标记团员中是否有用户本身
		if UserId == teammate.Uid{
			teammate.IsMe = 1
		}else{
			teammate.IsMe = 0
		}

		//将承包者放在数组第一位
		if teammate.Identity == 1{
			var teammates = []Teammate{teammate}
			Teammates = append(teammates,Teammates...)
		}else{
			Teammates = append(Teammates,teammate)
		}
	}

	return Teammates,nil
}