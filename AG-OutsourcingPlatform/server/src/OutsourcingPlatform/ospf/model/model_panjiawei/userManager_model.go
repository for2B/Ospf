package model_panjiawei

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
	"OutsourcingPlatform/internal/httpapi"
	"time"
	"strconv"
	"OutsourcingPlatform/ospf/constant"
	"fmt"
)

var ONCE sync.Once
type  UserManagerModel struct {
	*infrastructure.Infrastructure
}
type UserInfo struct {
	Uid int64
	Name string
	Phone string
	Location string
	Create_time time.Time
	Change_time time.Time
	Permissions int
	State int
}

type UserFullDetail struct {
	User_id int64
	Invited_id int64
	User_name	string
	Openid	string
	Head_img	string
	User_tel	string
	User_sex	string
	User_idcard	string
	User_bank	string
	Permissions int
	Regist_time	time.Time
	Personal_update_time time.Time
	Permissiongs_update_time time.Time
	Prove_data string
	Expertise_field	string
	Wechat	string
	Address	string
	Level	int
	User_state	int
	Invitation_qrcode_token_normal	string
	Invitation_qrcode_token_witness	string
}

type UserResp struct {
	Alluser []UserInfo
	Count int64
}
func NewUserManagerModel(i *infrastructure.Infrastructure) (um *UserManagerModel){
		ONCE.Do(func() {
		um = &UserManagerModel{
			Infrastructure:i,
		}
	})
	return um
}

func (m *UserManagerModel)Change_permissions(userId []int64,OperateType string,PermissionType string) error{

	fmt.Println(userId)
	fmt.Println(OperateType)
	fmt.Println(PermissionType)

	getNeedPermission := -1
	switch PermissionType {
	case "AC_1":
		getNeedPermission = constant.AC_LOGIN
		break
	case "AC_2":
		getNeedPermission = constant.AC_Undertake_project
		break
	case "AC_4":
		getNeedPermission = constant.AC_Release_project
		break
	case "AC_8":
		getNeedPermission = constant.AC_COMMENT
		break
	default:
		return httpapi.NewErr(530,"PermissionType error",nil)
	}

	if OperateType!="add" && OperateType!= "remove" {
		return httpapi.NewErr(531,"OperateType error",nil)
	}

	// 把userId 数组转换 为如（1,2,3,4），方便数据库语句只调用一次
	uidList := ""
	for _,id := range userId {
		idStr := strconv.FormatInt(id,10)
		if idStr != "" {
			uidList+=(idStr+",")
		}
	}
	if uidList == "" || len(uidList) == 0 {
		return nil
	}
	uidList = uidList[:len(uidList)-1]
	uidList = "("+uidList+")"
	fmt.Println("uidList",uidList)
	fmt.Println("getNeedPermission", getNeedPermission)
	//for _,id :=range userId {
		if OperateType =="add"{
			_,err:=m.DB.Exec(`update users set  permissions = permissions|$1  where user_id in `+uidList,getNeedPermission)
			if err!= nil {
				return httpapi.NewErr(532,"QueryRow error ",err)
			}
		} else if OperateType =="remove" {
			getNeedPermission = constant.AC_ALL^getNeedPermission
			fmt.Println(getNeedPermission)
			fmt.Println(15&getNeedPermission)
			_,err:=m.DB.Exec(`update users set  permissions = permissions&$1  where user_id in `+uidList,getNeedPermission)
			if err!= nil {
				return httpapi.NewErr(532,"QueryRow error ",err)
			}
		}
	//}
	return nil
}

// 查询用户
func (m *UserManagerModel)Get_users(page,limit,permission,state int,keyword string) (UserResp,error){

	var permission_statement string = ""
	var state_statement string = ""

	if permission == constant.AC_LOGIN || permission == constant.AC_Undertake_project ||
		permission == constant.AC_Release_project || permission == constant.AC_COMMENT {
		permission_statement = " and "+strconv.Itoa(permission)+"&permissions>0"
	}

	if state == 0  {state_statement =  "and 1&permissions=0"}
	if state == 1  {state_statement =  "and 1&permissions>0"}

	if keyword != "%" {
		keyword = "%"+keyword+"%"
	} else {
		keyword = "%"
	}
	where_statement := " where (user_name like $1 or user_tel like $1 or address like $1)"
	where_statement = where_statement+permission_statement+state_statement;
	users := []UserInfo{}

	rows,err :=m.DB.Query("select user_id,user_name,user_tel,address,regist_time,"+
			"personal_update_time,permissions from users "+where_statement+" order by user_id ASC LIMIT $3 OFFSET $2",keyword,(page-1)*limit,limit)
	if err!=nil {
		return UserResp{},httpapi.NewErr(513,"【m.DB.Query(statment,(page-1)*limit,limit)】 error",err)
	}
	for rows.Next() {
		user := UserInfo{}
		err = rows.Scan(&user.Uid,&user.Name,&user.Phone,&user.Location,&user.Create_time,&user.Change_time,&user.Permissions)
		user.State = 1&user.Permissions
		if err!=nil{continue}
		users =append(users,user)
	}
	var allnum int64 = 0
	err =m.DB.QueryRow(`select count(*) from users ` +where_statement,keyword).Scan(&allnum)
	if err!=nil {
		return UserResp{},httpapi.NewErr(513,"【m.DB.Query(`select count(*) from users;)】 error",err)
	}

	return UserResp{users,allnum},nil
}

func (m *UserManagerModel)Get_one_user(uid int) (UserFullDetail,error){
	statement := "select user_id,invited_id,user_name,openid,head_img,user_tel,user_sex,user_idcard,user_bank,permissions,regist_time,personal_update_time,"+
		"permissiongs_update_time,prove_data,expertise_field,wechat,address,level,"+
		"user_state,invitation_qrcode_token_normal,invitation_qrcode_token_witness"+ " from users where user_id=$1"
	rows :=m.DB.QueryRow(statement,uid)

	user := UserFullDetail{}
	err := rows.Scan(
		&user.User_id ,
		&user.Invited_id ,
		&user.User_name,
		&user.Openid,
		&user.Head_img,
		&user.User_tel,
		&user.User_sex,
		&user.User_idcard,
		&user.User_bank	,
		&user.Permissions ,
		&user.Regist_time	,
		&user.Personal_update_time ,
		&user.Permissiongs_update_time ,
		&user.Prove_data,
		&user.Expertise_field,
		&user.Wechat,
		&user.Address,
		&user.Level,
		&user.User_state,
		&user.Invitation_qrcode_token_normal,
		&user.Invitation_qrcode_token_witness,
		)
	if err!=nil {
		return UserFullDetail{},httpapi.NewErr(514,"【Get_one_user rows.Scan】",err)
	}

	return user,err
}

func (m *UserManagerModel)Delete_users(userId []int64) (error){
	// 把userId 数组转换 为如（1,2,3,4），方便数据库语句只调用一次
	uidList := ""
	for _,id := range userId {
		idStr := strconv.FormatInt(id,10)
		if idStr != "" {
			uidList+=(idStr+",")
		}
	}
	if uidList == "" || len(uidList) == 0 {
		return nil
	}
	uidList = uidList[:len(uidList)-1]
	uidList = "("+uidList+")"
	fmt.Println("uidList",uidList)


	_,err:=m.DB.Exec("delete from users where user_id in "+uidList)
	return err
}