package projectRelatedPersonModel

import (
	"strconv"
	"database/sql"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)



func (m *ProjectRelatedPersonModel) GetBiddingList(StrProjectId, StrCurrentPage, StrShowQuantity ,SortOrder string) (interface{}, error) {

	//Checking if parameter are empty
	if StrProjectId == "" {
		m.Logger.Error("StrProjectId is empty")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "StrProjectId is empty", nil)
	}

	if StrCurrentPage == "" {
		m.Logger.Error("StrCurrentPage is empty")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "StrCurrentPage is empty", nil)
	}

	if StrShowQuantity == "" {
		m.Logger.Error("StrShowQuantity is empty")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "StrShowQuantity is empty", nil)
	}

	//Converting parameter type
	IntProjectId, err := strconv.Atoi(StrProjectId)
	if err != nil {
		m.Logger.Error("Fail to convert StrProjectId to int")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "Fail convert StrProjectId to int", nil)
	}

	IntCurrentPage, err := strconv.Atoi(StrCurrentPage)
	if err != nil {
		m.Logger.Error("Fail to convert StrCurrentPage to int")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "Fail to convert StrCurrentPage to int", nil)
	}

	IntShowQuantity, err := strconv.Atoi(StrShowQuantity)
	if err != nil {
		m.Logger.Error("Fail to convert StrShowQuantity to int")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "Fail to convert StrShowQuantity to int", nil)
	}

	//Checking if parameter is valid
	if IntProjectId < 0 {
		m.Logger.Error("IntProjectId is invalid")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "IntProjectId is invalid", nil)
	}

	if IntCurrentPage <= 0 {
		m.Logger.Error("IntCurrentPage is invalid")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "IntCurrentPage is invalid", nil)
	}

	if IntShowQuantity <= 0 {
		m.Logger.Error("IntShowQuantity is invalid")
		return nil, httpapi.NewErr(constant.CHECK_PARAMS_ERR, "IntShowQuantity is invalid", nil)
	}

	if SortOrder != "descend" && SortOrder != "ascend" && SortOrder!=""{
		m.Logger.Error("SortOrder is invalid")
		return nil , httpapi.NewErr(constant.CHECK_PARAMS_ERR,"SortOrder is invalid",nil)
	}


	//get ContractorId
	//var ContractorId int
	//var isNullId sql.NullInt64
	//err = m.DB.QueryRow(`SELECT contractor_id FROM project WHERE project_id = $1;`, IntProjectId).Scan(&isNullId)
	//if err != nil {
	//	m.Logger.Error("ERR:There is no such project")
	//	return nil, httpapi.NewErr(constant.SQL_OPERATE_FAIL, "ERR:There is no such project", err)
	//}

	//if !isNullId.Valid {
	//	ContractorId = 0
	//}

	//get BiddingList
	var Bidders []Bidder
	var SqlStmt string
	var Contractor Bidder

	//找出承接者，放在第一位,如果在下面的for.scan中去找出竞标者的话，因为输出条数限制和排序，如果在第一页的数据里没有承包者那就不能保证第一行一定是承接者了
	//而且承接者只显示在第一页第一行
	if IntCurrentPage == 1{
		err = m.DB.QueryRow(`SELECT u1.user_id,u1.user_name,head_img,expertise_field,address,COUNT(p1.contractor_id) 
								FROM competitors c1
                				LEFT JOIN project p on c1.project_id = p.project_id
                				LEFT JOIN project p1 on p1.contractor_id = c1.user_id AND p1.project_status = 4
                				LEFT JOIN users u1 on c1.user_id = u1.user_id
                				WHERE c1.project_id = $1 AND c1.user_id = p.contractor_id GROUP BY u1.user_id;`,IntProjectId).
			Scan(&Contractor.UserId,&Contractor.Name,&Contractor.HeadImg,&Contractor.ExpertiseField,&Contractor.Address,&Contractor.FinishProjectCount)
		if err!=nil {
			if err == sql.ErrNoRows {
				Contractor.UserId = 0
			}else{
				m.Logger.Error("Fail to select contractor")
				return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to select contractor",nil)
			}
		}else{
			Contractor.Identity = 1
			Bidders = append(Bidders,Contractor)
		}
	}

	//降序输出
	if SortOrder == "descend" {
		SqlStmt = `SELECT u1.user_id,u1.user_name,head_img,expertise_field,address,COUNT(p1.project_id) 
		FROM competitors c1 
		LEFT JOIN project p1 ON  p1.contractor_id = c1.user_id  AND p1.project_status = 4
		LEFT JOIN users u1 ON c1.user_id = u1.user_id
		WHERE  c1.project_id = $1 GROUP by u1.user_id ORDER BY COUNT(p1.project_id) desc 
		LIMIT $2 OFFSET $3*($4-1);`
	}
	//升序输出
	if SortOrder == "ascend" {
		SqlStmt = `SELECT u1.user_id,u1.user_name,head_img,expertise_field,address,COUNT(p1.project_id) 
		FROM competitors c1 
		LEFT JOIN project p1 ON  p1.contractor_id = c1.user_id  AND p1.project_status = 4
		LEFT JOIN users u1 ON c1.user_id = u1.user_id
		WHERE  c1.project_id = $1 GROUP by u1.user_id ORDER BY COUNT(p1.project_id) asc 
		LIMIT $2 OFFSET $3*($4-1);`
	}
	//无排序输出
	if SortOrder == "" {
		SqlStmt = `SELECT u1.user_id,u1.user_name,head_img,expertise_field,address,COUNT(p1.project_id) 
		FROM competitors c1 
		LEFT JOIN project p1 ON  p1.contractor_id = c1.user_id  AND p1.project_status = 4
		LEFT JOIN users u1 ON c1.user_id = u1.user_id
		WHERE  c1.project_id = $1 GROUP by u1.user_id LIMIT $2 OFFSET $3*($4-1);`
	}

	row, err := m.DB.Query(SqlStmt, IntProjectId, IntShowQuantity, IntShowQuantity, IntCurrentPage)
	defer row.Close()
	if err != nil {
		m.Logger.Error("Fail to query")
		return nil, httpapi.NewErr(constant.SQL_OPERATE_FAIL, "Fail to query", err)
	}

	for row.Next() {

		var bidder Bidder
		err = row.Scan(
			&bidder.UserId,
			&bidder.Name,
			&bidder.HeadImg,
			&bidder.ExpertiseField,
			&bidder.Address,
			&bidder.FinishProjectCount)
		if err != nil {
			m.Logger.Error("Fail to scan for Bidder")
			return nil, httpapi.NewErr(constant.SQL_OPERATE_FAIL, "Fail to scan for Bidder", err)
		}
		bidder.Identity = 0
		//如果该Bidder为项目承包者，则将Identity设置为1，否则设置为0
		//并且在第一页时需要将其放在第一位
		if IntCurrentPage == 1 {
			if bidder.UserId != Contractor.UserId {
				Bidders = append(Bidders, bidder)
			}
		}else{
			Bidders = append(Bidders, bidder)
		}
	}

	//返回 BidderList
	var bidderList BidderList
	bidderList.Bidders = Bidders

	//竞标总人数
	bidderList.Bidders = Bidders
	err = m.DB.QueryRow(`SELECT COUNT(*) FROM competitors WHERE project_id = $1`,IntProjectId).Scan(&bidderList.Total)
	if err!=nil {
		m.Logger.Error("Fail to get total of Bidders")
		return nil,httpapi.NewErr(constant.SQL_OPERATE_FAIL,"Fail to get total of Bidders",err)
	}

	return bidderList, nil
}

