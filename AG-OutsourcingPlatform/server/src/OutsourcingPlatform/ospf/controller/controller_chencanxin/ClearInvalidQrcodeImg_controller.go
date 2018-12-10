package controller_chencanxin

import (
	"net/http"
	"OutsourcingPlatform/ospf/service/chencanxin/qrcode"
	"OutsourcingPlatform/ospf/infrastructure"
)

type ClearQrcodeController struct {
	*infrastructure.Infrastructure
}

func (c *ClearQrcodeController) Clear(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	pm:=qrcode.NewQrcodeManager()
	err:=pm.ClearInvalidQrcode(func(token string) {
		CheckQrcodeToken(token,c.Infrastructure,false)
	})
	if err!=nil{
		return nil,err
	}
	return nil,nil
}
