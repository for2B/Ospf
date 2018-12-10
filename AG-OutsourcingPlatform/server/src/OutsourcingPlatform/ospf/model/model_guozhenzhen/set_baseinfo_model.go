package model_guozhenzhen

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"log"
)

type SetBaseInfoModel struct {
	*infrastructure.Infrastructure
	User_id  int64
	Username string
	Sex      string
	Password string
	Wechat_account string
	Phone string
	Identity string
	Credit string
	Location string
}

var onceSetBaseModel sync.Once
var setBaseInfoModel *SetBaseInfoModel

func SetBaseInformationModel(i *infrastructure.Infrastructure) *SetBaseInfoModel {
	onceSetBaseModel.Do(func() {
		setBaseInfoModel = &SetBaseInfoModel{
			Infrastructure: i,
		}
	})
	return setBaseInfoModel
}

func (m *SetBaseInfoModel)SetBaseInfor(info SetBaseInfoModel) (SetBaseInfoModel,error){
	var baseinfo SetBaseInfoModel
	rows,err:=m.DB.Query("select user_id from users where user_tel = $1",info.Phone)
	if err!=nil {
		return baseinfo, httpapi.NewErr(constant.GLOBAL_DB_ERR,"DB.Query fail",nil)
	}
	for rows.Next() {
		rows.Scan(&baseinfo.User_id)
	}
	if baseinfo.User_id==0 {
		log.Print("34")
		sql:=`insert into users(user_name,user_sex,user_password,user_tel,wechat,user_idcard,user_bank,
    address) values($1,$2,$3,$4,$5,$6,$7,$8)`
		rowsaffected,err:=m.DB.Exec(sql,info.Username,info.Sex,info.Password,info.Phone,info.Wechat_account,info.Identity,
			info.Credit,info.Location)
		if err!=nil {
			return baseinfo, httpapi.NewErr(constant.GLOBAL_DB_ERR,"DB.Exec fail",nil)
		}
		log.Print(rowsaffected.RowsAffected())
	}else{
		sql:=`update users set user_name=$1, user_sex=$2,user_password=$3, user_tel=$4 ,wechat=$5, user_idcard=$6,
        user_bank=$7,address=$8  where user_id=$9`
		rowsaffected,err:=m.DB.Exec(sql,info.Username,info.Sex,info.Password,info.Phone,info.Wechat_account,info.Identity,
			info.Credit,info.Location,baseinfo.User_id)
		if err!=nil {
			return baseinfo, httpapi.NewErr(constant.GLOBAL_DB_ERR,"DB.Exec fail",err)
		}
		log.Print(rowsaffected.RowsAffected())
	}
	return baseinfo,nil
}