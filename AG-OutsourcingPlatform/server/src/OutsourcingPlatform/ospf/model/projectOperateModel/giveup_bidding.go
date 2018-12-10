package projectOperateModel

import (
	"database/sql"
	"strconv"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
)

//根据项目Id和OpenId，删除表中数据
func (m *ProjectOperateModel)GiveUpBidding(StrProjectId ,OpenId string) error{

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

	//先查找一下表中是否有竞标的记录
	var projectId int
	err = m.DB.QueryRow(`SELECT project_id FROM competitors WHERE user_id = $1;`,UserId).Scan(&projectId)
	if err != nil {
		if err == sql.ErrNoRows {
			m.Logger.Error("")
			return httpapi.	NewErr(constant.SQL_OPERATE_FAIL,"can't find the bidding record in table",err)
		}
		m.Logger.Error("Fail to select bidding record")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to select bidding record",err)
	}

	//删除对应记录
	stmt,err := m.DB.Prepare(`DELETE FROM competitors WHERE project_id = $1 AND user_id = $2;`)
	defer stmt.Close()
	if err != nil {
		m.Logger.Error("Fail to get stmt")
		return httpapi.NewErr(constant.DB_BEGIN_FAIL,"Fail to get stmt",err)
	}
	_,err = stmt.Exec(IntProjectId , UserId)
	if err != nil {
		m.Logger.Error("Fail to exec delete")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to exec delete",err)
	}

	return nil
}
