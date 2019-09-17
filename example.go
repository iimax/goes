package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	// 从 views文件夹加载所有模版
	// 要求扩展名为 .html，并且使用 'html/template' 包来解析模版内容
	app.RegisterView(iris.HTML("./public", ".html"))

	// GET
	app.Get("/", func(ctx iris.Context) {
		// 绑定数据
		ctx.ViewData("message", "Hello world!深圳")
		// render 模版
		ctx.View("index.html")
	})

	// GET /user/1 参数路由
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.HandleDir("/", "./public")
	// 启动服务器
	app.Run(iris.Addr(":8080"))
}

func main2() {
    app := iris.Default()
    app.Use(myMiddleware)

    app.Handle("GET", "/ping", func(ctx iris.Context) {
    	ctx.JSON(iris.Map{"message": "pong"})
    })
    
    // 监听HTTP请求
    app.Run(iris.Addr(":8080"))
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}