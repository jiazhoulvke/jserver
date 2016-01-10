package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var port int
var path string

func init() {
	flag.IntVar(&port, "port", 8888, "服务器端口 (默认:8888)")
	flag.StringVar(&path, "path", "", "服务器根目录 (默认为当前目录)")
}

func main() {
	var err error
	flag.Parse()
	if path == "" {
		path, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		_, err = os.Stat(path)
		if err != nil {
			panic("指定的目录:" + path + " 无效")
		}
	}
	fmt.Printf("目录: %s\n端口:%d\n", path, port)
	http.Handle("/", http.FileServer(http.Dir(path)))
	panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
