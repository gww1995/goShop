package main

import (
	"github.com/gin-gonic/gin"
	"net"
)

var goodsList gin.HandlerFunc

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer tcp.Close()
	return tcp.Addr().(*net.TCPAddr).Port, nil
}

func main() {
	router := gin.Default()
	router.Group("/goods")
	router.GET("/goods", goodsList)

}
