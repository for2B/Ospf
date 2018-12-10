package getProjectModel

import (
	"database/sql"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
)

//根据身份返回个状态的项目总数
func (m *GetProjectModel)CountCompartedStatus(ListType , OpenId string) (interface{} , error){

	//checking if parameter is empty
	if ListType == "" {
		m.Logger.Error("ListType is empty")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR	,"ListType is empty",nil)
	}

	if OpenId == "" {
		m.Logger.Error("OpenId is empty")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"OpenId is empty",nil)
	}

	//get UserId
	var UserId int64
	var err error

	err = m.DB.QueryRow(`SELECT user_id FROM users WHERE openid = $1;`,OpenId).Scan(&UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			m.Logger.Error("Could not find the user according to OpenId")
			return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Could not find the user according to OpenId",err)
		}
		m.Logger.Error("Fail to QueryRow for searching user")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to QueryRow for searching user",err) //数据库出错, 用户不会产生这种错误
	}


	switch ListType {
	case "my_publish": //获取我发布的项目各状态的总数
		MyPublishStatusCount , err := m.getMyPublish(UserId)
		if err!=nil {
			return nil,httpapi.NewErr(constant.FAIL_GET_RETDATA,"Fail to get MyPublishStatusCount",err)
		}
		return MyPublishStatusCount,nil
	case "my_contract":
		MyContractStatusCount , err := m.getMyContract(UserId)
		if err!=nil {
			return nil,httpapi.NewErr(constant.FAIL_GET_RETDATA,"Fail to get MyContractStatusCount",err)
		}
		return MyContractStatusCount,nil
	case "my_participate":
		MyParticipateStatusCount , err := m.getMyParticipate(UserId)
		if err!=nil {
			return nil,httpapi.NewErr(constant.FAIL_GET_RETDATA,"Fail to get MyParticipateStatusCount",err)
		}
		return MyParticipateStatusCount,nil
	default:
		m.Logger.Error("ListType is invalid")
		return nil,httpapi.NewErr(constant.CHECK_PARAMS_ERR,"ListType is invalid",nil)
	}

	return nil,nil
}

//返回我创建的项目各状态的总数
func (m *GetProjectModel)getMyPublish(UserId int64) (interface{},error){

	var NumberOfWaitContract,NumberOfWaitSign,NumberOfSign,NumberOfEnd,NumberOfTotal int

	SqlSelect := `SELECT COUNT (*) FROM project WHERE publisher_id = $1 AND project_status = $2`
	err := m.DB.QueryRow(SqlSelect,UserId,1).Scan(&NumberOfWaitContract)
	if err != nil {
		m.Logger.Error("Fail to get NumberOfWaitContract")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfWaitContract",err)
	}

	err = m.DB.QueryRow(SqlSelect,UserId,2).Scan(&NumberOfWaitSign)
	if err != nil {
		m.Logger.Error("Fail to get NumberOfWaitSign")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfWaitSign",err)
	}

	err = m.DB.QueryRow(SqlSelect,UserId,3).Scan(&NumberOfSign)
	if err != nil {
		m.Logger.Error("Fail to get NumberOfSign")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfSign",err)
	}

	err = m.DB.QueryRow(SqlSelect,UserId,4).Scan(&NumberOfEnd)
	if err != nil {
		m.Logger.Error("Fail to get NumberOfEnd")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfEnd",err)
	}

	NumberOfTotal = NumberOfWaitContract+NumberOfWaitSign+NumberOfSign+NumberOfEnd
	NumberOfMyPublish := NumberOfEachStatus{
		Total:NumberOfTotal,
		StatusOne:NumberOfWaitContract,
		StatusTwo:NumberOfWaitSign,
		StatusThree:NumberOfSign,
		StatusFour:NumberOfEnd,
	}

	return NumberOfMyPublish,nil

}

//返回我承接的项目各状态的总数
func (m *GetProjectModel)getMyContract(UserId int64)(interface{},error){

	var NumberOfBidding,NumberOfWaitSign,NumberOfSign,NumberOfEnd,NumberOfTotal,NumberOfBidFail int

	SqlSelectBidding := `SELECT COUNT(*) FROM project p1 ,competitors c1 WHERE p1.project_id = c1.project_id AND p1.contractor_id is null 
						AND p1.project_status = 1 AND c1.user_id = $1`

	err := m.DB.QueryRow(SqlSelectBidding,UserId).Scan(&NumberOfBidding)
	if err!=nil {
		m.Logger.Error("Fail to get NumberOfBidding")
		return nil , httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfBidding",err)
	}

	SqlSelectMyContract := `SELECT COUNT (*) FROM project WHERE contractor_id = $1 AND project_status = $2`
	err = m.DB.QueryRow(SqlSelectMyContract,UserId,2).Scan(&NumberOfWaitSign)
	if err!=nil {
		m.Logger.Error("Fail to get NumberOfWaitSign")
		return nil , httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfWaitSign",err)
	}

	err = m.DB.QueryRow(SqlSelectMyContract,UserId,3).Scan(&NumberOfSign)
	if err!=nil {
		m.Logger.Error("Fail to get NumberOfSign")
		return nil , httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfSign",err)
	}

	err = m.DB.QueryRow(SqlSelectMyContract,UserId,4).Scan(&NumberOfEnd)
	if err!=nil {
		m.Logger.Error("Fail to get NumberOfEnd")
		return nil , httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfEnd",err)
	}

	SqlBiddingFail := `SELECT COUNT (*) FROM project p1 , competitors c1 WHERE p1.project_id = c1.project_id
					   AND p1.contractor_id is not null AND p1.contractor_id != $1 AND p1.project_status != 1 AND c1.user_id = $2`
	err = m.DB.QueryRow(SqlBiddingFail,UserId,UserId).Scan(&NumberOfBidFail)
	if err!=nil {
		m.Logger.Error("Fail to get NumberOfBidFail")
		return nil , httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get NumberOfBidFail",err)
	}

	NumberOfTotal = NumberOfBidding+NumberOfWaitSign+NumberOfSign+NumberOfEnd+NumberOfBidFail
	NumberOfMyContract := NumberOfEachStatus{
		Total:NumberOfTotal,
		StatusOne:NumberOfBidding,
		StatusTwo:NumberOfWaitSign,
		StatusThree:NumberOfSign,
		StatusFour:NumberOfEnd,
		StatusFive:NumberOfBidFail}
	return NumberOfMyContract,nil
}

//返回我参与的项目各状态的总数
func (m *GetProjectModel)getMyParticipate(UserId int64)(interface{},error){
	//TODO Wait to complete
	return nil,nil
}
