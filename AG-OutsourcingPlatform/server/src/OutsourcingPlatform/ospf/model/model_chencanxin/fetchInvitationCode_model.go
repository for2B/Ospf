package model_chencanxin

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
	"OutsourcingPlatform/internal/httpapi"
	"database/sql"
	"OutsourcingPlatform/ospf/constant"
)

var onceFetchInvitationCodeModel sync.Once

type FetchInvitationCodeModel struct {
	*infrastructure.Infrastructure
}

var FetchInviCodeModel *FetchInvitationCodeModel

func GetFetchInviCodeModel(i *infrastructure.Infrastructure) *FetchInvitationCodeModel {
	onceFetchInvitationCodeModel.Do(func() {
		FetchInviCodeModel = &FetchInvitationCodeModel{
			Infrastructure: i,
		}
	})
	return FetchInviCodeModel
}

/*
作用:从数据库中获取当前用户的邀请码token,注意建表时邀请码token列默认为空字符串
输入:当前用户的openid
输出:返回token
*/
func (m *FetchInvitationCodeModel) GetQRcodeToken(openid,usertype string) (string ,error ) {
	qrcode:=""
	sqlstatement:="select invitation_qrcode_token_normal from users where openid = $1"
	if usertype==constant.REFEREE_USER_TYPE {
		sqlstatement="select invitation_qrcode_token_witness from users where openid = $1"
	}
	err:=m.DB.QueryRow(sqlstatement,openid).Scan(&qrcode)
	if err!=nil{
		if err == sql.ErrNoRows {
			return "",httpapi.NewErr(constant.GLOBAL_DB_ERR,"no this openid user row in table",err)
		}
		return "",httpapi.NewErr(constant.GLOBAL_DB_ERR,"GetQRcodeToken().db.QueryRow() failed",err)
	}
	return qrcode,nil
}

/*
作用:往数据库中插入当前用户的邀请码token,如果存在则覆盖
输入:token,openid
输出:无
*/
func (m *FetchInvitationCodeModel) InsertInvitationCodeToken(usertype,openid,token string) (error ) {
	//todo:
statement:="UPDATE users SET invitation_qrcode_token_normal = $1 WHERE openid = $2;"
if usertype==constant.REFEREE_USER_TYPE {
	statement="UPDATE users SET invitation_qrcode_token_witness = $1 WHERE openid = $2;"
}
	stmt, err := m.DB.Prepare(statement)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_DB_ERR, "m.DB.Prepare.update func fail", err)
	}
	result, err := stmt.Exec(token,openid)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_DB_ERR, "m.DB.Exec.update func fail", err)
	}


	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "result.RowsAffected func fail", err)
	}

	if rowsAffected == 0 {
		return httpapi.NewErr(constant.GLOBAL_PARAM_ERR, "the user openid is not exist", nil)
	}

	return nil
}

