package TreadPit

import "testing"

var MockLoginGodReq = map[string]string{
	"redirectURL":"/pc/project/pfprolist",
	"userType":"2",
}


var ReqParamsGetPlatListAll = `{"cycle_range_end":999999999999,"cycle_range_start":0,"date_range_end":6311433600,"date_range_start":4692528000,"money_range_end":999999999999,"money_range_start":0,"current_page":1,"show_nums":10}`
var ExpectGetPlatList = `{"code":0,"data":{"total":13,"project_list":[{"project_name":"T_TestProject1","project_id":100001,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":30,"project_detail":"无","project_money":1000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject2","project_id":100002,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":40,"project_detail":"无","project_money":5000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject3","project_id":100003,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":70,"project_detail":"无","project_money":8000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject4","project_id":100004,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":100,"project_detail":"无","project_money":1000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject5","project_id":100005,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":30,"project_detail":"无","project_money":4000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject6","project_id":100006,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":40,"project_detail":"无","project_money":9000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject7","project_id":100007,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":50,"project_detail":"无","project_money":3000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject8","project_id":100008,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":80,"project_detail":"无","project_money":10000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject9","project_id":100009,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":150,"project_detail":"无","project_money":12000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject10","project_id":100010,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":130,"project_detail":"无","project_money":17000,"project_status":1,"competitors_nums":0}]}}`

var ReqParamsGetPlatListInvalidParams = `{"cycle_range_end":0,"cycle_range_start":-1,"date_range_end":0,"date_range_start":0,"money_range_end":0,"money_range_start":0,"current_page":1,"show_nums":10}`
var ExpectGetPlatListInvalidParam = `{"code":301,"msg":"CycleFirst is invalid"}`
var ReqParamsGetPlatListFilter1 = `{"cycle_range_end":90,"cycle_range_start":30,"date_range_end":6311433600,"date_range_start":4692528000,"money_range_end":999999999999,"money_range_start":100000,"current_page":1,"show_nums":10}`
var ExpectGetPlatListFilter1 = `{"code":0,"data":{"total":3,"project_list":[{"project_name":"T_TestProject11","project_id":100011,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":30,"project_detail":"无","project_money":100000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject12","project_id":100012,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":70,"project_detail":"无","project_money":120000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject13","project_id":100013,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":80,"project_detail":"无","project_money":150000,"project_status":1,"competitors_nums":0}]}}`
var ReqParamsGetPlatListFilter2 = `{"cycle_range_end":90,"cycle_range_start":30,"date_range_end":6311433600,"date_range_start":4692528000,"money_range_end":10000,"money_range_start":0,"current_page":1,"show_nums":10}`
var ExpectGetPlatListFilter2 = `{"code":0,"data":{"total":7,"project_list":[{"project_name":"T_TestProject1","project_id":100001,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":30,"project_detail":"无","project_money":1000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject2","project_id":100002,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":40,"project_detail":"无","project_money":5000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject3","project_id":100003,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":70,"project_detail":"无","project_money":8000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject5","project_id":100005,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":30,"project_detail":"无","project_money":4000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject6","project_id":100006,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":40,"project_detail":"无","project_money":9000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject7","project_id":100007,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":50,"project_detail":"无","project_money":3000,"project_status":1,"competitors_nums":0},{"project_name":"T_TestProject8","project_id":100008,"publisher_id":27,"project_publish_time":"2118-09-14","project_circle":80,"project_detail":"无","project_money":10000,"project_status":1,"competitors_nums":0}]}}`

var mockProjectData = []TestProjectInfo{
	{100001,27,"T_TestProject1",30,1000,"无",1},
	{100002,27,"T_TestProject2",40,5000,"无",1},
	{100003,27,"T_TestProject3",70,8000,"无",1},
	{100004,27,"T_TestProject4",100,1000,"无",1},
	{100005,27,"T_TestProject5",30,4000,"无",1},
	{100006,27,"T_TestProject6",40,9000,"无",1},
	{100007,27,"T_TestProject7",50,3000,"无",1},
	{100008,27,"T_TestProject8",80,10000,"无",1},
	{100009,27,"T_TestProject9",150,12000,"无",1},
	{100010,27,"T_TestProject10",130,17000,"无",1},
	{100011,27,"T_TestProject11",30,100000,"无",1},
	{100012,27,"T_TestProject12",70,120000,"无",1},
	{100013,27,"T_TestProject13",80,150000,"无",1},
}

func InsertMockProjectData(t *testing.T,tc *TestCase){
	tx,err := TestDb.Begin()
	if err!=nil {
		tx.Rollback()
		t.Fatalf("%v Prepare mock data fail err:%v",tc.Describe,err)
		return
	}

	stmt ,err := tx.Prepare(`INSERT INTO
	project(project_id,publisher_id,project_name,project_circle,project_detail,project_money,project_status,project_publish_time,project_alter_time)
	values ($1,$2,$3,$4,$5,$6,$7,timestamp with time zone '2118-09-14 00:00:00.000000',timestamp with time zone '2118-09-14 00:00:00.000000');`)
	if err!=nil {
		t.Fatalf("%v Prepare mock data fail err:%v",tc.Describe,err)
		return
	}
	for _,insertRec := range mockProjectData{
		_,err = stmt.Exec(insertRec.Id,insertRec.PublisherId,insertRec.Name,insertRec.Circle,insertRec.Detail,insertRec.Money,insertRec.Status)
		if err!=nil {
			tx.Rollback()
			t.Fatalf("%v Prepare mock data fail err:%v",tc.Describe,err)
			return
		}
	}
	tx.Commit()
}

func ClearMockProjectData(t *testing.T,tc *TestCase){
	_, err := TestDb.Exec("delete from project  where project_name like 'T%';")
	if err!=nil {
		t.Fatalf("%v Prepare mock data fail err:%v",tc.Describe,err)
		return
	}
}

var ReqParamsGetIdentity = map[string]string{
	"project_id":"60",
}
var ExpectGetIdentity = `{"code":0,"data":{"identity":2}}`