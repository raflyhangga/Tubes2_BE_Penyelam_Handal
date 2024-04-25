package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:algorithm/:solutions", main_router)
	router.Run("localhost:9090")
}
