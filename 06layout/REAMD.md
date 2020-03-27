# layout常见大型 web 项目分层
https://www.jishuchi.com/read/GO/707

### web 框架
```text
web 框架大多数是 MVC 框架:
1.控制器（Controller）- 负责转发请求，对请求进行处理。
2.视图（View） - 界面设计人员进行图形界面设计。
3.模型（Model） - 程序员编写程序应有的功能（实现算法等等）、数据库专家进行数据管理和数据库设计(可以实现具体的功能)。
```

### 前后分离的架构--CLD 三层
```text
现在更为流行的一般是前后分离的架构,前后端之间通过 ajax 来交互，有时候要解决跨域的问题，但也已经有了较为成熟的方案
见图“images/web01”
现在比较流行的纯后端 api 模块一般采用下述划分方法：
1.Controller，
    与上述类似，服务入口，负责处理路由，参数校验，请求转发。
2.Logic/Service，逻辑(服务)层，
    一般是业务逻辑的入口，可以认为从这里开始，所有的请求参数一定是合法的。业务逻辑和业务流程也都在这一层中。
    常见的设计中会将该层称为 Business Rules。
3.DAO/Repository，
    这一层主要负责和数据、存储打交道。将下层存储以更简单的函数、接口形式暴露给 Logic 层来使用。负责数据的持久化工作。

每一层都会做好自己的工作，然后用请求当前的上下文构造下一层工作所需要的结构体或其它类型参数，
然后调用下一次的函数。在工作完成之后，再把处理结果一层层地传出到入口。
见图“images/web02”

划分为 CLD 三层之后，在 C 层之前我们可能还需要同时支持多种协议。本章前面讲到的 thrift、gRPC 和 http 并不是一定只选择其中一种，
有时我们需要支持其中的两种，比如同一个接口，我们既需要效率较高的 thrift，也需要方便 debug 的 http 入口。即除了 CLD 之外，
还需要一个单独的 protocol 层 (协议层)，负责处理各种交互协议的细节。这样请求的流程会变成下面这样：
见图“images/web03”

这样我们 controller 中的入口函数就变成了下面这样：
func CreateOrder(ctx context.Context, req *CreateOrderStruct) (*CreateOrderRespStruct, error) {
}
CreateOrder 有两个参数，ctx 用来传入 trace_id 一类的需要串联请求的全局参数，req 里存储了我们创建订单所需要的所有输入信息。
返回结果是一个响应结构体和错误。

既然工作流已经成型，我们可以琢磨一下怎么让整个流程对用户更加友好。
本节中所叙述的分层没有将 middleware 作为项目的分层考虑进去
见图“images/web05”
middleware 是和 http 协议强相关的，遗憾的是在 thrift 中看起来没有和 http 中对等的解决这些
非功能性逻辑代码重复问题的 middleware。所以我们在图上写 thrift stuff。这些 stuff 可能需要你手写去实现，
然后每次增加一个新的 thrift 接口，就需要去写一遍这些非功能性代码。(thrif这个分支放弃)
```

### 项目架构
```text
proto           protobuf层 ，生成业务逻辑的接口
protocol        协议层（http/thrift）
controller      控制层
middleware      中间件层
logic           业务逻辑层
dao             数据存储层
server          服务层
interface       服务网关（gateway）
google          生成服务网关的引用包
```

### thrift  --放弃
```text
下载：http://thrift.apache.org/download
安装：https://blog.csdn.net/Sicily_winner/article/details/86751896
```

```text
//router.pb.gw.go
//生成网关使用部分
protoc -I C:/protoc/protoc/bin -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
       --grpc-gateway_out=logtostderr=true:. \
       order.proto product.proto

//router.swagger.json
在对外公布REST接口时，我们一般还会提供一个Swagger格式的文件用于描述这个接口规范。
在网页中提供REST接口的文档和测试等功能。
protoc -I C:/protoc/protoc/bin -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
    --swagger_out=. \
    order.proto product.proto


//生成验证的命令
//validate.validator.pb.go
 protoc -I C:/protoc/protoc/bin -I. \
        -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
        --gogo_out=.  --govalidators_out=gogoimport=true:.  \
        order.proto product.proto

//.router.pb.go
//生成微服务
 protoc -I C:/protoc/protoc/bin -I. \
        -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
        --go_out=plugins=grpc:. \
        order.proto product.proto
```

### 存在的问题
```text
1.传参的中文出现乱码
2.post访问获取不到数据
```

### gorm框架
```text
https://gorm.io/docs/
```


