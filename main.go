package main

import (
	"github.com/gin-gonic/gin"
)

var goodsList gin.HandlerFunc

func main() {
	router := gin.Default()
	router.Group("/goods")
	router.GET("/goods", goodsList)

}
