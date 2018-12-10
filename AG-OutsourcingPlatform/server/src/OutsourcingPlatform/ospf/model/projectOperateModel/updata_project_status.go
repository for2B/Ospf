package projectOperateModel

import (
	"strconv"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

//根据项目ID和状态值更新项目状态
func (m *ProjectOperateModel)UpdateProjectStatus(StrProjectId,StrProjectStatus string) error{

	//Checking if parameters are empty
	if StrProjectId == "" {
		m.Logger.Error("StrProjectId is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrProjectId is empty",nil)
	}

	if StrProjectStatus == "" {
		m.Logger.Error("StrProjectStatus is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrProjectStatus is empty",nil)
	}

	//Converting parameters type
	IntProjectId , err := strconv.Atoi(StrProjectId)
	if err != nil {
		m.Logger.Error("Fail to convert StrProjectId to int")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProjectId to int",nil)
	}

	IntProjectStatus , err := strconv.Atoi(StrProjectStatus)
	if err != nil {
		m.Logger.Error("Fail to convert StrProjectStatus to int")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProjectStatus to int",nil)
	}

	//Checking if parameters are valid
	if IntProjectId < 0 {
		m.Logger.Error("IntProjectId is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProjectId is invalid",nil)
	}
	if IntProjectStatus < 0 || IntProjectStatus > 4 {
		m.Logger.Error("IntProjectStatus is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProjectStatus is invalid",nil)
	}

	//update project_status
	stmt,err := m.DB.Prepare(`UPDATE project SET project_status = $1 WHERE project_id = $2;`)
	defer stmt.Close()
	if err!=nil {
		m.Logger.Error("Fail to get stmt")
		return httpapi.NewErr(constant.DB_BEGIN_FAIL,"Fail to get stmt",err)
	}
	_,err = stmt.Exec(IntProjectStatus,IntProjectId)
	if err!=nil {
		m.Logger.Error("Fail to exec stmt")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to exec stmt",err)
	}

	return nil
}