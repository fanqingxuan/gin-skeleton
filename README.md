# gin-skeleton

用gin框架配合golang方面比较优秀的库，搭建的一个项目结构，方便快速开发项目。
用最少的依赖实现80%项目可以完成的需求

### 特点

- 自研sqlx库，用于mysql存储层操作
- 集成go redis，用于操作缓存
- 集成uber用于记录日志

  项目中日志按天进行了分解:

  - request日志

    记录http的request日志，方便查看请求参数
  - error日志

    记录http请求产生的panic、warn和error等级的日志，方便快速定位错误问题
  - info日志

    记录开发过程中的debug、info级别的日志

    日志格式如下,包括了traceId、file、line，以及我们记录的重要信息:

  ```shell
     {"level":"ERROR","time":"2023-02-08 12:53:51.369","caller":"user/infohandler.go:19","msg":"Handler ShouldBind Parse","traceId":"67d783ba-e34f-429f-a077-97e87d00e2d3"}


  ```
- 集成gopkg.in/yaml.v3，用于解析yaml文件的配置项
- 集成go-cache,用于本地缓存

### 目录结构

- config -配置
- handler -业务控制器&路由配置,可以根据不同的子目录拆分模块handler
- logs -日志目录
- middleware -中间件目录
- model -数据模型，数据库查询操作通常在这里完成
- logic -业务逻辑目录,可以根据不同的子目录拆分模块
- svc -定义了log、db、redis等基础服务组件的封装
- types 定义了请求参数、相应格式的数据结构
- main.go -程序入口文件

### 其它

- 日志按天进行分割
- 业务日志可以根据traceId查看当次请求的所有日志
- 响应，返回固定的字段，包括code、message、data、traceid
- 本项目倡导简单精良，go.mod文件依赖如下

```shell
	github.com/gin-gonic/gin v1.8.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	go.uber.org/zap v1.23.0
	gopkg.in/yaml.v3 v3.0.1
```

### 快速启动

```shell
1. https://github.com/fanqingxuan/gin-skeleton.git
2. cd gin-skeleton && go mod tidy
3. go run main.go -f .env.yaml
```
