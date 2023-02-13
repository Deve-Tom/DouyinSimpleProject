# DouyinSimpleProject

## Structure

```
.
├── LICENSE
├── README.md
├── cmd          // 命令行工具
├── conf.yaml    // 配置文件
├── config       // 配置一些全局变量以及初始化操作
├── controller   // Controller 层业务逻辑
├── dao          // DAO 层业务逻辑
├── dto          // 定义 DTO 对象
├── entity       // 实体类
├── go.mod
├── go.sum
├── main.go      // 主函数
├── middleware   // 中间件
├── router       // 路由设置
├── service      // Service 层业务逻辑
└── utils        // 一些工具函数
```
## Getting started

```
$ go run main.go
```

## Release
- 2023.2.13 王帅
  1. 组织项目架构, 采用分层架构: `Entity - Service - Controller`
  2. 实现简单的用户登录, 注册和获取用户信息业务逻辑
  4. 添加 `GORM` 模块, 定义 `User` 和 `Video` 实体类，测试 `Migration` 功能，生成表之间的外键
  4. 使用 `GORM` 的 `Gen` 功能根据 `Entity` 类自动生成 `DAO` 层代码
  5. 添加 `JWT` 中间件，验证 `Token`

- 2023.2.3 贺胜 initial
  1. 创建了bisicFunction和test文件夹，basicFunction用于保存基本功能模块的源文件，test文件夹用于保存单元测试源文件
  2. 在主目录下创建了main.go源文件以及router.go源文件，main.go源文件用于在之后进行集成（在开发之中不做任何处理，可暂时忽略），router.go文件引用自simple-demo项目，删除了附属功能2的所有内容
  3. 在basicFunction文件夹下创建了common.go源文件，本源文件引用自simple-demo项目，内部定义了信息结构体（可能以后会用到）
  4. 创建了go.mod模块管理文件
  5. 在go.mod之中导入了"github.com/RaymondCode/simple-demo/service"以及"github.com/gin-gonic/gin"两个模块，这两个模块在之后可能会被使用到
