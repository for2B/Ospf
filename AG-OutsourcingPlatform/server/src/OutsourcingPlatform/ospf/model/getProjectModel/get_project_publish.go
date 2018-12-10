package getProjectModel

import (
	"strconv"
	"database/sql"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

//根据StrProjectStatus的值返回对应的项目列表，如果值为""（对应整型为 0)，则返回所有项目状态,其他参数不能为空
func (m *GetProjectModel)GetProjectPublish(StrProjectStatus ,OpenId,StrCurrentPage,StrShowQuantity string) (interface{},error) {

	var IntProjectStatus int
	var err error

	//Checking if the parameters are empty
	if StrProjectStatus == "" {
		IntProjectStatus = 0
	}

	if OpenId == ""{
		m.Logger.Error("OpenId is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"OpenId is empty",nil)
	}

	if StrCurrentPage == ""{
		m.Logger.Error("StrCurrentPage is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrCurrentPage is empty",nil)
	}

	if StrShowQuantity == ""{
		m.Logger.Error("StrShowQuantity is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrShowQuantity is empty",nil)
	}

	//Converting parameter type
	if StrProjectStatus != "" {
		IntProjectStatus ,err = strconv.Atoi(StrProjectStatus)
		if err!=nil{
			m.Logger.Error("Fail to convert StrProjectStatus to int")
			return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProjectStatus to int",nil)
		}
	}

	IntCurrentPage, err := strconv.Atoi(StrCurrentPage)
	if err!=nil{
		m.Logger.Error("Fail to convert StrCurrentPage to int")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrCurrentPage to int",nil)
	}

	IntShowQuantity , err := strconv.Atoi(StrShowQuantity)
	if err != nil{
		m.Logger.Error("Fail to convert StrShowQuantity to int")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrShowQuantity to int",nil)
	}

	//Checking if the parameters are valid
	if IntProjectStatus < 0 || IntProjectStatus > 6 {
		m.Logger.Error("IntPeojectStatus values is not valid,it should be between 1 and 5")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntPeojectStatus values is not valid,it should be between 1 and 5",nil)
	}

	if IntCurrentPage < 0 {
		m.Logger.Error("IntCurrentPage values is not valid")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntCurrentPage values is not valid",nil)
	}

	if IntShowQuantity < 0 {
		m.Logger.Error("IntShowQuantity values is not valid")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntShowQuantity values is not valid",nil)
	}

	//get UserId
	var UserId int64

	err = m.DB.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&UserId)
	if err!=nil{
		if err == sql.ErrNoRows{
			m.Logger.Error("Could not find the user according to OpenId")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		m.Logger.Error("Fail to QueryRow for searching user")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}

	//get ProjectList
	SqlStmt := `SELECT project_name,p1.project_id,publisher_id,TO_CHAR(project_publish_time,'YYYY-MM-DD') project_publish_time,
       project_circle,project_money,project_detail,project_status, COUNT(c1.project_id) competitors_nums
		FROM project p1 LEFT JOIN competitors c1 on p1.project_id = c1.project_id 
		WHERE publisher_id = $1`
	SqlStatus := `AND project_status = $2`
	SqlGroup :=`GROUP BY p1.project_id`
	SqlLimit :=` LIMIT $2 OFFSET $3*($4-1);`
	SqlLimit2 :=` LIMIT $3 OFFSET $4*($5-1);`

	var row *sql.Rows
	switch IntProjectStatus {
	case 0:  // 获取所有状态的项目列表
		row,err = m.DB.Query(SqlStmt+SqlGroup+SqlLimit,UserId,IntShowQuantity,IntShowQuantity,IntCurrentPage)
		if err!=nil {
			m.Logger.Error("SQL query fail")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"SQL query fail",err)
		}
	default: //根据IntProjectStatus 获取对应状态的项目
		row,err = m.DB.Query(SqlStmt+SqlStatus+SqlGroup+SqlLimit2,UserId,IntProjectStatus,IntShowQuantity,IntShowQuantity,IntCurrentPage)
		if err!=nil {
			m.Logger.Error("SQL query fail")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"SQL query fail",err)
		}
	}
	defer row.Close()

	var projectList []Project
	for row.Next(){
		var oneProject Project
		err = row.Scan(
			&oneProject.ProjectName,
			&oneProject.ProjectId,
			&oneProject.PublisherId,
			&oneProject.ProjectPublicTime,
			&oneProject.ProjectCircle,
			&oneProject.ProjectMoney,
			&oneProject.ProjectDetail,
			&oneProject.ProjectStatus,
			&oneProject.CompetitorsNums)
		if err!=nil {
			m.Logger.Error("")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to scan row for oneProject",err)
		}
		projectList = append(projectList,oneProject)
	}

	return projectList,nil
}