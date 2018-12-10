package projectOperateModel

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

//获取项目详情参数，并更新项目数据
func (m *ProjectOperateModel)ModifyProject(ProjectName,ProjectDetail string,ProjectId,ProjectCircle int,ProjectMoney float64) error{

	//Checking if parameter are invalid
	if ProjectName  == "" {
		m.Logger.Error("ProjectName is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectName is empty",nil)
	}

	if ProjectDetail == "" {
		m.Logger.Error("ProjectDetail is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectDetail is empty",nil)
	}

	if ProjectId <= 0 {
		m.Logger.Error("ProjectId is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectId is invalid",nil)
	}

	if ProjectCircle <= 0 {
		m.Logger.Error("ProjectCircle is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectCircle is invalid",nil)
	}

	if ProjectMoney <= 0 {
		m.Logger.Error("ProjectMoney is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectMoney is invalid",nil)
	}

	//update project
	stmt,err := m.DB.Prepare( `UPDATE project 
		SET project_name = $1,project_circle = $2 ,project_money = $3,project_detail= $4,project_alter_time = CURRENT_DATE 
		WHERE project_id = $5 AND project_status = 1;`)
	defer stmt.Close()
	if err!=nil {
		m.Logger.Error("Fail to get stmt")
		return httpapi.NewErr(constant.DB_BEGIN_FAIL,"Fail to get stmt",err)
	}
	_,err = stmt.Exec(ProjectName,ProjectCircle,ProjectMoney,ProjectDetail,ProjectId)
	if err!=nil{
		m.Logger.Error("Fail to exec update project")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to exec update project",err)
	}

	return nil
}



















