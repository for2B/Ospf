package controller_chencanxin

import (
	"OutsourcingPlatform/ospf/model/model_chencanxin"
	"net/http"
	"OutsourcingPlatform/internal/httpparse"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"encoding/json"
	"os"
	"fmt"
	"OutsourcingPlatform/utils"
	"strings"
	"io/ioutil"
)

type GenerateGolangApiCodeController struct {
	*model_chencanxin.GenerateGolangApiCodeModel
}

type GenerateGolangApiCodeControllerData struct {
	ApiName            string `json:"apiName"`
	Developer          string `json:"developer"`
	ControllorCode     string `json:"controllorCode"`
	ModelFileCode      string `json:"modelFileCode"`
	RegistFunctionCode string `json:"registFunctionCode"`
	RegistStatement    string `json:"registStatement"`
}

func (c *GenerateGolangApiCodeController) ToGenerateGolangApiCode(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// check request method
	if r.Method != "POST" {
		return nil, httpapi.NewErr(constant.REQUEST_METHOd_FAIL, "must be post method", nil)
	}

	//get api param
	getbody, err := httpparse.GetBody(r)
	if err != nil {
		return nil, httpapi.Err{Code: constant.HTTP_ERR_READBODY, Msg: "get body param failed"}
	}

	var data GenerateGolangApiCodeControllerData
	err = json.Unmarshal([]byte(getbody), &data)
	if err != nil {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "UnmJson failed", err)
	}

	// check param
	if &data == nil {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "&data can not be nil", nil)
	}
	if data.ControllorCode == "" {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "data.controllorCode can not be an empty string", nil)
	}
	if data.ModelFileCode == "" {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "data.modelFileCode can not be an empty string", nil)
	}
	if data.RegistFunctionCode == "" {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "data.registFunctionCode can not be an empty string", nil)
	}
	if data.RegistStatement == "" {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "data.registStatement can not be an empty string", nil)
	}
	if data.Developer == "" {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "data.developer can not be an empty string", nil)
	}
	if data.Developer != "chencanxin" && data.Developer != "guochuting" && data.Developer != "chenhuiliang" && data.
		Developer != "panjiawei" && data.Developer != "guozhenzhen" {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "data.developer can only be 'chencanxin' ,"+
			"'guochuting','chenhuiliang','panjiawei','guozhenzhen'", nil)
	}
	if data.ApiName == "" {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "data.ApiName can not be an empty string", nil)
	}else{
	//	ensure the first letter is upper
	splitResult:=strings.SplitAfterN(data.ApiName,"",2)
	splitResult[0]=strings.ToUpper(splitResult[0])
	data.ApiName=splitResult[0]+splitResult[1]
	}

	//create new file
	err = CreateAllFile(data)
	if err != nil {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "CreateAllFile  failed", err)
	}

	//append registCode to httpserver.go file
	path, err := utils.GetProDir() //path = .../server/bin
	if err != nil {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "utils.GetProDir failed", err)
	}
	path = strings.Trim(path, "bin")
	finalPath := path + "src/OutsourcingPlatform/ospf/httpserver_" + data.Developer + ".go"
	appendToFile(finalPath, data.RegistFunctionCode)

	// add regist statement to file
	AddRegistStatement(finalPath, data.RegistStatement)
	return nil, nil
}

//创建*_model.go文件和*_controller.go文件,并写入api参数传来的内容
func CreateAllFile(data GenerateGolangApiCodeControllerData) error {
	controllerFileName := fmt.Sprintf("%s_controller.go", data.ApiName)
	modelFileName := fmt.Sprintf("%s_model.go", data.ApiName)
	err := CreateFile(controllerFileName, data.ControllorCode, data.Developer, "controller")
	if err != nil {
		return httpapi.NewErr(constant.HTTP_ERR_READBODY, "CreateFile controllerFileName  failed", err)
	}
	err = CreateFile(modelFileName, data.ModelFileCode, data.Developer, "model")
	if err != nil {
		return httpapi.NewErr(constant.HTTP_ERR_READBODY, "CreateFile modelFileName failed", err)
	}
	return nil
}

/*
作用:在参数filename指定的路径创建文件,在文件中写入参数content内容.
参数developerName:开发者姓名:取值只能下列其中一个:chencanxin,chenhuiliang,guochuting,guozhenzhen, panjiawei.
参数fileType:controller或者model两种取值.
*/
func CreateFile(fileName, content, developerName, fileType string) error {
	dir, err := utils.GetProDir()
	if err != nil {
		return httpapi.NewErr(constant.HTTP_ERR_READBODY, "utils.GetProDir failed", err)
	}

	dir = strings.Trim(dir, "bin")
	finalDir := dir + "src/OutsourcingPlatform/ospf/" + fileType + "/" + fileType + "_" + developerName + "/"

	f, err := os.Create(finalDir + fileName)
	defer f.Close()
	if err != nil {
		return httpapi.NewErr(constant.HTTP_ERR_READBODY, "os.Create failed", err)
	}
	_, err = f.Write([]byte(content))
	if err != nil {
		return httpapi.NewErr(constant.HTTP_ERR_READBODY, "f.Write failed", err)
	}
	return nil
}

/*
在指定开发者的*_httpserver.go文件中追加注册新接口的函数体
*/
func appendToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "os.OpenFile(fileName, os.O_WRONLY, 0644)", err)
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

/*
在指定开发者的*_httpserver.go文件中追加注册新接口的函数调用
*/
func AddRegistStatement(filepath, newRegistStatement string) error {

	//read the origin string
	inputFile := filepath
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "ioutil.ReadFile(inputFile) failed", err)
	}
	fileString := string(buf)
	//fmt.Printf("%s\n", fileString)

	//get the final string
	s := strings.SplitAfterN(fileString, "(s.router)", 2)
	finalString := s[0] + "\n	" + newRegistStatement  + s[1]

	// write final buf to file
	finalBuf := []byte(finalString)
	err = ioutil.WriteFile(inputFile, finalBuf, 0x644)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "ioutil.WriteFile(inputFile, buf, 0x644) failed", err)
	}
	return nil
}
