package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	
	router.Use(cors.Default())
	router.GET("/:algorithm/:solutions", main_router)
	router.Run("localhost:9090")
}
