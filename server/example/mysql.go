package main

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

type struct Admin {
	id int
	username string
	password string
	status int
}

func setQuestionInfo(title string, creater int, datetime string) Question {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	var question Question 

	Db.Exec(&question, "INSERT INTO QuestionInfo (title, creater, status, datetime) VALUES(?, ?, ?, ?)", title, creater, 1, datetime)

	return question
}

func changeQeustionInfo(id int, title string) bool {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	Db.Exec("UPDATE QuestionInfo SET title = ? WHERE id = ?", title, id)

	return true
}

func removeQeustionInfo(id int) bool {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	Db.Exec("UPDATE QuestionInfo SET status = 0 WHERE id = ?", id)

	return true
}

func resetPassword(id int, phone string, verifyCode string, pwd string) bool {
	user := getUserInfo(id)

	if user == nil || user.phone != phone {
		return Error.new("用户不存在")
	}
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	Db.Exec("UPDATE Admin SET password = ? WHERE id = ?", pwd, id)

	return true
}

func getFrontUserInfo(phone string, pwd string) FrontUser {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()
	
	var user User

	err := Db.Get(&user, "SELECT * FROM user WHERE phone = ? AND password = ?", username, password)

	if err != nil{
		fmt.Println("insert failed", err)
	}

	if user.status == 0 {
		return Error.new("账号被禁用")
	}

	if user.status == 2 {
		return Error.new("账号被锁定")
	}
	
	return user
}


func getUserInfo(username, password string) Admin {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()
	
	var admin Admin

	err := Db.Get(&admin, "SELECT * FROM admin WHERE username = ? AND password = ? AND status = 1", username, password)

	if err != nil{
		fmt.Println("insert failed", err)
	}
	
	return admin
}