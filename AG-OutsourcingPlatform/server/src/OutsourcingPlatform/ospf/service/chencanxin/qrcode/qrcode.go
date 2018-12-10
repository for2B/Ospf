package qrcode

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"OutsourcingPlatform/utils"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"os"
	"io/ioutil"
	"strings"
)

type QrcodeManager struct {
}

/*
作用:用于应用启动时新键一个二维码管理器

*/
func NewQrcodeManager() QrcodeManager {
	var qrcode QrcodeManager
	return qrcode
}

/*
作用:新建一个二维码图片
输入:
输出:
*/
func (q *QrcodeManager) CreateQrcode(content string, level qrcode.RecoveryLevel, size int, filename string) error {
	err := qrcode.WriteFile(content, level, size, filename)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "qrcode.writefile() failed", err)
	}
	return nil
}

/*
作用:新建一个邀请二维码图片
输入:带token的链接,token(以token作为文件名称),后端服务器域名
输出:返回图片请求地址
*/
func (q *QrcodeManager) Create_Invitation_Qrcode(sdomain, content, token string) (string, error) {
	path, err := GetInvitationQrcodePath(token)
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "GetInvitationQrcodePath() failed in create_invitation_qrcode() ", err)
	}
	err = q.CreateQrcode(content, qrcode.Medium, 256, path)
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "CreateQrcode() failed in create_invitation_qrcode() ", err)
	}
	responseURL := "http://" + sdomain + "/assets/qrcode/" + token + ".png"
	return responseURL, nil
}

/*
作用:删除指定文件名称的二维码图片
输入:token(文件名称)
*/
func (q *QrcodeManager) Delete_Invitation_Qrcode(token string) (error) {
	path, err := GetInvitationQrcodePath(token)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "GetInvitationQrcodePath() failed in create_invitation_qrcode() ", err)
	}
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

/*
作用:获取指定token的邀请二维码的文件路径
输入:token(即文件名称)
输出:邀请二维码的文件路径
*/
func GetInvitationQrcodePath(token string) (string, error) {
	demopath := "%s/dist/assets/qrcode/%s.png"
	binpath, err := utils.GetProDir()
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "getprodir () failed in create_invitation_qrcode() ", err)
	}
	path := fmt.Sprintf(demopath, binpath, token)
	return path, nil
}

/*
作用:获取存放邀请二维码的目录路径
输入:无
输出:存放邀请二维码的目录路径
*/
func GetInvitationQrcodeDir() (string, error) {
	demopath := "%s/dist/assets/qrcode"
	binpath, err := utils.GetProDir()
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "GetInvitationQrcodeDir () failed in  ", err)
	}
	path := fmt.Sprintf(demopath, binpath)
	return path, nil
}

/*
作用:删除目录中所有过期的二维码图片和cache中过期的token
输入:
输出:删除的token名称,对应的过期时间,对应的生成者用户ID
*/
func (q *QrcodeManager) ClearInvalidQrcode(checkAndDel func(string)) (error) {
	pngs, err := q.GetDirPngList()
	if err != nil {
		return err
	}
	if len(pngs)==0{
		return nil
	}
	//check exiretiem one by one
	for _, p := range pngs {
		checkAndDel(p)
	}

	return nil
}

/*
作用:获取目录下所有的二维码图片列表,截取token组成数组返回
输入:无
输出:token为值的字符串数组
*/
func (q *QrcodeManager) GetDirPngList() ([]string, error) {
	dir, err := GetInvitationQrcodeDir()
	if err != nil {
		return nil, err
	}
	fmt.Println(dir)
	files, _ := ioutil.ReadDir(dir)
	var png []string
	for _, f := range files {
		token := f.Name()
		token = strings.TrimRight(token, ".png")
		png = append(png, token)
	}
	return png, nil
}
