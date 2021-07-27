package main

import (
	controller2 "awesomeProject/gin_project/controller"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	controller := controller2.Controller{}
	controller.Init()


}