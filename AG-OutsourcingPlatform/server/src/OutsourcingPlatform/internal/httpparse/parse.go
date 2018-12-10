package httpparse

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"context"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

/*
作用:获取请求方所携带的相关信息,存进上下文context
输入:
输出:
*/
func Parse(r *http.Request) (context.Context, error) {
	Form := make(map[string][]string)
	MultipartFormFile := make(map[string][]*multipart.FileHeader)
	MultipartFormValue := make(map[string][]string)
	ctx := r.Context()

	r.ParseMultipartForm(1 << 32)
	r.ParseForm()

	if len(r.Form) > 0 {
		for k, v := range r.Form {
			Form[k] = []string(v)
		}
	}

	if r.MultipartForm != nil && len(r.MultipartForm.File) > 0 {
		for k, v := range r.MultipartForm.File {
			MultipartFormFile[k] = []*multipart.FileHeader(v)
		}
	}

	if r.MultipartForm != nil && len(r.MultipartForm.Value) > 0 {
		for k, v := range r.MultipartForm.Value {
			MultipartFormValue[k] = []string(v)
		}
	}

	buf, _ := ioutil.ReadAll(r.Body)
	ctx = context.WithValue(ctx, "Body", string(buf))
	ctx = context.WithValue(ctx, "Form", Form)
	ctx = context.WithValue(ctx, "MultipartFormFile", MultipartFormFile)
	ctx = context.WithValue(ctx, "MultipartFormValue", MultipartFormValue)

	return ctx, nil
}

/*
作用:获取post提交过来的body内容
输入:
输出:
*/
func GetBody(r *http.Request) (s string, err error) {
	body, ok := r.Context().Value("Body").(string)
	if !ok {
		return "", errors.New("get body fail : body is not string")
	}

	return body, nil
}

/*
作用:获取url参数
输入:
输出:
*/
func GetForm(r *http.Request) (form map[string][]string, err error) {
	form, isOK := r.Context().Value("Form").(map[string][]string)
	if !isOK {
		return nil, errors.New("get form fail : form is not map[string][]string")
	}

	return form, nil
}

/*
作用:获取客户端通过multipart/form-data上传的内容
输入:
输出:
*/
func GetMultipartFormFile(r *http.Request) (form map[string][]*multipart.FileHeader, err error) {
	form, ok := r.Context().Value("MultipartFormFile").(map[string][]*multipart.FileHeader)
	if !ok {
		return nil, errors.New("get multipart form file fail : multipart form file is not map[string][]*multipart.FileHeader")
	}

	return form, nil
}

func GetMultipartFormValue(r *http.Request) (form map[string][]string, err error) {
	form, ok := r.Context().Value("MultipartFormValue").(map[string][]string)
	if !ok {
		return nil, errors.New("get multipart form value fail : multipart form value is not map[string][]string")
	}

	return form, nil
}

/*
作用:获取客户端对应保存在cache中的session会话信息
输入:
输出:
*/
func GetSession(r *http.Request) (session *constant.Session, err error) {

	session, ok := r.Context().Value("session").(*constant.Session)
	if !ok {
		return nil, httpapi.NewErr(constant.GLOBAL_ASSERT_ERR,
			"get session from context fail:is not session struct type", nil)
	}
	return session, nil
}

/*
作用:获取当前用户的openid
输入:
输出:
*/
func GetOpenid(r *http.Request) (openid string, err error) {

	session, ok := r.Context().Value("session").(*constant.Session)
	if !ok {
		return "", httpapi.NewErr(constant.GLOBAL_ASSERT_ERR,
			" getopenid() failed : get session from context fail:is not session struct type ", nil)
	}
	return session.Openid, nil
}
