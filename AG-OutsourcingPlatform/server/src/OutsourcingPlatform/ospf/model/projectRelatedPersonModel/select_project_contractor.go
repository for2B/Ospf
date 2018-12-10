package projectRelatedPersonModel

import (
	"strconv"
	"OutsourcingPlatform/ospf/model/projectOperateModel"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
)

//根据项目Id和用户Id，选择项目承包者并更新项目状态
func (m *ProjectRelatedPersonModel)SelectProjectContractor(StrProjectId,StrUserId string) error{

	//Checking if parameters are empty
	if StrProjectId == "" {
		m.Logger.Error("StrProjectId is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrProjectId is empty",nil)
	}

	if StrUserId == "" {
		m.Logger.Error("StrUserId is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrUserId is empty",nil)
	}

	//Converting parameters type
	IntProjectId , err := strconv.Atoi(StrProjectId)
	if err != nil {
		m.Logger.Error("Fail to convert StrProjectId to int")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProjectId to int",nil)
	}
	IntUserId , err := strconv.Atoi(StrUserId)
	if err != nil {
		m.Logger.Error("Fail to convert StrUserId to int")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to conver StrUserId to int",nil)
	}

	//Checking if parameters is valid
	if IntProjectId < 0 {
		m.Logger.Error("IntProjectId is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProjectId is invalid",nil)
	}

	if IntUserId < 0 {
		m.Logger.Error("IntUserId is Invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntUserId is invalid",nil)
	}

	//update contractor_id
	stmt,err := m.DB.Prepare(`UPDATE project SET contractor_id = $1 WHERE project_id = $2;`)
	defer stmt.Close()
	if err !=nil {
		m.Logger.Error("Fail to get stmt")
		return httpapi.NewErr(constant.DB_BEGIN_FAIL,"Fail to get stmt",err)
	}
	_,err = stmt.Exec(IntUserId,IntProjectId)
	if err != nil {
		m.Logger.Error("Fail to exec stmt")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to exec stmt",err)
	}

	//选择承包者后对应更新项目状态
	UpdateProStatus := projectOperateModel.ProduceProjectOperateModel(m.Infrastructure)
	err = UpdateProStatus.UpdateProjectStatus(StrProjectId,"2")
	if err!=nil {
		m.Logger.Error("Fail to UpdateProjectStatus")
		return err
	}

	//将承包者插入到teammate表中 1.承包者
	err = projectOperateModel.InsertIntoTeammate(IntProjectId,IntUserId,1,m.DB)
	if err!=nil {
		m.Logger.Error("Fail to insert user to teammate")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL, "insert into teammater fail", err)
	}

	return nil
}
