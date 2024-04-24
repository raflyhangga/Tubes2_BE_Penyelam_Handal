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
	Total int				`json:"total_visited"`
}

func ids_router(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	link_1 := PREFIX + queryParams.Get("init")
	link_2 := PREFIX + queryParams.Get("goal")

	fmt.Println("Starting iterative deepening search...")
	fmt.Println(link_1)
	fmt.Println(link_2)
	solution,duration := scrapper.IDS_interface(link_1,link_2)
	
	var pack Package
	pack.Duration = duration.String()
	pack.Total = scrapper.TotalVisitedLink
	for _,node := range solution {
		pack.Solutions = append(pack.Solutions,node.Paths)
	}
	if(len(solution) != 0) {
		context.IndentedJSON(http.StatusFound,pack)
	} else {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"No Solution"})
	}
	scrapper.TotalVisitedLink = 0
}

func bfs_router(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	link_1 := PREFIX + queryParams.Get("init")
	link_2 := PREFIX + queryParams.Get("goal")

	fmt.Println("Starting breadth first search...")
	fmt.Println(link_1)
	fmt.Println(link_2)
	solution,duration := scrapper.BFS_interface(link_1,link_2)

	var pack Package
	pack.Duration = duration.String()
	pack.Total = scrapper.TotalVisitedLink

	for _,node := range solution {
		pack.Solutions = append(pack.Solutions,node.Paths)
	}

	if(len(solution) != 0) {
		context.IndentedJSON(http.StatusFound,pack)
	} else {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"No Solution"})
	}
	scrapper.TotalVisitedLink = 0
}

func main(){
	router := gin.Default()
	router.GET("/ids",ids_router)
	router.GET("/bfs",bfs_router)
	router.Run("localhost:9090")
}
