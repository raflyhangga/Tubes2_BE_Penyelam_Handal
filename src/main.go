package main

import (
	"fmt"
	"net/http"
	"wikipedia-scraper-engine/scrapper"

	"github.com/gin-gonic/gin"
)

const PREFIX = "/wiki/"

type Package struct {
	Solutions [][]string `json:"solutions"`
	Duration  string     `json:"duration"`
	Total     int        `json:"total_visited"`
}

func addDomainPrefix(list []string) []string {
	realLink := make([]string, len(list))
	for i, link := range list {
		realLink[i] = scrapper.DOMAIN_PREFIX + link
	}
	return realLink
}

func ids_router(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	link_1 := queryParams.Get("init")
	link_2 := queryParams.Get("goal")
	sumsolution := queryParams.Get("onesolution")

	fmt.Println("Starting Iterative Deepening Search...")
	fmt.Println(PREFIX + link_1)
	fmt.Println(PREFIX + link_2)
	solution, duration := scrapper.IDS_interface(PREFIX+link_1, PREFIX+link_2, sumsolution)

	var pack Package
	pack.Duration = duration.String()
	pack.Total = scrapper.Total_Visited_Link
	for _, node := range solution {
		path := addDomainPrefix(node.Paths)
		current := scrapper.DOMAIN_PREFIX + node.Current
		pack.Solutions = append(pack.Solutions, append(path, current))
	}

	if len(solution) != 0 {
		context.IndentedJSON(http.StatusFound, pack)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Solution"})
	}

	scrapper.Total_Visited_Link = 0
}

func bfs_router(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	link_1 := queryParams.Get("init")
	link_2 := queryParams.Get("goal")
	sumsolution := queryParams.Get("onesolution")

	fmt.Println("Starting Breadth First Search...")
	fmt.Println(PREFIX + link_1)
	fmt.Println(PREFIX + link_2)
	solution, duration := scrapper.BFS_interface(PREFIX+link_1, PREFIX+link_2, sumsolution)

	var pack Package
	pack.Duration = duration.String()
	pack.Total = scrapper.Total_Visited_Link
	for _, node := range solution {
		path := addDomainPrefix(node.Paths)
		current := scrapper.DOMAIN_PREFIX + node.Current
		pack.Solutions = append(pack.Solutions, append(path, current))
	}

	if len(solution) != 0 {
		context.IndentedJSON(http.StatusFound, pack)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Solution"})
	}

	scrapper.Total_Visited_Link = 0
}

func main() {
	router := gin.Default()
	router.GET("/ids", ids_router)
	router.GET("/bfs", bfs_router)
	router.Run("localhost:9090")
}
