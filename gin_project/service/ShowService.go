package service

import (
	"awesomeProject/gin_project/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//查询数据

func ShowService(c *gin.Context){
	db := DbService{}
	db.InitDb()
	defer db.CloseDB()

	//查
	//querySql := "select id, no, name, score from student "
	rows  := db.QueryAll()
	defer rows.Close()
	//if err != nil {
	//	log.Printf("query data error:%v\n", err)
	//	return
	//}
	s := new(entity.Student)


	for rows.Next() {
		rows.Scan(&s.Id, &s.No, &s.Name, &s.Score)
		log.Println(*s)
		data := map[string]interface{}{
			"ID":    &s.Id,
			"No":    &s.No,
			"Name":  &s.Name,
			"Score": &s.Score,
		}
		c.HTML(http.StatusOK, "list.html", data)
	}
}
