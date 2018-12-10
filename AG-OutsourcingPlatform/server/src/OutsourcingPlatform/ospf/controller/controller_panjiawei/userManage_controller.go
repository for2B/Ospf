package controller_panjiawei

import (
	"OutsourcingPlatform/ospf/model/model_panjiawei"
	"net/http"
	"encoding/json"
	"OutsourcingPlatform/internal/httpapi"
	"strconv"
	"errors"

	"OutsourcingPlatform/ospf/constant"
)

type UserManagerController struct {
	*model_panjiawei.UserManagerModel
}

type Post struct {
	UserID []int64 `json:"user_id"`
	OperateType string `json:"operate_type"`
	PermissionType string `json:"permission_type"`
}

func (c *UserManagerController)getPostBody( r *http.Request)(Post,error){
	var post Post
	p :=r.Context().Value("Body").(string)
	err :=json.Unmarshal([]byte(p),&post)
	return post,err
}
func (c *UserManagerController)getUrlParms( r *http.Request, key string)(string,error){
	p :=r.Context().Value("Form").(map[string][]string)
	if  str:=p[key][0]; str != "" {
		return str,nil
	}
	return "",errors.New("not in it")
}
func (c *UserManagerController)ChangeUserPermissions(w http.ResponseWriter, r *http.Request) (interface{}, error){
	post,err := c.getPostBody(r)
	if err!= nil {
		return "",httpapi.NewErr(503,"【ChangeUserPermissions】 jgetPostBody",err)
	}
	err = c.Change_permissions(post.UserID,post.OperateType,post.PermissionType)
	isok := "ok"
	if err != nil {
		isok = "fail"
	}
	return isok,err
}
func (c *UserManagerController)DeleteUsers(w http.ResponseWriter, r *http.Request) (interface{}, error){
	post,err := c.getPostBody(r)
	if err!= nil {
		return "",httpapi.NewErr(503,"【DeleteUsers】 jgetPostBody",err)
	}
	err =  c.Delete_users(post.UserID)
	isok := "ok"
	if err != nil {
		isok = "fail"
	}
	return isok,err
}


func (c *UserManagerController)GetOneUser(w http.ResponseWriter, r *http.Request)(interface{}, error) {
	uid,err:=c.getUrlParms(r,"uid")
	uid_num,err:= strconv.Atoi(uid)
	if err!=nil  {
		return nil,httpapi.NewErr(522,"url pramas [uid] is invaild",err)
	}
	data,err := c.Get_one_user(uid_num)
	return data,err
}
func (c *UserManagerController)GetUsers(w http.ResponseWriter, r *http.Request)(interface{}, error) {


	limit,err:=c.getUrlParms(r,"limit")
	if err!=nil {
		return nil,httpapi.NewErr(521,"url pramas [limit] is null",err)
	}
	limit_num,err:= strconv.Atoi(limit)
	if err!=nil || limit_num<=0 {
		return nil,httpapi.NewErr(522,"url pramas [limit] is invaild",err)
	}

	page,err:=c.getUrlParms(r,"page")
	if err!=nil {
		return nil,httpapi.NewErr(523,"url pramas [page] is null",err)
	}
	page_num,err:= strconv.Atoi(page)
	if err!=nil || page_num<=0 {
		return nil,httpapi.NewErr(524,"url pramas [page] is invaild",err)
	}

	permission,_:=c.getUrlParms(r,"permission")
	permission_num,err :=strconv.Atoi(permission)
	if err!=nil || (permission_num!=0 && permission_num != constant.AC_LOGIN && permission_num != constant.AC_Undertake_project &&
		permission_num != constant.AC_Release_project && permission_num != constant.AC_COMMENT){
		return nil,httpapi.NewErr(525,"url pramas [permission] is invaild",err)
	}

	state,_:=c.getUrlParms(r,"state")
	state_num,_ :=strconv.Atoi(state)
	if err!=nil || (state_num!=0 && state_num != 1 && state_num != 4){
		return nil,httpapi.NewErr(526,"url pramas [state] is invaild",err)
	}
	keyword,_:=c.getUrlParms(r,"skeyword")

	//fmt.Println("page_num",page_num)
	//fmt.Println("limit_num",limit_num)
	//fmt.Println("permission_num",permission_num)
	//fmt.Println("state_num",state_num)
	//fmt.Println("skeyword",keyword)

	data,err:= c.Get_users(page_num,limit_num,permission_num,state_num,keyword)
	return data,err
}

