package constant

//常量================================================================
const (
	RE_PROXY_SCHEME = "/xxxxx/"
)

//用户等级
const (
	USER_VISITOR = "VISITOR" //游客
	USER_REFEREE = "REFEREE" //证明人
	USER_NORMAL  = "NORMAL"  //普通用户
	USER_ADMIN   = "ADMIN"   //管理员
	USER_SUPER   = "SUPER"   //上帝
)

//构建类型
const (
	BUILD_TYPE_DEPLOY = "DEPLOY"
)

//错误码===============================================================

//全局错误码
const (
	GLOBAL_SUCCESS                = 0 //成功
	GLOBAL_SYS_ERR                = 202 //系统出错
	GLOBAL_PARAM_ERR              = 201 //请求参数错误
	GLOBAL_NO_PERMISSION          = 203 //没有权限
	GLOBAL_DB_ERR                 = 204 //数据库操作失败
	GLOBAL_HTTP_METHOD_ERR        = 205 //HTTP请求方法失败
	GLOBAL_ASSERT_ERR             = 206 //断言失败
	GLOBAL_REQUEST_BODY_PARSE_ERR = 207 //请求体解析错误
)

//SESSION相关错误码
const (
	SESSION_EXPIRED   = 208 //session过期
	SESSION_FETCH_ERR = 209 //获取session失败
	SESSION_SET_ERR   = 210 //SESSION写入cache失败

)

//COOKIE相关错误码
const (
	COOKIE_NULL    = 211 //客户端没有cookie
	COOKIE_EMPTY   = 212 //COOKIE内容为空
	ERROR_INVALID_SID = 213 //sid非法(cookie值非法)
)

//权限列表
const (
	AC_LOGIN = 1<<iota     // 允许登录
	AC_Undertake_project	// 允许承接项目
	AC_Release_project		// 允许发布项目
	AC_COMMENT			// 允许评论
	AC_ALL  =  AC_LOGIN+AC_Undertake_project+AC_Release_project+AC_COMMENT	  //合集  1111
)


