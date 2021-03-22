package util

import (
	"encoding/json"
	"net/http"
)

type ResponseSuccess struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ResultFail(w http.ResponseWriter, err string) {
	ResultFailWithCode(w, err, http.StatusBadRequest)
}

func ResultFailWithCode(w http.ResponseWriter, err string, code int) {
	response := ResponseError{
		Code: code,
		Msg:  err,
	}

	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(res)
}

func ResultJsonOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response := ResponseSuccess{
		Code: 200,
		Data: data,
	}

	res, _ := json.Marshal(response)

	w.Write(res)
}
