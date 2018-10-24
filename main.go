package main

import (
	"flag"
	"fmt"
	"net"
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
			fmt.Println("无法获取当前目录地址:", err)
			os.Exit(1)
		}
	} else {
		_, err = os.Stat(path)
		if err != nil {
			fmt.Println("指定的目录:" + path + " 无效")
			os.Exit(1)
		}
	}
	ip, err := getIP()
	if err != nil {
		fmt.Println("无法获取ip地址:", err)
		os.Exit(1)
	}
	fmt.Printf("目录: %s\n", path)
	fmt.Printf("服务地址: http://%s:%d\n", ip, port)
	http.Handle("/", http.FileServer(http.Dir(path)))
	panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipnet.IP.IsLoopback() {
			continue
		}
		if ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("not found")
}
