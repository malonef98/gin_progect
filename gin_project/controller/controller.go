package controller

import (
	"awesomeProject/gin_project/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (this *Controller) Init()  {
	//1. 创建路由
	r := gin.Default()

	//html渲染
	r.LoadHTMLFiles("/Users/mayifan/Desktop/list.html","/Users/mayifan/Desktop/index.html")

	//2.
	r.POST("/add", service.AddService)

	r.GET("/show",service.ShowService)

	r.POST("/delete",service.Deleteervice)

	r.POST("/update",service.UpdateService)

	//3. 监听
	r.Run()
}