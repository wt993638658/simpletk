# tiktok

基于 kitex RPC微服务 + gin HTTP服务完成的第五届字节跳动青训营-极简抖音后端项目

## 一、项目特点

1. 采用RPC框架（Kitex）脚手架生成代码进行开发，基于 **RPC 微服务** + **gin 提供 HTTP 服务**

2. 基于《[接口文档在线分享](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)[- Apifox](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)》提供的接口进行开发，使用《[极简抖音](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)[App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 》提供的APK进行Demo测试， **功能完整实现** ，前端接口匹配良好。

3. 代码结构采用 (HTTP API 层 + RPC Service 层+Dal 层) 项目 **结构清晰** ，代码 **符合规范**

4. 使用 **JWT** 进行用户token的校验

5. 使用 **ETCD** 进行服务发现和服务注册；

6. 使用 **Minio** 实现视频文件和图片的对象存储

7. 使用 **Gorm** 对 MySQL 进行 ORM 操作；

8. 使用 **OpenTelemetry** 实现链路跟踪；

9. 数据库表建立了索引和外键约束，对于具有关联性的操作一旦出错立刻回滚，保证数据一致性和安全性

## 二、项目地址

- **<https://github.com/wt993638658/simpletk>**

## 三、项目说明

### 1. 项目模块介绍

| 服务名称    | 模块介绍                   | 技术框架 | 传输协议 | 注册中心 | 链路跟踪 | 数据存储 | 日志 | 配置存取 |
|---------|------------------------| --- | --- | --- | --- | --- | --- | --- |
| api     | API服务将HTTP请求发送给RPC微服务端 | `Gorm` `Kitex` `Hertz` | `http` | `etcd`| `opentelemetry` | 下一步计划采用Redis | `zapklog` | `viper` |
| user    | 用户管理微服务                | `Gorm` `Kitex` `Hertz` `JWT` | `proto3` ||| `MySQL` `gorm` |
| relation | 用户关注微服务                |
| feed    | 视频流微服务                 |
| favorite | 用户点赞微服务                |
| comment | 用户评论微服务                |
| message | 消息微服务                  |
| publish | 视频发布微服务                |||||`MySQL` `gorm` `minio对象存储` |
| dal     | 数据层实现                  | `MySQL` `gorm` | - || | `MySQL` `gorm` |

### 2. 服务调用关系
![img.png](img.png)

- HTTP 使用 Hertz 开放 HTTP 端口, 通过封装的RPC客户端与微服务中的服务端通信;

- RPC 微服务, 接收客户端的请求, 在各自的 command 中实现与数据库交互的业务逻辑;

- DAL 提供数据层实现, pack 部分实现将数据库输出封装为服务端的响应结构体;

- MiddleWare 提供中间件业务逻辑代码实现.

### 3. 代码介绍

#### 3.1 代码目录结构介绍

| 目录                                                                           | 子目录                                                                                                                                                                                         | 说明 | 备注                                                                                                                                               |
|------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------| --- |--------------------------------------------------------------------------------------------------------------------------------------------------|
| [cmd](https://github.com/wt993638658/simpletk/tree/master/cmd)               | [api](https://github.com/wt993638658/simpletk/tree/master/cmd/api)                                                                                                                          | api 服务的 **业务代码** | 包含 [gin](https://github.com/wt993638658/simpletk/blob/master/cmd/api/main.go)和 [RPC_client](https://github.com/wt993638658/simpletk/tree/master/cmd/api/rpc) |
|| [comment](https://github.com/wt993638658/simpletk/tree/master/cmd/comment)   | command 服务的业务代码                                                                                                                                                                             |
|| [favorite](https://github.com/wt993638658/simpletk/tree/master/cmd/favorite) | favorite 服务的业务代码                                                                                                                                                                            |
|| [feed](https://github.com/wt993638658/simpletk/tree/master/cmd/feed)         | feed 服务的业务代码                                                                                                                                                                                |
|| [message](https://github.com/wt993638658/simpletk/tree/master/cmd/message)   | message 服务的业务代码                                                                                                                                                                             |
|| [publish](https://github.com/wt993638658/simpletk/tree/master/cmd/publish)         | publish 服务的业务代码                                                                                                                                                                             |
|| [relation](https://github.com/wt993638658/simpletk/tree/master/cmd/publish)        | relation 服务的业务代码                                                                                                                                                                            |
|| [user](https://github.com/wt993638658/simpletk/tree/master/cmd/user)               | user 服务的业务代码                                                                                                                                                                                |
| [config](https://github.com/wt993638658/simpletk/tree/master/config)               | 微服务及 pkg 的 **配置文件**                                                                                                                                                                         |
| [dal](https://github.com/wt993638658/simpletk/tree/master/dal)                     | [db](https://github.com/wt993638658/simpletk/tree/master/dal/db)                                                                                                                                  | 包含 [Gorm 初始化](https://github.com/wt993638658/simpletk/blob/master/dal/db/init.go) 、[Gorm 结构体及 数据库操作逻辑](https://github.com/wt993638658/simpletk/blob/master/dal/db/user.go) |
|| [pack](https://github.com/wt993638658/simpletk/tree/master/dal/pack)               | 将 [Gorm 结构体](https://github.com/wt993638658/simpletk/blob/master/dal/pack/user.go#L25) 封装为 [protobuf 结构体](https://github.com/wt993638658/simpletk/blob/master/kitex_gen/user/user.pb.go#L268)的 **业务逻辑** | Protobuf 结构体由 Kitex自动生成 |
| [idl](https://github.com/wt993638658/simpletk/tree/master/idl)                     | proto **接口定义文件**                                                                                                                                                                            |
| [kitex_gen](https://github.com/wt993638658/simpletk/tree/master/kitex_gen)         | Kitex **自动生成的代码**                                                                                                                                                                           |
| [pkg](https://github.com/wt993638658/simpletk/tree/master/pkg)                     | [dlog](https://github.com/wt993638658/simpletk/tree/master/pkg/dlog)                                                                                                                              | 基于 klog 和 zap 封装的 **Logger** 及其接口 |
|| [errno](https://github.com/wt993638658/simpletk/tree/master/pkg/errno)             | **错误码**                                                                                                                                                                                     | 错误码设计逻辑:[a76yyyy/ErrnoCod](https://github.com/a76yyyy/ErrnoCode) |
|| [jwt](https://github.com/wt993638658/simpletk/tree/master/pkg/jwt)                 | 基于 [golang-jwt](http://github.com/golang-jwt/jwt)的代码封装                                                                                                                                      |
|| [middleware](https://github.com/wt993638658/simpletk/tree/master/pkg/middleware)   | Kitex的中间件                                                                                                                                                                                   |
|| [minio](https://github.com/wt993638658/simpletk/tree/master/pkg/minio)             | **Minio** 对象存储初始化及代码封装                                                                                                                                                                      |
|| [ttviper](https://github.com/wt993638658/simpletk/tree/master/pkg/ttviper)         | **Viper** 配置存取初始化及代码封装                                                                                                                                                                      |

#### 3.2 代码运行

1. 提前修改 [config](https://github.com/wt993638658/simpletk/tree/master/config)目录的相关配置

2. 运行基础依赖

    ``` Shell
    # 自行安装 docker 及 docker-compose
    docker-compose up -d 
    ```

3. 运行 user 服务

    ``` Shell
    cd cmd/user
    sh build.sh
    sh output/bootstrap.sh 
    ```

4. 运行 comment 服务

    ``` Shell
    cd cmd/comment
    sh build.sh
    sh output/bootstrap.sh 
    ```

5. 运行 favorite 服务

    ``` Shell
    cd cmd/favorite
    sh build.sh
    sh output/bootstrap.sh 
    ```

6. 运行 feed 服务

    ``` Shell
    cd cmd/feed
    sh build.sh
    sh output/bootstrap.sh 
    ```

7. 运行 publish 服务

    ``` Shell
    cd cmd/publish
    sh build.sh
    sh output/bootstrap.sh 
    ```

8. 运行 relation 服务

    ``` Shell
    cd cmd/relation
    sh build.sh
    sh output/bootstrap.sh 
    ```

9. 运行 message 服务
    ``` Shell
    cd cmd/message
    sh build.sh
    sh output/bootstrap.sh 
    ```

10. 运行 api 服务

    ``` Shell
    cd cmd/api
    chmod +x ./run.sh
    sh ./run.sh 
    ```