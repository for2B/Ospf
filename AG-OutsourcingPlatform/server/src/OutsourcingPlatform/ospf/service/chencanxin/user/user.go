package user

import (
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/infrastructure"
	"database/sql"
	"errors"
)

var ErrOfNilUser = errors.New("User struct pointer can not be nil!")

type User struct {
	Name     string
	Phone    string
	Status   string //状态:"启用" "禁用"
	Avatar   string //头像图片的url地址
	Openid   string
	UserType string //取值有证明人或者普通开发者
	*infrastructure.Infrastructure
}

/*
*新建一个user结构体,usertype属性默认为游客类型,opeinid默认为空
 */
func NewUser(i *infrastructure.Infrastructure) User {
	u := User{
		UserType:       constant.USER_VISITOR,
		Openid:         "",
		Infrastructure: i,
	}
	return u
}

/*
该方法用于获取该用户是否为新用户,新用户返回true,否则返回false
*/
func (u *User) IsNewUser(db sql.DB) bool {
	//db.Prepare()

	return true
}

/*
该方法用于获取该用户的openid
*/
func (u *User) GetOpenid() error {
	if u == nil {
		return ErrOfNilUser
	}

	return nil
}
