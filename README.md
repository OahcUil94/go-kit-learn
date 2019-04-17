# go-kit学习

### demo1
- 简单web接口实现、基础功能封装、Fail接口错误统一处理，统一响应格式，自定义错误
- 请求：curl -d '{"orderId": "112233"}' 127.0.0.1:8080/order/create
- 响应：{"code":0,"msg":"success","data":{"orderId":"#112233","source":"APP","isPay":1}}

### demo2
- 添加自定义服务中间件、端点中间件、请求的before、after使用，以及ctx的使用、进一步封装结构
- 监控地址：http://127.0.0.1:8088/metrics
- 其他同上

### demo3
- GRPC服务端和客户端编写。并同时兼容http和grpc两种请求方式，以及cmd/service封装和复用
- grpc模拟客户端请求文件：cmd/grpc_client/main.go

### demo4
- 编写中...
- 使用限流器、断路器、普罗米修斯指标等功能实现

### demo5
- 编写中...
- 添加mysql、redis，配置文件操作，参数校验等封装使用

### demo6
- 编写中...
- 使用etcd服务注册发现，内部服务之间的订阅和发布（Nats）、使用Dockerfile和docker-compose命令来创建和管理容器

### 其他目录
- configs: 配置文件目录
- deployments: 部署配置目录，如：docker-compose.yml 
- internal: 微服务项目的共享和私有代码目录
- tools: 工具目录，可以从internal目录中导入代码使用
- 参考：https://github.com/golang-standards/project-layout

### 参考
- https://gokit.io/examples
