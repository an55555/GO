# GO~

## 路由

### 创建路由实例

```javascript
    route := odserver.Default(
```

### 路由方式一：

### 创建基本路由

```javascript
route.Get("/user",SayHello)
```

支持链式

```javascript
route.Get("/user",SayHello).Get("/user1",SayHello)
```

#### 统一开始域

```javascript
route.Start("/new").Get("/1", SayHello).Get("/2", SayHello)
```

### 路由方式二

#### 创建基本路由

```javascript
route.Target("/user").AddGet(SayHello).AddPost(SayHello2)
```

对路径`/user`同时注册了Get和Post两种请求方式

#### 统一开始域

```javascript
route.Start("/cookie").
    Target("/read").GoGet(ReadCookieServer).GoPost(ReadCookieServer).And().
    Target("/write").GoGet(WriteCookieServer)
```

等价于

```javascript
route.Target("/cookie/read").GoGet(ReadCookieServer)
route.Target("/cookie/read").GoGet(ReadCookieServer)
route.Target("/cookie/write").GoPost(WriteCookieServer)
```

### 参数获取

#### URL请求参数

`route.Target("/params/{id}").Get(GetParams)`

```javascript
func GetParams(c *odserver.Context) {
	fmt.Println(c.Params)
}
```