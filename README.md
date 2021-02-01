# gin-demo
用gin框架搭建的一个项目结构，方便快速开发项目。

### 特点

- 集成gorm，用户mysql存储层操作

- 集成go redis，用户操作缓存

- 集成uber/zap, zap是一个高效的日志组件

  项目中日志进行了分解:

  - request日志  

    记录http请求的request和response结果

  - panic日志

    记录http请求产生的panic错误日志，方便快速定位错误问题

  - app日志

    记录开发过程中，我们记录的业务日志，这是我们最常用的日志，日志当中记录了requestId，方便快速根据请求查看当前请求的日志流，
    日志格式如下,包括了TrequestId、file、keywords，以及我们记录的重要信息:
    
 ```shell

    2021-02-01 20:34:17.918	DEBUG	1nsdi6gxbL01kiv0beQiTIPv6kT	service/user.go:19	测试下	[{"Name":"测试","Age":4},{"Name":"测试","Age":3},{"Name":"测试","Age":0}]


 ```

- 集成gopkg.in/ini，用户解析我们的配置项

### 目录结构

- config -配置
- constant -常量
- controller -业务控制器
- logs -日志目录
- middleware -中间件目录
- model -表模型目录
- router -路由配置目录
- service -服务存放目录
- dto -定义除数据表实体外的其它结构体
- dao -数据库操作
- core -定义了log、db、redis等基础组件的封装、封装
- util -工具目录
- main.go -程序入库文件

### 其它

- 日志按天进行分割
- 根目录的.env.local进行本地配置,.env作为线上配置，避免误将本地配置提交到git仓库
- 业务日志可以根据requestId查看当次请求的所有日志
