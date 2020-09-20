package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type struct Admin {
	id int
	username string
	password string
	status int
}

type struct Question {
	id int
	title string
	creater int
	status int
	datetime string
}

type struct QuestionRequestData {
	title string
	creater int
	datetime string
}

type struct ResetPassword {
	id int
	phone string
	verifyCode string
	pwd string
}

func main()  {
	http.HandleFunc("/list", ListHandler)
	http.HandleFunc("/add", AddNewHandler)
	http.HandleFunc("/change", ChangeHandler)
	http.HandleFunc("/delete", RemoveHandler)
	http.HandleFunc("/reset-password", ResetPasswordHandler)
	err := http.ListenAndServe("0.0.0.0:8880", nil)

	if err != nil {
		fmt.Println("服务器错误")
	}

}

func ListHandler(w http.ResponseWriter, r *http.Request)  {
	data := getQuestionList()

	jsons, errs := json.Marshal(data)
	if errs != nil {
	  fmt.Println(errs.Error())
	}

	fmt.Fprintf(w, jsons)
}

func AddNewHandler(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	var data QuestionRequestData
	if err = json.Unmarshal(body, &a); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}
	setQuestionInfo(data)
	fmt.Fprintf(w, "success")
}

func ChangeHandler(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	var data QuestionRequestData
	if err = json.Unmarshal(body, &a); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}
	changeQeustionInfo(data)
	fmt.Fprintf(w, "success")
}

func RemoveHandler(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	var data QuestionRequestData
	if err = json.Unmarshal(body, &a); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}
	removeQeustionInfo(data)
	fmt.Fprintf(w, "success")
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	var data ResetPassword
	if err = json.Unmarshal(body, &a); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}
	resetPassword(data)
	fmt.Fprintf(w, "success")
}

func getQuestionList() Question[] {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	rows, _ := Db.Query("SELECT * FROM QuestionInfo")
	var id, creater, status int
	var title, datetime string
	var questions []Question
 
	for rows.Next() {
		err := rows.Scan(&id, &title, &creater, &status, &datetime)
		if err != nil { /* error handling */}
		questions = append(questions, Question{ id, title, creater, status, datetime })
	}
 
	return questions
}

func setQuestionInfo(data QuestionRequestData) Question {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	var question Question 

	Db.Exec(&question, "INSERT INTO QuestionInfo (title, creater, status, datetime) VALUES(?, ?, ?, ?)", data.title, data.creater, 1, data.datetime)

	return question
}

func changeQeustionInfo(data QuestionRequestData) bool {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	Db.Exec("UPDATE QuestionInfo SET title = ? WHERE id = ?", data.title, data.id)

	return true
}

func removeQeustionInfo(data QuestionRequestData) bool {
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	Db.Exec("UPDATE QuestionInfo SET status = 0 WHERE id = ?", data.id)

	return true
}

func resetPassword(data ResetPassword) bool {
	user := getUserInfo(data.id)

	if user == nil || user.phone != data.phone {
		return Error.new("用户不存在")
	}
	Db, err:= sqlx.Open("mysql", "root:password@tcp(192.168.0.1:3306)/survey")
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return
	}
	defer Db.Close()

	Db.Exec("UPDATE Admin SET password = ? WHERE id = ?", data.pwd, data.id)

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