package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var distPath string = os.Getenv("DIST_PATH")

var Router *RouterHandler = new(RouterHandler)

type RouterHandler struct {
}

var mux = make(map[string]func(http.ResponseWriter, *http.Request))

func (p *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if fun, ok := mux[r.URL.Path]; ok {
		fun(w, r)
		return
	}

	var path string

	if r.URL.Path[1:] == "" {
		path = "index.html"
	} else {
		path = r.URL.Path[1:]
	}

	filePath := distPath + path

	fileExit, _ := PathExists(filePath)

	if fileExit {
		file, _ := ioutil.ReadFile(filePath)
		w.Write(file)
		return
	}

	http.Error(w, "error URL:"+r.URL.String(), http.StatusBadRequest)

}

func (p *RouterHandler) Router(relativePath string, handler func(http.ResponseWriter, *http.Request)) {
	mux[relativePath] = handler
}

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
