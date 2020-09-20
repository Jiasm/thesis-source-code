package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", HelloWorldHandler)
	err := http.ListenAndServe("0.0.0.0:8880", nil)

	if err != nil {
		fmt.Println("服务器错误")
	}

}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Fprintf(w,"HelloWorld!")
}
