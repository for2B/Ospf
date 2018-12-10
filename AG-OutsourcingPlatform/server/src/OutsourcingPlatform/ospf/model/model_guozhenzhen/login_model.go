package model_guozhenzhen

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"sync"
	"log"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

type UserModel struct {
	*infrastructure.Infrastructure
	Phone string
	Password string
}

var onceLoginModel sync.Once
var userModel *UserModel

func GetUserModel(i *infrastructure.Infrastructure) *UserModel {
	onceLoginModel.Do(func() {
		userModel = &UserModel{
			Infrastructure: i,
		}
	})
	return userModel
}

func (m *UserModel)Getuser(usertel string) (UserModel,error){
   var user UserModel
   user.Phone=usertel
   rows,err:=m.DB.Query("select user_password from users where user_tel = $1",usertel)
   if err!=nil {
   	log.Println("err",err)
	  return user, httpapi.NewErr(constant.GLOBAL_DB_ERR,"DB.Query fail",nil)
   }
   for rows.Next(){
   	rows.Scan(&user.Password)
   }
	rows.Close()
   return user,nil
}



