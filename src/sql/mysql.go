package main

import (
	"database/sql"
	. "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql","root:123456@/test?charset=utf8")
	checkErr(err)
	
	// 插入数据
	stmt,err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res,err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	Println(id)

	stmt,err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res,err = stmt.Exec("neil", id)
	checkErr(err)

	affect,err := res.RowsAffected()
	checkErr(err)

	Println(affect)
	
	rows,err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		Println(uid)
		Println(username)
		Println(department)
		Println(created)
	}

	//删除数据
	/* stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	Println(affect) */

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}