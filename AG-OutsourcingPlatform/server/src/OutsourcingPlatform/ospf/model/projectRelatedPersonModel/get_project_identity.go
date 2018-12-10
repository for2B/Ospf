package projectRelatedPersonModel

import (
	"strconv"
	"database/sql"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

const CaCheBucket = "ProjectIdentity"

//根据ProjectId获取该用户相对于某一项目的身份身（1 承包者 2 发布者 3普通成员 4路人）
func (m *ProjectRelatedPersonModel)GetProjectIdentity(StrProjectId,OpenId string) (interface{},error){

	//Checking if parameters are empty
	if StrProjectId == "" {
		m.Logger.Error("StrProjectId is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrProjectId is empty",nil)
	}

	if OpenId == "" {
		m.Logger.Error("OpenId is empty")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"OpenId is empty",nil)
	}

	//get from cache
	var Identity interface{}
	err := m.CacheManager.Get(CaCheBucket,[]string{StrProjectId,OpenId},&Identity)
	if err!=nil{
		m.Logger.Error("Get From Cache fail"+err.Error())
	}
	if Identity!=nil{
		return Identity,nil
	}

	//Converting parameter type
	IntProjectId , err := strconv.Atoi(StrProjectId)
	if err != nil {
		m.Logger.Error("Fail to convert StrProjectId to int")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProject to int",nil)
	}

	//Checking if parameter is valid
	if IntProjectId < 0 {
		m.Logger.Error("IntProjectId is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProjectId is invalid",nil)
	}

	//get UserId
	var UserId int64

	err = m.DB.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			m.Logger.Error("Could not find the user according to OpenId")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		m.Logger.Error("Fail to QueryRow for searching user")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}

	//查找teammate表来获取用户对某一个项目的身份
	var identity ProjectIdentity
	err = m.DB.QueryRow(`SELECT identity_id FROM teammate WHERE project_id = $1 AND user_id= $2;`,IntProjectId,UserId).Scan(&identity.Identity)
	if err!=nil {
		if err == sql.ErrNoRows {
			identity.Identity = 4
		}else{
			m.Logger.Error("Fail to select Identity")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to select Identity",err)
		}
	}

	//判断是否是竞标者 竞标者的话身份也为1
	var user_id int
	if identity.Identity == 4{
		err = m.DB.QueryRow(`SELECT user_id FROM competitors WHERE project_id = $1 AND user_id = $2`,IntProjectId,UserId).Scan(&user_id)
		if err!=nil{
			if err != sql.ErrNoRows {
				m.Logger.Error("Fail to judge if the user is bidder")
				return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to judge if the user is bidder",nil)
			}
		}else{
			identity.Identity = 1
		}

	}

	//Checking if return value is valid
	if identity.Identity < 0 || identity.Identity > 4 {
		m.Logger.Error("Identity's values is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Identity's values is invalid",nil)
	}

	//Insert into cache
	m.CacheManager.Set(CaCheBucket,[]string{StrProjectId,OpenId},identity)

	return identity,nil
}
