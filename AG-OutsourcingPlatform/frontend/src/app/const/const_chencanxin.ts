



// export const domain="http://localhost:6616"
// export const domain="http://chuil.online:6616" //开发 生成 发布均要一个
export const domain="http://chuil.online/ospf" //开发 生成 发布均要一个

export const REFEREE_QRCODE_TYPE="REFEREE"
export const NORMAL_QRCODE_TYPE="NORMAL"


//后端api
export var URL_MOCK_LOGIN=domain+"/api/mock_login"
export var URL_WECHAT_LOGIN=domain+"/api/wechat_login"
export var URL_FETCH_QRCODE=domain+"/api/fetch_invitation_code?type="
export var URL_GENERATE_GOLANG_API_CODE=domain+"/api/generateGolangApiCode"


//pc前端路由
export const PC_ROUTE = {
  mock_login: {name: "模拟登陆", url: "/mock_login"},
  myqualification:{name:"我的个人能力资质",url:"/pc/qulafication/my"},

}


//phone前端路由
export const PHONE_ROUTE = {
  mock_login: {name: "模拟登陆", url: "/student/regist"},
  qualification:{name:"我的个人能力资质",url:"/pc/"}
}


//错误码===============================================================
export const FB = {
  GLOBAL_SUCCESS               : {msg: "成功", code: 0},
  GLOBAL_SYS_ERR               : {msg: "系统出错", code: 202},
  GLOBAL_PARAM_ERR             : {msg: "请求参数错误", code: 201},
  GLOBAL_NO_PERMISSION         : {msg: "没有权限", code: 203},
  GLOBAL_DB_ERR                : {msg: "数据库操作失败", code: 204},
  GLOBAL_HTTP_METHOD_ERR       : {msg: "HTTP请求方法失败", code: 205},
  GLOBAL_ASSERT_ERR            : {msg: "断言失败", code: 206},
  GLOBAL_REQUEST_BODY_PARSE_ERR: {msg: "请求体解析错误", code: 207},

  SESSION_EXPIRED  : {msg: "session过期", code: 208},
  SESSION_FETCH_ERR: {msg: "获取session失败", code: 209},
  SESSION_SET_ERR  : {msg: "SESSION写入cache失败", code: 210},


  COOKIE_NULL      : {msg: "客户端没有cookie", code:211 },
  COOKIE_EMPTY     : {msg: "COOKIE内容为空", code:212 },
  ERROR_INVALID_SID: {msg: "sid非法(cookie值非法)", code:213 },

  QRCODE_TOKEN_INVALID: {msg: "二维码已失效", code:401 },
}





