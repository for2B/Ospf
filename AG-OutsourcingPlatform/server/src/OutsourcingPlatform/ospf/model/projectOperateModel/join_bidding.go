package projectOperateModel

import (
	"strconv"
	"database/sql"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)


//根据项目ID和OpenId，在竞标表中插入数据
func (m *ProjectOperateModel)JoinBidding(StrProjectId ,OpenId string) error{

	//Checking if parameter are empty
	if StrProjectId == "" {
		m.Logger.Error("StrProjectId is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrProjectId is empty",nil)
	}

	if OpenId == "" {
		m.Logger.Error("OpenId is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"OpenId is empty",nil)
	}

	//Converting parameter type
	IntProjectId , err := strconv.Atoi(StrProjectId)
	if err != nil {
		m.Logger.Error("Fail to convert StrProjectId to int")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail convert StrProjectId to int",nil)
	}

	//Checking if parameter is valid
	if IntProjectId < 0 {
		m.Logger.Error("IntProjectId is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProjectId is invalid",nil)
	}

	//get UserId
	var UserId int64

	err = m.DB.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&UserId)
	if err!=nil{
		if err == sql.ErrNoRows {
			m.Logger.Error("Could not find the user according to OpenId")
			return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		m.Logger.Error("Fail to QueryRow for searching user")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}


	//将竞标记录插入表中
	stmt,err := m.DB.Prepare(`INSERT INTO competitors(project_id, user_id,compete_id) VALUES ($1,$2,CURRENT_DATE)`)
	defer stmt.Close()
	if err!=nil {
		m.Logger.Error("Fail to get stmt")
		return httpapi.NewErr(constant.DB_BEGIN_FAIL,"Fail to get stmt",err)
	}
	_,err = stmt.Exec(IntProjectId , UserId)
	if err!=nil{
		m.Logger.Error("Fail to insert record to competitors")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to insert record to competitors",err)
	}

	return nil

}








