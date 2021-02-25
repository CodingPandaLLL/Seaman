package main

import (
	"Seaman/config"
	"Seaman/controller"
	"Seaman/datasource"
	"Seaman/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

func main() {

	app := newApp()

	//应用App设置
	configation(app)

	//路由设置
	mvcHandle(app)

	config := config.InitConfig()

	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr), //在端口8080进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)
}

//构建App
func newApp() *iris.Application {
	app := iris.New()

	//设定应用图标
	app.Favicon("./static/favicons/favicon.ico")

	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")
	app.HandleDir("/img", "./uploads")

	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context iris.Context) {
		context.View("index.html")
	})

	return app
}

/**
 * MVC 架构模式处理
 */
func mvcHandle(app *iris.Application) {

	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	//获取redis实例
	//redis := datasource.NewRedis()
	//设置session的同步位置为redis
	//sessManager.UseDatabase(redis)

	//实例化mysql数据库引擎
	engine := datasource.NewMysqlEngine()

	//安全模块功能
	securityService := service.NewSecurityService(engine)

	admin := mvc.New(app.Party("/service/security"))
	admin.Register(
		securityService,
		sessManager.Start,
	)
	admin.Handle(new(controller.SecurityController))

	//用户功能模块
	userService := service.NewUserService(engine)
	user := mvc.New(app.Party("/service/user"))
	user.Register(
		userService,
		sessManager.Start,
	)
	user.Handle(new(controller.UserController))

	//组织机构功能模块
	orgService := service.NewOrgService(engine)
	org := mvc.New(app.Party("/service/org"))
	org.Register(
		orgService,
		sessManager.Start,
	)
	org.Handle(new(controller.OrgController))

	//角色功能模块
	appRoleService := service.NewAppRoleService(engine)
	appRole := mvc.New(app.Party("/service/appRole"))
	appRole.Register(
		appRoleService,
		sessManager.Start,
	)
	appRole.Handle(new(controller.AppRoleController))

	//群组功能模块
	appGroupService := service.NewAppGroupService(engine)
	appGroup := mvc.New(app.Party("/service/appGroup"))
	appGroup.Register(
		appGroupService,
		sessManager.Start,
	)
	appGroup.Handle(new(controller.AppGroupController))

	//项目功能模块
	projectService := service.NewProjectService(engine)
	project := mvc.New(app.Party("/service/project"))
	project.Register(
		projectService,
		sessManager.Start,
	)
	project.Handle(new(controller.ProjectController))

	//项目文件功能模块
	projectFileService := service.NewProjectFileService(engine)
	projectFile := mvc.New(app.Party("/service/projectFile"))
	projectFile.Register(
		projectFileService,
		sessManager.Start,
	)
	projectFile.Handle(new(controller.ProjectFileController))

}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found ",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    " interal error ",
			"data":   iris.Map{},
		})
	})
}
