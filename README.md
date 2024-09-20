创建一个基于gin 的web 服务
1. 在目录internal/router中创建_router.go结尾的文件，用于路由注册
2. 在目录internal/controller中创建_controller.go结尾的文件，用于业务逻辑处理
3. 在目录internal/service中创建_service.go结尾的文件，用于业务逻辑处理
4. 在目录internal/model中创建_model.go结尾的文件，用于数据模型定义，包括表结构，表名，表注释等，使用年gorm 的gorm.Model结构体
5. 在目录internal/dao中创建_dao.go结尾的文件，用于数据访问对象定义，包括表的增删改查等操作,其中list支持分页，search支持搜索，sort支持排序。 参数支持可以选查询。数值参数支持范围查询，时间参数支持范围查询。字符串参数支持模糊查询
6. 在目录internal/config中创建_config.go结尾的文件，用于配置定义
7. 在目录internal/utils中创建_utils.go结尾的文件，用于工具类定义
8. 在目录internal/middleware中创建_middleware.go结尾的文件，用于中间件定义

全部都是函数式编程，没有使用面向对象编程。但每个包在init.go文件中， 包涵必要的包内全局变量。例如dao包中， 包涵了db, redis, log,cache, 等全局变量。
上面目录结构的前缀包括 user, admin, system, config。 并且每个目录中都包含一个init.go文件，用于初始化。