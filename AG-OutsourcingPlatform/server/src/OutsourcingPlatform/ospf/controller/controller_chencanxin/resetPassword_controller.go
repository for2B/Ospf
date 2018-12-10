package controller_chencanxin

import (
	"net/http"
	"OutsourcingPlatform/ospf/model/model_chencanxin"
)


type ResetPasswordController struct {
	*model_chencanxin.ResetPasswordModel
}

func (c *ResetPasswordController) ToResetPassword(w http.ResponseWriter, r *http.Request) (interface{}, error) {


	return nil,nil
}
