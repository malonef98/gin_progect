package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func Deleteervice(c *gin.Context) {

	db := DbService{}
	db.InitDb()
	defer db.CloseDB()

	//表单参数设置默认值
	type1 := c.DefaultPostForm("type", "alert")


	no := c.PostForm("no")

	flag := db.DeleteDB(no)

	if flag {
		c.String(200,
			fmt.Sprintf(type1,"删除成功"))
	}else {
		c.String(200,
			fmt.Sprintf(type1,"删除失败"))
	}
}

