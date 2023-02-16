# DouyinSimpleProject

## Process

- 基础功能实现（更新日期：2023.2.16）

| 功能       | 是否完成 | 是否存在BUG                                                  |
| ---------- | -------- | ------------------------------------------------------------ |
| 视频Feed流 | 基本完成 | 视频播放时，以倒序方式播放，但播放数据库到数据库种第一条记录后会出现死循环第一条记录视频 |
| 视频投稿   | 完成     | 暂未发现                                                     |
| 个人主页   | 基本完成 | 个人主页处无法出现自己发表视频的封面、无法展示作品数量、无法展示视频获赞数量 |
| 喜欢列表   | 基本完成 | 喜欢列表处同个人主页的BUG一致无法显示视频的封面              |
| 用户评论   | 基本完成 | 无法在未登陆时查看评论列表                                   |

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
├── public       // 静态文件
├── router       // 路由设置
├── service      // Service 层业务逻辑
└── utils        // 一些工具函数
```
## Getting started

1. Install `ffmpeg` to get the video frame.

    We use `ffmpeg` just to extract the video frame as the cover image.

    Please refer to some documents or search in google to find `How to install ffmpeg?`.
    
    If you can't install it, it doesn't matter.  And we just simply use the `default.jpg` instead.

2. Change `server.host` in `conf.yaml` to your **Local Machine's IP Address**.

    Please refer to some documents or search in google to find `How to get your computer's IP Address?`

3. Set `BaseUrl` in the app.

    Refer to [https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#mC5eiD](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#mC5eiD).
    
    If you successfully get your IP Address, please use **your own IP Address** instead `http://192.168.1.7:8080` in the above document to fill the `BaseUrl` field in the advanced setting.

    For example, if your own IP Addresss is `192.168.108`, you should fill the `BaseUrl` to `http://192.168.108:8080` (`8080` is the port we set in the `conf.yaml`).

4. Just run this command in the root folder of your project:
    ```
    $ go run main.go
    ```

5. Can also run this command in the root folder of your project:

```
$ go build
$ ./DouyinSimpleProject
```

## Release

- 2023.2.16 贺胜 test

  - 目前BUG
  
    - 无法在未登陆时查看评论列表
  
    - 视频播放时，以倒序方式播放，但播放数据库到数据库种第一条记录后会出现死循环第一条记录视频

    - ~~密码与用户匹配时不兼容第一个版本~~

    - 用户信息界面无法正常显示视频封面图片
  
- 2023.2.16 王帅

  - 实现简单的评论操作和获取评论列表功能

- 2023.2.16 谢毛毛

  - 在用户注册/登录，增加对用户名和密码长度的有效性验证；

    - 限制用户名和密码非空，用户名长度大于32, 密码长度大于6小于32；

  - 使用bcrypt, 对用户密码的加密;

    - 加密密码长度增加到size = 200;
    
    


- 2023.2.15 王帅

  - 实现简单的点赞, 取消点赞, 查看点赞列表功能

  - 修复功能异常部分

    - 可以获取用户注册时的昵称
    
    - 用户获赞数API文档中没有定义

    - 关注数和粉丝数是社交模块的功能，目前尚未实现，默认均为0

    - 修复视频列表的获取
    
    - 用户喜欢作品数和发布作品数API文档中没有定义（可能app自身问题）
    
    - 目前可以在首页正常浏览视频，用户信息页面app本身没有实现浏览视频功能
    
    

- 2023.2.15贺胜 test

  - 功能正常部分：

    - 数据库自动建立正常

    - 用户注册可以正确被存储

    - 视频上传正常，可以在数据库种正常找到记录

  - 功能异常部分：

    - 用户仅能够在注册后正确登陆，无法正常获取用户注册时的昵称、获赞数、关注数、粉丝数

    - 视频列表无法获取，无法获得用户的喜欢作品数，发布作品数以及无法正常游览其他视频
    
    

- 2023.2.14 王帅

  - 实现简单的发布视频, 用户获取发布的视频，视频流功能

  - 采用 ffmpeg 提取视频帧，作为视频封面。若提取失败，使用默认图片

  

- 2023.2.13 王帅

  - 组织项目架构, 采用分层架构: `Entity - Service - Controller`

  - 实现简单的用户登录, 注册和获取用户信息业务逻辑

  - 添加 `GORM` 模块, 定义 `User` 和 `Video` 实体类，测试 `Migration` 功能，生成表之间的外键

  - 使用 `GORM` 的 `Gen` 功能根据 `Entity` 类自动生成 `DAO` 层代码

  - 添加 `JWT` 中间件，验证 `Token`

  

- 2023.2.3 贺胜 initial

  - 创建了bisicFunction和test文件夹，basicFunction用于保存基本功能模块的源文件，test文件夹用于保存单元测试源文件

  - 在主目录下创建了main.go源文件以及router.go源文件，main.go源文件用于在之后进行集成（在开发之中不做任何处理，可暂时忽略），router.go文件引用自simple-demo项目，删除了附属功能2的所有内容

  - 在basicFunction文件夹下创建了common.go源文件，本源文件引用自simple-demo项目，内部定义了信息结构体（可能以后会用到）

  - 创建了go.mod模块管理文件

  - 在go.mod之中导入了"github.com/RaymondCode/simple-demo/service"以及"github.com/gin-gonic/gin"两个模块，这两个模块在之后可能会被使用到
