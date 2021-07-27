package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)


func UpdateService(c *gin.Context) {

	db := DbService{}
	db.InitDb()
	defer db.CloseDB()

	no := c.PostForm("no")
	name := c.PostForm("name")
	score := c.PostForm("score")

	s , err := strconv.ParseFloat(score,32)
	if err != nil {
		fmt.Println()
		return
	}

	db.UpdateDB(no,name,s)

}