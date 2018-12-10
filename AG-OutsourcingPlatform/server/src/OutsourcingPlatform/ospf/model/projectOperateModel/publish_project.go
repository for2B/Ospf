package projectOperateModel

import (
	"database/sql"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
)

//获取请求参数，并将项目信息插入数据库中
func (m *ProjectOperateModel)PublishProject(ProjectName,ProjectDetail string,ProjectCircle int,ProjectMoney float64,OpenId string) error {

	//Checking if parameter are invalid
	if ProjectName  == "" {
		m.Logger.Error("ProjectName is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectName is empty",nil)
	}

	if ProjectDetail == "" {
		m.Logger.Error("ProjectDetail is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectDetail is empty",nil)
	}

	if OpenId == "" {
		m.Logger.Error("OpenId is empty")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR	,"OpenId si empty",nil)
	}

	if ProjectCircle < 0 {
		m.Logger.Error("ProjectCircle is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectCircle is invalid",nil)
	}

	if ProjectMoney < 0 {
		m.Logger.Error("ProjectMoney is invalid")
		return httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ProjectMoney is invalid",nil)
	}

	//get UserId
	var PublisherId int
	var err error
	err = m.DB.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&PublisherId)
	if err!=nil{
		if err == sql.ErrNoRows {
			m.Logger.Error("Could not find the user according to OpenId")
			return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		m.Logger.Error("Fail to QueryRow for searching user")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}

	//insert new column into project
	var ProjectId int
	SqlInert := `INSERT INTO 
					project(publisher_id,project_name,project_circle,project_detail,project_money,project_status,project_publish_time,project_alter_time) 
					values ($1,$2,$3,$4,$5,1,CURRENT_DATE,CURRENT_DATE) returning project_id; `
	err = m.DB.QueryRow(SqlInert,PublisherId,ProjectName,ProjectCircle,ProjectDetail,ProjectMoney).Scan(&ProjectId)
	if err!=nil{
		m.Logger.Error("Fail to insert new column")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to insert new column",err)
	}

	//发布项目后要在teammate 表中插入发布者与这个项目的关系记录 2为发布者
	err = InsertIntoTeammate(ProjectId,PublisherId,2,m.DB)
	if err!=nil{
		m.Logger.Error("Fail to insert new column into teammate")
		return httpapi.NewErr(constant.SQL_OPERATE_FAIL, "Fail to insert new column into teammate", err)
	}

	return nil
}

















