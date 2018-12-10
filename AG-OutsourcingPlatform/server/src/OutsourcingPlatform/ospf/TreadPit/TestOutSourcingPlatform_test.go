package TreadPit

import (
	"testing"
	"fmt"
	"time"
	"net/http"
	"log"
	"os"
	"github.com/valyala/fasthttp"
	"net/url"
)

type TestDbData func(*testing.T,*TestCase)

type TestCase struct {
	Protocol 	string //Url协议
	Host 		string //请求域名+端口
	Api 		string //调用的api
	Method 		string //请求方法
	Describe 	string
	ExpectResp 	string //期望返回值
	ReqParsMsp 	map[string]string //请求参数
	ReqParsJson string //请求参数
	InsertMockData TestDbData //往数据库里添加测试数据
	ClearMockData TestDbData //清楚数据库中的测试数据
	UserType 	string //
}

const srv = "127.0.0.1:6616"
const reTryTimes = 10

var TestCookie *fasthttp.Cookie

func TestMain(m *testing.M) {
	serverOk := make(chan int)
	url := fmt.Sprintf("http://%s/api/get_project_list", srv)
	go func() {
		reTry := 0
		for {
			if reTry >= reTryTimes {
				panic(m)
			}
			time.Sleep(2 * time.Second)
			resp, err := http.Get(url)
			if err != nil {
				log.Println(fmt.Sprintf("retry...%d", reTry))
				reTry++
				continue
			}
			if resp.StatusCode != 200 {
				log.Println(fmt.Sprintf("retry...%d", reTry))
				reTry++
				continue
			}
			serverOk <- 1
		}
	}()
	<-serverOk
	log.Println("Server Ready")
	InitTestDb()
	TestCookie = &fasthttp.Cookie{}

	ExitCode := m.Run()

	ClearTestData()
	if ExitCode!=0 {
		fmt.Println("test Fail")
	}else {
		fmt.Println("test Success!")
	}
	os.Exit(ExitCode)

}

func ReqTest(t *testing.T,testCase *TestCase,client *fasthttp.Client){

	if testCase.Protocol == ""{
		testCase.Protocol = "HTTP"
	}

	if testCase.Host == "" {
		testCase.Host = srv
	}

	if testCase.InsertMockData !=nil {
		testCase.InsertMockData(t,testCase)
	}

	if testCase.ClearMockData !=nil {
		defer testCase.ClearMockData(t,testCase)
	}


	var err error

	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}

	if len(TestCookie.Value())!=0 && len(TestCookie.Key())!=0 {
		req.Header.SetCookie(string(TestCookie.Key()),string(TestCookie.Value()))
	}

	ParseUrl,err := url.Parse(testCase.Protocol+"://"+testCase.Host+testCase.Api)
	if err!=nil{
		t.Errorf("%s:err:%s",testCase.Describe,err.Error())
		return
	}

	switch testCase.Method {
	case "GET":

		//Method Get Params
		Queries := url.Values{}
		for k,v := range testCase.ReqParsMsp{
			Queries.Add(k,v)
		}
		ParseUrl.RawQuery = Queries.Encode()

		url := ParseUrl.String()
		req.SetRequestURI(url)
		req.Header.SetMethod(http.MethodGet)
		err = client.Do(req,resp)
		break
	case "POST":
		//PostJson数据
		req.Header.Set("Content-Type","application/json")
		body := testCase.ReqParsJson

		if body == "" {
			t.Errorf("%s:err,%s",testCase.Describe,"ReqParsJson empty")
			return
		}

		req.SetRequestURI(ParseUrl.String())
		req.Header.SetMethod(http.MethodPost)
		req.SetBodyString(body)
		err = client.Do(req,resp)

		break
	default:
		t.Error(testCase.Describe + "unknown request method: " + testCase.Method)
		return
		break
	}

	if err!=nil {
		t.Fatalf("%v fail,client.Do fail,err:%v",testCase.Describe,err)
	}

	if testCase.ExpectResp !="" && testCase.ExpectResp != string(resp.Body()){
		t.Fatalf("%v fail,expect %v,but get %v",testCase.Describe,testCase.ExpectResp,string(resp.Body()))
	}
}

func LoginTest(t *testing.T,testCase *TestCase,client *fasthttp.Client){
	if testCase.Protocol == ""{
		testCase.Protocol = "HTTP"
	}

	if testCase.Host == "" {
		testCase.Host = srv
	}
	if testCase.InsertMockData !=nil {
		testCase.InsertMockData(t,testCase)
	}

	if testCase.ClearMockData !=nil {
		defer testCase.ClearMockData(t,testCase)
	}
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}
	ParseUrl,err := url.Parse(testCase.Protocol+"://"+testCase.Host+testCase.Api)
	if err!=nil{
		t.Errorf("%s:err:%s",testCase.Describe,err.Error())
		return
	}
	switch testCase.Method {
	case "GET":
		//Method Get Params
		Queries := url.Values{}
		for k,v := range testCase.ReqParsMsp{
			Queries.Add(k,v)
		}
		ParseUrl.RawQuery = Queries.Encode()
		url := ParseUrl.String()
		req.SetRequestURI(url)
		req.Header.SetMethod(http.MethodGet)
		err = client.Do(req,resp)
		break
	}
	if err!=nil {
		t.Fatal()
	}

	if testCase.ExpectResp !="" && testCase.ExpectResp != string(resp.Body()){
		t.Fatalf("%v fail,expect %v,but get %v",testCase.Describe,testCase.ExpectResp,string(resp.Body()))
	}
	TestCookie.SetKey("ag_outsourcingplatform")
	if !resp.Header.Cookie(TestCookie){
		t.Fatalf("TestLogin fail,because get cookie:%v fail",TestCookie.Key())
	}

}

func TestProjectManage(t *testing.T){
	client := &fasthttp.Client{}
	var testCase *TestCase

	//测试(模拟)登陆
	testCase = &TestCase{
		Api:"/api/mock_login",
		Method:http.MethodGet,
		Describe:"test mock login , get cookie ,set session ",
		ReqParsMsp:MockLoginGodReq,
		ExpectResp:"",
	}
	LoginTest(t,testCase,client)

	//测试获取所有等待承接项目，第一页前10条
	testCase = &TestCase{
		Api:"/api/get_project_list",
		Method:http.MethodPost,
		Describe:"test get all project list for palt,should return 10 records",
		InsertMockData:InsertMockProjectData,
		ClearMockData:nil,
		ReqParsJson:ReqParamsGetPlatListAll,
		ExpectResp:ExpectGetPlatList,
	}
	ReqTest(t,testCase,client)

	//测试获取周期30~90天内，金额10w以上的平台项目
	testCase = &TestCase{
		Api:"/api/get_project_list",
		Method:http.MethodPost,
		Describe:"test filter to get project list with cycle 30~90 money > 10w",
		ReqParsJson:ReqParamsGetPlatListFilter1,
		ExpectResp:ExpectGetPlatListFilter1,
	}
	ReqTest(t,testCase,client)

	//测试获取周期30~90天内，金额1w以内的平台项目
	testCase = &TestCase{
		Api:"/api/get_project_list",
		Method:http.MethodPost,
		Describe:"test filter to get project list whit cycle 30~90 and money < 1w",
		ReqParsJson:ReqParamsGetPlatListFilter2,
		ExpectResp:ExpectGetPlatListFilter2,
		ClearMockData:nil,
	}
	ReqTest(t,testCase,client)

	//如果请求的参数非法呢
	testCase = &TestCase{
		Api:"/api/get_project_list",
		Method:http.MethodPost,
		Describe:"test empty prams",
		ReqParsJson:ReqParamsGetPlatListInvalidParams,
		ExpectResp:ExpectGetPlatListInvalidParam,
	}
	ReqTest(t,testCase,client)
}
