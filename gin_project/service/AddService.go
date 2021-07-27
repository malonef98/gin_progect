package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func AddService(c *gin.Context) {

	db := DbService{}
	db.InitDb()
	defer db.CloseDB()

	//表单参数设置默认值
	type1 := c.DefaultPostForm("type", "alert")

	//接收username,password
	no := c.PostForm("no")
	name := c.PostForm("name")
	score := c.PostForm("score")

	s , err := strconv.ParseFloat(score,32)
	if err != nil {
		fmt.Println()
		return
	}

	flag := db.Inert(no,name,s)

	if flag {
		c.HTML(http.StatusOK, "index.html", nil)
		c.String(200,
			fmt.Sprintf(type1, no, name, score))
	}else {
		c.HTML(http.StatusOK, "index.html", nil)
		c.String(200,
			fmt.Sprintf(type1,"增加失败"))
	}


}