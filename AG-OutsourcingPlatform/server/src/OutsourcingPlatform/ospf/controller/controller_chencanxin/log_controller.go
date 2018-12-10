package controller_chencanxin

import (
	"net/http"
	"OutsourcingPlatform/ospf/model/model_chencanxin"
	"OutsourcingPlatform/utils"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"strings"
	"fmt"
)

type LogController struct {
	*model_chencanxin.LogModel
}

func (c *LogController) GetLogs(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dir, err := utils.GetProDir()
	if err != nil {
		return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "utils.GetProDir failed", err)
	}
	//fmt.Println(dir)//.../server/bin
	dir = strings.Trim(dir, "bin")
	finalDir := dir + "src/OutsourcingPlatform/ospf/httpserver_chencanxin.go"

	fmt.Println(finalDir)

	return c.GetLog()
}
