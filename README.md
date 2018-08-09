# GO~

## 路由

### 创建路由实例

```javascript
    route := odserver.Default(
```

### 创建基本路由

```javascript
route.Target("/user").Get(SayHello).Post(SayHello2)
```

对路径`/user`同时注册了Get和Post两种请求方式

### 统一开始域

```javascript
route.Start("/cookie").
    Target("/read").Get(ReadCookieServer).And().
    Target("/write").Get(WriteCookieServer).And()
```

等价于

```javascript
route.Target("/cookie/read").Get(ReadCookieServer).And().
route.Target("/cookie/write").Get(WriteCookieServer).And()
```

### 参数获取