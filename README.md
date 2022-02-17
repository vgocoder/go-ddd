README

[English](./README_EN.md)

===========================
# Go DDD Scaffold

> go-ddd是一个遵循领域驱动设计（DDD）的Golang脚手架项目，旨在展示如何构建易于开发、维护且易于使用的Go应用程序，尤其是从长远来看！

## 脚手架理念

1. 规范开发流程：写代码之前，必须先做数据库设计以及接口文档设计，
2. 减少重复开发工作：接口路由自动生成，数据库表Model自动生成，常见的CRUD只需要简单几行代码
3. 专注业务：主要业务逻辑集中在领域层（Domain），业务逻辑与产品需求遵循同一套设计语言
4. 代码易读：遵循领域驱动设计（DDD）模式，规范代码分层

## 脚手架特点

1. 自动根据数据库表生成Model
2. 自动根据OpenAPI3.0接口文档生成路由
3. 遵循领域驱动设计（DDD）模式
4. 使用go-chi路由:轻量，快，扩展性好，完全兼容标准库net/http
5. 使用Go语言的热加载工具: 监听文件或目录的变化，自动编译，重启程序,大大提高开发期的工作效率
6. 使用spf13/cobra，提供了简单的接口来创建命令行程序

## 脚手架目录说明
>脚手架目录遵循领域驱动设计（DDD）模式

### Domain
> 这是定义应用程序的域和业务逻辑的地方

### Infrastructure
> 此层包含独立于我们的应用程序而存在的所有内容：外部库，数据库引擎等。

### Application
> 该层用作域和界面层之间的通道。将请求从接口层发送到域层，由域层处理请求并返回响应。

### Interfaces
> 该层包含与其他系统交互的所有内容，例如Web服务，RMI接口或Web应用程序以及批处理前端。

## 脚手架依赖类库

### 路由中间件：[go-chi](https://github.com/go-chi/chi)

``
go get -u github.com/go-chi/chi/v5
``

### 数据库中间件：[sqlboiler](https://github.com/volatiletech/sqlboiler)
> 安装ORM

``
go install github.com/volatiletech/sqlboiler/v4@latest
``

> 安装驱动

> postgresql驱动

``
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
``

mysql驱动

``
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
``

### 接口代码生成器：[oapi-codegen](https://github.com/deepmap/oapi-codegen)
> 安装

``
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
``

### 服务热启动：[air](https://github.com/cosmtrek/air/)

> binary will be $(go env GOPATH)/bin/air

``
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
``

> or install it into ./bin/

``
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
air -v
``

### 命令行工具：[spf13/cobra](https://github.com/spf13/cobra)
> 安装

``
go get -u github.com/spf13/cobra
``

> 引用

``
import "github.com/spf13/cobra"
``

## 服务启动

``
air
``

or

``
./bin/air
``