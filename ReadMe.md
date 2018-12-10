# 前端框架说明

1, baseService包含有吐司,模态框,title(设置页面标题)等基础服务;

2, http拦截器具有以下功能:
- 统一检查了后端返回response结构体是否undefined和是否为null.
- response结构体中是否含有code,因此具体业务中不必重复检查.
- 如果code不为0(即后端接口有返回错误),拦截器将弹出模态框说明错误码和具体信息.

3, 如果请求的接口必定有返回data属性,可以调用`GetDataByHttpGet(url: string, _next:(data)=>void)`和`GetDataByPost(_url: string, _body: object,_next:(data)=>void)`,省去检查data属性是否为空的代码

4,[localhost:4200/super/buildCode](localhost:4200/super/buildCode) 页面可以自动生成后端新接口代码

# 后端框架说明

1, post请求的参数使用`httpparse.GetBody(r  *http.Request)`获取.

```
	//demo for get api param by http post
	getbody, err := httpparse.GetBody(r)
	if err != nil {
		return nil, httpapi.Err{Code: constant.HTTP_ERR_READBODY, Msg: "get body param failed"}
	}
```

2, url参数使用`httpparse.GetForm(r  *http.Request)`获取.

3, 当前用户的Openid使用`httpparse.GetOpenid(r *http.Request)`获取.

**(错误请指出并及时改正,更多需补充的事项请更新并通知团队)**