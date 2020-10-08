package main

import (
	"time"
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"github.com/jmoiron/sqlx"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id int
	Username string
	Password string
	Status int
}

type Question struct {
	Id 				int 			`json:"id"`
	Title 		string 		`json:"title"`
	Creater 	int 			`json:"creater"`
	Status 		int 			`json:"status"`
	Datetime 	time.Time `json:"datetime"`
}

type QuestionRequestData struct {
	Id				int
	Title 		string
	Creater 	int
	Datetime 	time.Time
}

type ResetPassword struct {
	Id int
	Phone string
	VerifyCode string
	Pwd string
}

const dbConfig = "test:123456@tcp(127.0.0.1:3306)/survey?parseTime=true"

func main()  {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/list", ListHandler)
	http.HandleFunc("/info", GetHandler)
	http.HandleFunc("/add", AddNewHandler)
	http.HandleFunc("/change", ChangeHandler)
	http.HandleFunc("/delete", RemoveHandler)
	// http.HandleFunc("/reset-password", ResetPasswordHandler)

	fmt.Println("server run as http://127.0.0.1:8880")
	err := http.ListenAndServe("0.0.0.0:8880", nil)

	if err != nil {
		fmt.Println("服务器错误")
	}
}

func CORSHandler(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	w.Header().Add("Access-Control-Allow-Credentials", "true") //设置为true，允许ajax异步请求带cookie信息
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
	w.Header().Set("content-type", "application/json;charset=UTF-8")             //返回数据格式是json
}

func LoginHandler(w http.ResponseWriter, r *http.Request)  {
	CORSHandler(w)

	body, err := ioutil.ReadAll(r.Body)
	var data Admin
	if err = json.Unmarshal(body, &data); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}

	user, err := getUserInfo(data.Username, data.Password)

	if err != nil || user.Status != 1 {
		fmt.Fprintf(w, "用户不存在")
	} else {
		fmt.Fprintf(w, "success")
	}
}

func ListHandler(w http.ResponseWriter, r *http.Request)  {
	CORSHandler(w)

	data := getQuestionList()

	jsons, errs := json.Marshal(data)
	if errs != nil {
	  fmt.Println(errs.Error())
	}

	fmt.Fprintf(w, string(jsons))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	CORSHandler(w)

	var id string
	values := r.URL.Query()
	id = values.Get("id")

	intId, errs := strconv.Atoi(id)

	question := getQuestionInfo(intId)

	jsons, errs := json.Marshal(question)
	if errs != nil {
	  fmt.Println(errs.Error())
	}

	fmt.Fprintf(w, string(jsons))
}

func AddNewHandler(w http.ResponseWriter, r *http.Request) {
	CORSHandler(w)

	body, err := ioutil.ReadAll(r.Body)
	var data QuestionRequestData
	if err = json.Unmarshal(body, &data); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}
	id := setQuestionInfo(data)
	fmt.Fprintf(w, strconv.FormatInt(id, 10))
}

func ChangeHandler(w http.ResponseWriter, r *http.Request)  {
	CORSHandler(w)

	body, err := ioutil.ReadAll(r.Body)
	var data QuestionRequestData
	if err = json.Unmarshal(body, &data); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}
	changeQeustionInfo(data)
	fmt.Fprintf(w, "success")
}

func RemoveHandler(w http.ResponseWriter, r *http.Request)  {
	CORSHandler(w)

	body, err := ioutil.ReadAll(r.Body)
	var data QuestionRequestData
	if err = json.Unmarshal(body, &data); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
	}
	removeQeustionInfo(data)
	fmt.Fprintf(w, "success")
}

// func ResetPasswordHandler(w http.ResponseWriter, r *http.Request)  {
// 	body, err := ioutil.ReadAll(r.Body)
// 	var data ResetPassword
// 	if err = json.Unmarshal(body, &data); err != nil {
// 			fmt.Printf("Unmarshal err, %v\n", err)
// 			return
// 	}
// 	resetPassword(data)
// 	fmt.Fprintf(w, "success")
// }

func getQuestionList() []Question {
	Db, err := sqlx.Connect("mysql", dbConfig)
	
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return nil
	}
	defer Db.Close()

	rows, _ := Db.Query("SELECT id, title, creater, status, datetime FROM QuestionInfo")

	var id, creater, status int
	var title string
	var datetime time.Time
	var questions []Question
 
	for rows.Next() {
		err := rows.Scan(&id, &title, &creater, &status, &datetime)
		if err != nil {
			fmt.Println(err)
		}
		questions = append(questions, Question{ id, title, creater, status, datetime })
	}
 
	return questions
}

func setQuestionInfo(data QuestionRequestData) int64 {
	Db, err := sqlx.Connect("mysql", dbConfig)

	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return 0
	}
	defer Db.Close()

	res, err := Db.Exec("INSERT INTO QuestionInfo (title, creater, status, datetime) VALUES(?, ?, ?, ?)", data.Title, data.Creater, 1, data.Datetime)

	if err != nil{
		fmt.Println("exec mysql failed,",err)
		return 0
	}

	id, err := res.LastInsertId()

	if err != nil{
		fmt.Println("get last insert id failed,",err)
		return 0
	}

	return id
}

func changeQeustionInfo(data QuestionRequestData) bool {
	Db, err := sqlx.Connect("mysql", dbConfig)
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return false
	}
	defer Db.Close()

	Db.Exec("UPDATE QuestionInfo SET title = ? WHERE id = ?", data.Title, data.Id)

	return true
}

func removeQeustionInfo(data QuestionRequestData) bool {
	Db, err := sqlx.Connect("mysql", dbConfig)
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return false
	}
	defer Db.Close()

	Db.Exec("UPDATE QuestionInfo SET status = 0 WHERE id = ?", data.Id)

	return true
}

// func resetPassword(data ResetPassword) bool {
// 	user, err := getUserInfo(data.id)

// 	if err != nil ||  user == nil || user.phone != data.phone {
// 		return Error.new("用户不存在")
// 	}
// 	Db, err := sqlx.Connect("mysql", "test:123456@tcp(127.0.0.1:3306)/survey")
// 	if err != nil{
// 		fmt.Println("connect to mysql failed,",err)
// 		return
// 	}
// 	defer Db.Close()

// 	Db.Exec("UPDATE Admin SET password = ? WHERE id = ?", data.pwd, data.id)

// 	return true
// }

// func getFrontUserInfo(phone string, pwd string) FrontUser {
// 	Db, err := sqlx.Connect("mysql", "test:123456@tcp(127.0.0.1:3306)/survey")
// 	if err != nil{
// 		fmt.Println("connect to mysql failed,",err)
// 		return
// 	}
// 	defer Db.Close()

// 	var user User

// 	err := Db.Get(&user, "SELECT * FROM user WHERE phone = ? AND password = ?", username, password)

// 	if err != nil{
// 		fmt.Println("insert failed", err)
// 	}

// 	if user.status == 0 {
// 		return Error.new("账号被禁用")
// 	}

// 	if user.status == 2 {
// 		return Error.new("账号被锁定")
// 	}

// 	return user
// }

func getQuestionInfo(id int) Question {
	var question Question

	Db, err := sqlx.Connect("mysql", dbConfig)
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return question
	}
	defer Db.Close()

	err = Db.Get(&question, "SELECT id, title, creater, status, datetime FROM QuestionInfo WHERE id = ?", id)

	if err != nil{
		fmt.Println("insert failed", err)
		return question
	}

	return question
}

func getUserInfo(username, password string) (Admin, error) {
	var admin Admin

	Db, err := sqlx.Connect("mysql", dbConfig)
	if err != nil{
		fmt.Println("connect to mysql failed,",err)
		return admin, err
	}
	defer Db.Close()

	err = Db.Get(&admin, "SELECT id, username, password, status FROM admin WHERE username = ? AND password = ? AND status = 1", username, password)

	if err != nil{
		fmt.Println("insert failed", err)
		return admin, err
	}

	return admin, nil
}