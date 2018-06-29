package main

import (
	"apiservice/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建路由引擎
	g := gin.New()
	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	log.Printf("Strat to listening the incoming requests on http addresss: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}
