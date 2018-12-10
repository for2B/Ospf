package getProjectModel

import (
	"strconv"
	"database/sql"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

//根据StrProjectStatus的值返回对应的项目列表，如果值为""（对应整型为 0)，则返回所有项目状态,其他参数不能为空
func (m *GetProjectModel)GetProjectContract(StrProjectStatus,OpenId,StrCurrentPage,StrShowQuantity string) ([]Project,error) {

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

	if StrCurrentPage == "" {
		m.Logger.Error("StrCurrentPage is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrCurrentPage is empty",nil)
	}

	if StrShowQuantity == "" {
		m.Logger.Error("StrShowQuantity is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"StrShowQuantity is empty",nil)
	}

	//Converting parameter type
	if StrProjectStatus != "" {
		IntProjectStatus , err = strconv.Atoi(StrProjectStatus)
		if err!=nil{
			m.Logger.Error("Fail to convert StrProjectStatus to int")
			return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrProjectStatus to int",nil)
		}
	}

	IntCurrentPage , err := strconv.Atoi(StrCurrentPage)
	if err!=nil {
		m.Logger.Error("Fail to convert StrCurrentPage to int")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrCurrentPage to int",nil)
	}

	IntShowQuantity , err := strconv.Atoi(StrShowQuantity)
	if err != nil {
		m.Logger.Error("Fail to convert StrShowQuantity to int")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"Fail to convert StrShowQuantity to int",nil)
	}

	//Checking if the parameters are valid
	if IntProjectStatus < 0 || IntProjectStatus > 6 {
		m.Logger.Error("IntPeojectStatus values is invalid,it should be between 1 and 5")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntPeojectStatus values is invalid,it should be between 1 and 5",nil)
	}

	if IntCurrentPage < 0 {
		m.Logger.Error("IntCurrentPage values is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntCurrentPage values is invalid",nil)
	}

	if IntShowQuantity < 0 {
		m.Logger.Error("IntShowQuantity values is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntShowQuantity values is invalid",nil)
	}

	//get UserId
	var UserId int64

	err = m.DB.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			m.Logger.Error("Could not find the user according to OpenId")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		m.Logger.Error("Fail to QueryRow for searching user")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}

	//get ProjectList
	var ProjectList []Project

	switch IntProjectStatus {
	case 0: //获取所有状态的项目
		ProjectList,err = m.getAllList(UserId,IntCurrentPage,IntShowQuantity)
		if err!=nil{
			return nil,err
		}
	case 1: //获取项目状态为‘参与竞标’的项目
		WaitContractorList ,err := m.getBiddingList(UserId,IntCurrentPage,IntShowQuantity)
		if err!=nil {
			return nil,err
		}
		ProjectList = append(ProjectList,WaitContractorList...)
	case 2,3,4: //获取项目状态为‘等待签订合同’、‘签订合同’、‘项目结束’的项目列表
		WaitSignEndList , err := m.getWaitSignEndList(UserId,IntProjectStatus,IntCurrentPage,IntShowQuantity)
		if err!=nil {
			return nil,err
		}
		ProjectList = append(ProjectList,WaitSignEndList...)
	case 5://获取‘竞标失败'的项目
		BidFailList,err := m.getBidFailList(UserId,IntCurrentPage,IntShowQuantity)
		if err!=nil {
			return nil,err
		}
		ProjectList = append(ProjectList,BidFailList...)
	default:
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"IntProjectStatus is invalid",nil)
	}

	return ProjectList,nil

}
//获取参与竞标的项目列表
func (m *GetProjectModel)getBiddingList(UserId int64 , CurrentPage , ShowQuantity int) ([]Project,error){
	var row *sql.Rows

	row,err := m.DB.Query(`
		SELECT project_name,p1.project_id,publisher_id,TO_CHAR(project_publish_time,'YYYY-MM-DD') project_publish_time,
        project_circle,project_money,project_detail,project_status, COUNT(c1.project_id) competitors_nums
		FROM project p1 LEFT JOIN competitors c1 ON p1.project_id = c1.project_id
		WHERE p1.contractor_id is null AND p1.project_status = 1 AND c1.user_id = $1
		GROUP BY p1.project_id LIMIT $2 OFFSET $3*($4-1);	`,UserId,ShowQuantity,ShowQuantity,CurrentPage)

	if err!=nil {
		m.Logger.Error("SQL query fail")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"SQL query fail",err)
	}

	defer row.Close()

	var ProjectList []Project

	for row.Next() {
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
			&oneProject.CompetitorsNums)//获取，返回都要检查
		if err != nil {
			m.Logger.Error("Fail to scan row for oneProject")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to scan row for oneProject",err)
		}
		ProjectList = append(ProjectList,oneProject)
	}
	return ProjectList,nil
}

//根据不同的状态值获取对应的项目，状态0表示获取等待签订合同，签订合同和项目结束的活动
func (m *GetProjectModel)getWaitSignEndList(UserId int64 , ProjectStatus , CurrentPage ,ShowQuantity int ) ([]Project,error){

	SqlStmt := `SELECT project_name,p1.project_id,publisher_id,TO_CHAR(project_publish_time,'YYYY-MM-DD') project_publish_time,
        project_circle,project_money,project_detail,project_status, COUNT(c1.project_id) competitors_nums
		FROM project p1 LEFT JOIN competitors c1 on p1.project_id = c1.project_id 
		WHERE contractor_id = $1`
	SqlGroup :=` GROUP BY p1.project_id `

	var row *sql.Rows
	var err error

	switch ProjectStatus {
	case 0:
		SqlStatus := ` AND project_status != 1`
		SqlLimit :=` LIMIT $2 OFFSET $3*($4-1)`
		row,err = m.DB.Query(SqlStmt+SqlStatus+SqlGroup+SqlLimit,UserId,ShowQuantity,ShowQuantity,CurrentPage)
		if err != nil {
			m.Logger.Error("SQL query fail")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"SQL query fail",err)
		}
	case 2,3,4:
		SqlStatus := ` AND project_status != 1 AND project_status = $2`
		SqlLimit :=` LIMIT $3 OFFSET $4*($5-1)`
		row,err = m.DB.Query(SqlStmt+SqlStatus+SqlGroup+SqlLimit,UserId,ProjectStatus,ShowQuantity,ShowQuantity,CurrentPage)
		if err!=nil {
			m.Logger.Error("SQL query fail")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"SQL query fail",err)
		}
	default:
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"query project for contractor by status fail",err)
	}
	defer row.Close()

	var ProjectList []Project

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
			m.Logger.Error("Fail to scan row for oneProject")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to scan row for oneProject",err)
		}
		ProjectList = append(ProjectList,oneProject)
	}
	return ProjectList,nil
}

//连接竞标表和项目表来获取竞标失败的活动
func (m *GetProjectModel)getBidFailList(UserId int64 , CurrentPage ,ShowQuantity int) ([]Project,error){
	var row *sql.Rows
	row,err := m.DB.Query(`
		SELECT project_name,p1.project_id,publisher_id,TO_CHAR(project_publish_time,'YYYY-MM-DD') project_publish_time,
     	  project_circle,project_money,project_detail,project_status, COUNT(c1.project_id) competitors_nums
		FROM project p1 LEFT JOIN competitors c1 on p1.project_id = c1.project_id
		WHERE contractor_id != $1 AND project_status != 1 AND c1.user_id = $2
 		GROUP BY p1.project_id LIMIT $3 OFFSET $4*($5-1);
		 	`,UserId,UserId,ShowQuantity,ShowQuantity,CurrentPage)
	if err!=nil {
		m.Logger.Error("SQL query fail")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"SQL query fail",err)
	}
	defer row.Close()

	var ProjectList []Project

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
			m.Logger.Error("Fail to scan row for oneProject")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to scan row for oneProject",err)
		}
		oneProject.ProjectStatus = 5 //查出来的是竞标失败的活动列表，对应的状态设置为5
		ProjectList = append(ProjectList,oneProject)
	}
	return ProjectList,nil
}

//分别调用三个函数来获取所有状态的项目
func (m *GetProjectModel)getAllList(UserId int64 ,IntCurrentPage,IntShowQuantity int)([]Project,error){

	var AllList []Project

	WaitContractorList ,err := m.getBiddingList(UserId,IntCurrentPage,IntShowQuantity)
	if err!=nil {
		return nil,err
	}
	AllList = append(AllList,WaitContractorList...)

	WaitSignEndList , err := m.getWaitSignEndList(UserId,0,IntCurrentPage,IntShowQuantity)
	if err!=nil {
		return nil,err
	}
	AllList = append(AllList,WaitSignEndList...)

	BidFailList,err := m.getBidFailList(UserId,IntCurrentPage,IntShowQuantity)
	if err!=nil {
		return nil,err
	}
	AllList = append(AllList,BidFailList...)

	return AllList,nil
}
