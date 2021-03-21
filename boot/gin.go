package boot

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"ms_novel/global"
	"ms_novel/middleware"
	"ms_novel/routers"
)

func GinLaunch() {
	// 启动小说服务器,初始化gin实例
	global.Gin = gin.New()

	// 加载全局中间件
	global.Gin.Use(
		middleware.GlobalExceptionMiddleware,
		middleware.GlobalLoggerMiddleware(),
	)

	// 加载路由
	routers.InitRouter(global.Gin)

	// 设置gin的运行模式
	gin.SetMode(gin.DebugMode)

	// 服务优雅关闭和重启
	// global.Gin.Run()
	endless.ListenAndServe(":8081", global.Gin)

	// 程序结束前关闭数据库连接，避免大量数据库连接未关闭，造成连接数满
	defer MysqlClose()
}

// 内置的初始化函数，服务启动时，只会执行一次
func init() {
	// 初始化全局验证器
	validateInit()
	// 初始化env全局配置
	envInit()
	// 初始化yaml配置
	yamlInit()
	// 数据库连接
	mysqlConnect()
}
