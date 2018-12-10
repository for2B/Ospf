package projectRelatedPersonModel

import (
	"OutsourcingPlatform/ospf/model/projectOperateModel"
	"strconv"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

//取消当前项目的承接者，并更新项目状态为1（等待承接）
func (m *ProjectRelatedPersonModel)Cancelcontractor(StrProjectId string) error{

	//checking if the parameter is empty
	if StrProjectId == "" {
		m.Logger.Error("StrProjectId is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProjectId is empty",nil)
	}

	//Converting parameter type
	IntProjectId , err := strconv.Atoi(StrProjectId)
	if err !=nil {
		m.Logger.Error("Fail to convert StrProjectId to int")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProjectId to int",nil)
	}

	//Checking if the parameter is valid
	if IntProjectId  < 0 {
		m.Logger.Error("IntProject is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProject is invalid",nil)
	}

	//get ContractorId
	var ContractorId int
	err = m.DB.QueryRow(`SELECT contractor_id FROM project WHERE project_id = $1`,IntProjectId).Scan(&ContractorId)
	if err!=nil {
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"no contractor",err)
	}

	//set contractor_id = null
	stmt,err := m.DB.Prepare(`UPDATE project SET contractor_id = null WHERE project_id = $1;`)
	defer stmt.Close()
	if err!=nil {
		m.Logger.Error("Fail to get stmt")
		return httpapi.NewErr(constant.DB_BEGIN_FAIL,"Fail to get stmt",err)
	}
	_,err = stmt.Exec(IntProjectId)
	if err!=nil {
		m.Logger.Error("Fail to exec stmt")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to exec stmt",err)
	}

	//取消承接也需要删去teammate表中的记录
	stmt,err = m.DB.Prepare(`DELETE FROM teammate WHERE project_id = $1 AND user_id = $2`)
	if err!=nil {
		m.Logger.Error("Fail to get stmt")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get stmt",err)
	}
	_,err = stmt.Exec(IntProjectId,ContractorId)
	if err!=nil{
		m.Logger.Error("Fail to exec stmt")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"exec fail",err)
	}

	//取消项目承接人后要更新项目状态
	UpdateProStatus := projectOperateModel.ProduceProjectOperateModel(m.Infrastructure)
	err = UpdateProStatus.UpdateProjectStatus(StrProjectId,"1")
	if err!=nil{
		return err
	}

	return nil
}
