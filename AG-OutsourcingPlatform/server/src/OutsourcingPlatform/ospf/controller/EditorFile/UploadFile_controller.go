package EditorFile

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
	"OutsourcingPlatform/ospf/model/editorFileModel"
)

type Link struct {
	Link string `json:"link"`
}

type UpLoadFileController struct {
	*EditorFileModel.UploadFileModel
}

var PATH = "./bin/project_item_src"
var linkpath = "http://chl.ish2b.cn/ospf/files/"

//获取前端上传文件并保存到PATH目录下
func (c *UpLoadFileController) HandleUploadFile(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	if r.Method == "POST" {

		Md5Time := utils.GetMd5Time()

		//get file
		ReadFile, handler, err := r.FormFile("filename")
		if err != nil {
			c.Logger.Error("get r.FormFile failed", zap.String("Err:", err.Error()))
			return nil, err
		}
		defer ReadFile.Close()

		FileName := c.UploadFileModel.FilePath + "/project_item_src/" + Md5Time + handler.Filename

		//write ReadFile into WriteFile
		WriteFile, err := os.OpenFile(FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			c.Logger.Error("OpenFile failed", zap.String("Err:", err.Error()))
			return nil, err
		}
		defer WriteFile.Close()
		_, err = io.Copy(WriteFile, ReadFile)
		if err != nil {
			c.Logger.Error("copy file failed", zap.String("Err:", err.Error()))
			return nil, err
		}

		//return file's link
		lk := Link{Link: linkpath + Md5Time + handler.Filename}

		buf, err := json.Marshal(lk)
		if err != nil {
			c.Logger.Error("json Marshal fail", zap.String("Err:", err.Error()))
			return nil, err
		}

		fmt.Fprint(w, string(buf))
	}

	return nil, httpapi.NewErr(constant.REQUEST_METHOd_FAIL, "request method fail", nil)
}
