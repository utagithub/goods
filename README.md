# goods


> ⚠️⚠️⚠️ 账号 / 密码： admin / 123456

## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库

- 配置文件简单的模型映射，快速能够得到想要的配置

- TODO: 单元测试

## 🎁 内置

1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
1. 接口文档：根据业务代码自动生成相关的api接口文档。
1. 定时任务：自动化任务，目前支持接口调用和函数调用。

## 准备工作

你需要在本地安装 [go] [gin] [node](http://nodejs.org/) 和 [git](https://git-scm.com/)


## 📦 本地开发

### 环境要求

go 1.18

node版本: v14.16.0

npm版本: 6.14.11

### 开发目录创建

```bash

# 创建开发目录
mkdir goods
cd goods
```

### 获取代码

```bash

git clone https://github.com/goods-team/goods-ui.git

```

### 启动说明

#### 服务端启动说明

```bash
# 进入 goods 项目
cd ./goods

# 编译项目
go build

# 修改配置
# 文件路径  goods/config/settings.yml
vi ./config/setting.yml

# 1. 配置文件中修改数据库信息
# 注意: settings.database 下对应的配置数据
# 2. 确认log路径
```

:::tip ⚠️注意 在windows环境如果没有安装中CGO，会出现这个问题；

```bash
E:\goods>go build
# github.com/mattn/go-sqlite3
cgo: exec /missing-cc: exec: "/missing-cc": file does not exist
```

or

```bash
D:\Code\goods>go build
# github.com/mattn/go-sqlite3
cgo: exec gcc: exec: "gcc": executable file not found in %PATH%
```


:::

#### 初始化数据库，以及服务启动

``` bash
# 首次配置需要初始化数据库资源信息
# macOS or linux 下使用
$ ./goods migrate -c config/settings.dev.yml

# ⚠️注意:windows 下使用
$ goods.exe migrate -c config/settings.dev.yml


# 启动项目，也可以用IDE进行调试
# macOS or linux 下使用
$ ./goods server -c config/settings.yml


# ⚠️注意:windows 下使用
$ goods.exe server -c config/settings.yml
```

