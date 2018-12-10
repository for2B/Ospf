package getProjectModel

import (
	"database/sql"
	"go.uber.org/zap"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

//根据请求参数返回平台项目列表
func (m *GetProjectModel)GetPlatformProjectList(CycleFirst,CycleLast,MoneyRangeFirst,MoneyRangeLast,DateRangeFirst,DateRangeLast,CurrentPage,ShowQuantity int,db *sql.DB,Logger *zap.Logger,OpenId string) (interface{},error) {

	//Checking if parameters ate valid
	if CycleFirst < 0 {
		Logger.Error("CycleFirst is invalid")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"CycleFirst is invalid",nil)
	}

	if CycleLast < CycleFirst {
		Logger.Error("CycleLast is less than CycleFirst")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"CycleLast is less then CycleFirst",nil)
	}

	if MoneyRangeFirst < 0 {
		Logger.Error("MoneyRangeFirst is invalid")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"MoneyRangeFirst is invalid",nil)
	}

	if MoneyRangeLast < MoneyRangeFirst {
		Logger.Error("MoneyRangeLast is less then MoneyRangeFirst")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"MoneyRangeLast is less then MoneyRangeFirst",nil)
	}

	if DateRangeFirst < 0 {
		Logger.Error("DateRangeFirst is invalid")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"DateRangeFirst is invalid",nil)
	}

	if DateRangeLast < DateRangeFirst {
		Logger.Error("DateRangeLast is less than DateRangeFirst")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"DataRangeLast is less than DateRangeFirst",nil)
	}

	if CurrentPage < 0 {
		Logger.Error("CurrentPage is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"CurrentPage is invalid",nil)
	}

	if ShowQuantity < 0 {
		Logger.Error("ShowQuantity is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ShowQuantity is invalid",nil)
	}


	//get UserId
	var UserId int
	err := db.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			Logger.Error("Could not find the user according to OpenId")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		Logger.Error("Fail to QueryRow for searching user")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}



	// get ProjectList
	row , err := db.Query(`
		SELECT project_name,p1.project_id,publisher_id,TO_CHAR(project_publish_time,'YYYY-MM-DD') project_publish_time,
			project_circle,project_money,project_detail,project_status,COUNT(c1.project_id) competitors_nums
		FROM project p1 LEFT JOIN competitors c1 ON p1.project_id = c1.project_id
		WHERE (EXTRACT(EPOCH FROM project_publish_time)-8*3600) >= $1 AND (EXTRACT(EPOCH FROM project_publish_time)-8*3600)<= $2 
		AND project_circle>=$3 AND project_circle<=$4 
		AND project_money>=$5 AND project_money<=$6 AND project_status = 1
		AND p1.project_id not in (select c2.project_id FROM competitors c2 WHERE c2.user_id = $7 AND c2.project_id = p1.project_id)
		GROUP BY p1.project_id  ORDER BY p1.project_id asc
		LIMIT $8
		OFFSET $9*($10-1);`,DateRangeFirst,DateRangeLast,CycleFirst,CycleLast,MoneyRangeFirst,MoneyRangeLast,UserId,ShowQuantity,ShowQuantity,CurrentPage)
	if err!=nil {
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"SQL query fail",err)
	}
	defer row.Close()
	var projects []Project
	for row.Next(){
		var aProject Project
		row.Scan(
			&aProject.ProjectName,
			&aProject.ProjectId,
			&aProject.PublisherId,
			&aProject.ProjectPublicTime,
			&aProject.ProjectCircle,
			&aProject.ProjectMoney,
			&aProject.ProjectDetail,
			&aProject.ProjectStatus,
			&aProject.CompetitorsNums)
		projects = append(projects,aProject)
	}

	var ProjectList ProjectList
	ProjectList.Projects = projects

	//get project total
	err=db.QueryRow(`
		SELECT COUNT(*) FROM project p1
		 WHERE (EXTRACT(EPOCH FROM project_publish_time)-8*3600) >= $1 AND (EXTRACT(EPOCH FROM project_publish_time)-8*3600)<= $2 
		AND project_circle>=$3 AND project_circle<=$4 AND project_status = 1
		AND project_money>=$5 AND project_money<=$6
		AND p1.project_id not in (select c2.project_id FROM competitors c2 WHERE c2.user_id = $7 AND c2.project_id = p1.project_id); `,
		DateRangeFirst,DateRangeLast,CycleFirst,CycleLast,MoneyRangeFirst,MoneyRangeLast,UserId).Scan(&ProjectList.Total)
	if err!=nil{
		Logger.Error("Fail to get project total")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get project total",err)
	}

	return ProjectList,nil
}
