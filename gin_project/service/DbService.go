package service

import (
	db2 "awesomeProject/gin_project/utils/db"
	"database/sql"
	"fmt"
)

type DbService struct {
	db *sql.DB
	err error
}

func (this *DbService) InitDb()  {
	this.db,this.err = db2.InitDB()
	if this.err!= nil {
		fmt.Println(this.err)
	}
}

func (this *DbService) QueryAll() *sql.Rows {
	return db2.QueryMultiRowDemo(this.db)
}

func (this *DbService) DeleteDB(no string) bool  {
	if this.FindById(no) {
		db2.DeleteRowDemo(this.db,no)
		fmt.Println("删除成功")
		return true
	}
	return false
}

func (this *DbService) CloseDB()  {
	this.db.Close()
}

func (this *DbService) FindById(no string) bool {
	if db2.QueryRowDemo(this.db,no) {
		fmt.Println("找不到对应学号的学生的数据")
		return false
	}
	return true
}

func (this *DbService) Inert(no string,name string,score float64) bool {
	if db2.QueryRowDemo(this.db,no) == false{
		fmt.Println("错误：学号输入重复")
		return false
	}
	return db2.PrepareInsertDemo(this.db,no,name,score)
	return true
	//db2.InsertRowDemo(this.db,name,age)
}

func (this *DbService) UpdateDB(no string,name string,score float64)  {
	db2.UpdateRowDemo(this.db,no,name,score)
}