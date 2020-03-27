# web 开发
https://www.jishuchi.com/read/GO/701
```text
只要你的路由带有参数，并且这个项目的 api 数目超过了 10，就尽量不要使用 net/http 中默认的路由。
在Go开源界应用最广泛的 router 是 httpRouter，很多开源的 router 框架都是基于 httpRouter 
进行一定程度的改造的成果。
```
### 开源框架
```text
开源界有这么几种框架，第一种是对 httpRouter 进行简单的封装，然后提供定制的 middleware 和一些简单的小工具
集成比如 gin，主打轻量，易学，高性能。第二种是借鉴其它语言的编程风格的一些 MVC 类框架，例如 beego，
方便从其它语言迁移过来的程序员快速上手，快速开发。还有一些框架功能更为强大，除了 db 设计，大部分代码直接生成，例如 goa。
```

### restful --http method
在 restful 中除了 GET 和 POST 之外，还使用了 http 协议定义的几种其它的标准化语义<br/>
restful 风格的 API 重度依赖请求路径。会将很多参数放在请求 URI 中。
```text
const (
    MethodGet     = "GET"
    MethodHead    = "HEAD"
    MethodPost    = "POST"
    MethodPut     = "PUT"
    MethodPatch   = "PATCH" // RFC 5789
    MethodDelete  = "DELETE"
    MethodConnect = "CONNECT"
    MethodOptions = "OPTIONS"
    MethodTrace   = "TRACE"
)
```

### restful 中常见的请求路径
 ```text
GET     /repos/:owner/:repo/comments/:id/reactions
POST    /projects/:project_id/columns
PUT     /user/starred/:owner/:repo
DELETE  /user/starred/:owner/:repo
```

### web 架构
```text
01httprouter                    请求路由
02middleware                    中间件         中间件仓库//https://github.com/gin-gonic/contrib
03vaildator                     请求校验       validator库：https://github.com/go-playground/validator
04db                            与数据库打交道  database/sql库:http://go-database-sql.org/
05ratelimit                     服务流量限制
06layout                        常见大型 web 项目分层  gorm+grpc+gateway
07table-driven                  表驱动开发
08servicediscovery              服务发现
```