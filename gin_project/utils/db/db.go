package db

import (
	"awesomeProject/gin_project/entity"
	"database/sql"
	"fmt"
	"log"
)

// 定义一个初始化数据库的函数
func InitDB() (*sql.DB,error) {
	// DSN:Data Source Name
	dsn := "root:chaindigg@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("open db fail err:%v",err)
		return nil,err
	}
	return db,err
}

// 查询多条数据示例
func QueryMultiRowDemo(db *sql.DB) *sql.Rows {
	sqlStr := "select id, no , name, score from student where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	return rows
}


// 插入数据
func InsertRowDemo(db *sql.DB,name string,age int) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func UpdateRowDemo(db *sql.DB,no string,name string,score float64) {
	sqlStr := "update student set no=?,name=?,score=? where no = ?"
	ret, err := db.Exec(sqlStr, no,name,score,no)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func DeleteRowDemo(db *sql.DB,no string) {
	sqlStr := "delete from student where no = ?"
	ret, err := db.Exec(sqlStr, no)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 查询单条数据示例
func QueryRowDemo(db *sql.DB,no string) bool {
	sqlStr := "select id, no, name ,score from student where no=?"
	var u entity.Student
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, no).Scan(&u.Id, &u.No, &u.Name,&u.Score)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return true
	}
	fmt.Printf("id:%d no:%d name:%s score:%d\n", u.Id, u.No,u.Name, u.Score)
	return false
}

// 预处理插入示例
func PrepareInsertDemo(db *sql.DB, no string,name string,score float64) bool {
	sqlStr := "insert into student(no,name,score) values (?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(no,name, score)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	fmt.Println("insert success.")
	return true
}

func PrepareDelete(db *sql.DB,id int) {
	sqlStr := "delete from  student(id) values (?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Printf("detelt failed, err:%v\n", err)
		return
	}
	fmt.Println("delete success.")
}