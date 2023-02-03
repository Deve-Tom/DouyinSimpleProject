# DouyinSimpleProject

- 2023.2.3 贺胜 initial
  1. 创建了bisicFunction和test文件夹，basicFunction用于保存基本功能模块的源文件，test文件夹用于保存单元测试源文件
  2. 在主目录下创建了main.go源文件以及router.go源文件，main.go源文件用于在之后进行集成（在开发之中不做任何处理，可暂时忽略），router.go文件引用自simple-demo项目，删除了附属功能2的所有内容
  3. 在basicFunction文件夹下创建了common.go源文件，本源文件引用自simple-demo项目，内部定义了信息结构体（可能以后会用到）
  4. 创建了go.mod模块管理文件
  5. 在go.mod之中导入了"github.com/RaymondCode/simple-demo/service"以及"github.com/gin-gonic/gin"两个模块，这两个模块在之后可能会被使用到
