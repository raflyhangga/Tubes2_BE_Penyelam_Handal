package main

import (
	"fmt"
	"net/http"
	"wikipedia-scraper-engine/scrapper"

	"github.com/gin-gonic/gin"
)

const PREFIX = "https://en.wikipedia.org/wiki/"

type Package struct {
	Solutions [][]string 	`json:"solutions"`
	Duration string			`json:"duration"`
}

func ids_router(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	link_1 := PREFIX + queryParams.Get("init")
	link_2 := PREFIX + queryParams.Get("goal")

	solution,duration := scrapper.IDS_interface(link_1,link_2)
	var pack Package
	pack.Duration = duration.String()
	for _,node := range solution {
		pack.Solutions = append(pack.Solutions, append(node.Paths,link_2))
	}
	if(len(solution) != 0) {
		context.IndentedJSON(http.StatusFound,pack)
	} else {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"No Solution"})
	}
}

func bfs_router(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	link_1 := PREFIX + queryParams.Get("init")
	link_2 := PREFIX + queryParams.Get("goal")

	fmt.Println(link_1)
	fmt.Println(link_2)
	solution,duration := scrapper.BFS_interface(link_1,link_2)

	var pack Package
	pack.Duration = duration.String()
	for _,node := range solution {
		pack.Solutions = append(pack.Solutions, append(node.Paths,link_2))
	}

	if(len(solution) != 0) {
		context.IndentedJSON(http.StatusFound,pack)
	} else {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"No Solution"})
	}
}

func main(){
	router := gin.Default()
	router.GET("/ids",ids_router)
	router.GET("/bfs",bfs_router)
	router.Run("localhost:9090")
}
