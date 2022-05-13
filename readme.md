# request
磐石开发框架默认http请求模块

## 函数说明
### http.client
- 参数为空默认创建一个请求的客户端,可以复用
```lua
    local client = http.client()
```

### http.client.R
- 创建一个请求对象
```lua
   local r = http.client().R() 
   local resp = r.GET("http://www.a.com")
```
### http.client().*
### http.client().R().*
- 发起请求或者设置请求参数
- R.GET
- R.HEAD
- R.POST
- R.PUT
- R.PATCH
- R.DELETE
- R.CONNECT
- R.OPTIONS
- R.TRACE
- R.POST

```lua
    local r  = http.client().R()
    local r1 = r.GET("http://www.a.com")
    local r2 = r.POST("http://www.a.com/api")
```

### http.client.R.output
- 这是个一个实现下载的库,可以用保存请求结果到文件
```lua
    http.client().R().output("1.html").GET("http://www.a.com/1.html")
```

## response
这个是返回对象的
- response.ERR 
- response.code
- response.body
- response.header
```lua
    local r = http.client().R().GET("http://www.a.com/1.html")
    print(r.ERR)
    print(r.code)
    print(r.body)
    print(r.header["content-length"])
```
