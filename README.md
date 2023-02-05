# gin-demo
用gin框架搭建的一个项目结构，方便快速开发项目。

### 特点

- 自研sqlx库，用于mysql存储层操作

- 集成go redis，用于操作缓存

- 集成uber/zapgo-file-rotatelogs, 记录日志

  项目中日志进行了分解:

  - request日志  

    记录http请求的request和response结果

  - panic日志

    记录http请求产生的panic错误日志，方便快速定位错误问题

  - app日志

    记录开发过程中，我们记录的业务日志，这是我们最常用的日志，日志当中记录了traceId，方便快速根据请求查看当前请求的日志流，
    日志格式如下,包括了traceId、file、keywords，以及我们记录的重要信息:
    
  ```shell
      2022-11-05 00:40:23.576	DEBUG	dao/user.go:24	83e51872-35a3-455d-9c5f-bd64ee140d4c	sql	[0.527ms] [rows:0] SELECT * FROM `users` WHERE `users`.`uid` = 7777 ORDER BY `users`.`uid` LIMIT 1

  ```
  - cron 定时任务业务日志
  - cron_panic 定时任务产生的panic日志

- 集成gopkg.in/yaml.v3，用于解析yaml文件的配置项
- 集成go-cache,用于本地缓存
- 集成pkg/errors,用于更好的error体验

### 目录结构

- config -配置
- handler -业务控制器&路由配置,可以根据不同的子目录拆分模块handler
- logs -日志目录
- middleware -中间件目录
- entity -表模型目录
- logic -业务逻辑目录,可以根据不同的子目录拆分模块
- dao -数据库操作
- svc -定义了log、db、redis等基础服务组件的封装
- types 定义了请求参数的数据结构
- cron 定时任务
- server.go -程序入库文件

### 其它

- 日志按天进行分割
- 业务日志可以根据traceId查看当次请求的所有日志
- 响应，返回固定的字段，包括code、message、data、traceid