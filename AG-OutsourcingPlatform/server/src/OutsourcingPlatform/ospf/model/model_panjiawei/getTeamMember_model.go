package model_panjiawei

import (
"sync"
"OutsourcingPlatform/ospf/infrastructure"
	"strings"
	"OutsourcingPlatform/internal/httpapi"
)

var onceGetUploadFile sync.Once

type GetMemberModel struct {
	*infrastructure.Infrastructure
}
/*
"head_img":"o3i2uasjflj324uoiu.img",
             "user_name":"陈宇",
             "address":"广东省 广州市 番禺区",
             "expertise_field":["Go","Angular"],
 */
type Member struct{
	Uid int64
	Name string
	Phone string
	Location string
	Rating int64
	Head_img	string
	Expertise_field []string
}

func New_GetMemberModel (i *infrastructure.Infrastructure) (getmembermodel *GetMemberModel){
	onceGetUploadFile.Do(func() {
		getmembermodel = &GetMemberModel{
			Infrastructure:i,
		}
	})
	return getmembermodel
}

func (m *GetMemberModel) GetMemberInDB(userID int64)(interface{},error){
	member := Member{}
	statement := `select user_id,user_name,user_tel,address,head_img,expertise_field from users where user_id = $1;`
	row:=m.DB.QueryRow(statement,userID)
	var expertise_field string
	err := row.Scan(&member.Uid,&member.Name,&member.Phone,&member.Location,&member.Head_img,&expertise_field)
	if err != nil {
		return "",httpapi.NewErr(503,"【func GetMemberInDB】 rows scan error",err)
	}
	expertise_slice := strings.Split(expertise_field,",")
	member.Expertise_field = expertise_slice

	return member,nil
}
