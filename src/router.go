package main

import (
	"net/http"
	"wikipedia-scraper-engine/scrapper"
	"github.com/gin-gonic/gin"
)

// Serves the router endpoint for Iterative Depth Search
func ids_router(context *gin.Context, source string, goal string, solution_mode string) {
	solution, duration := scrapper.IDS_interface(PREFIX+source, PREFIX+goal, solution_mode)
	sendRequest(context, wrapPackage(duration, solution))
}

// Serves the router endpoint for Breadth First Search
func bfs_router(context *gin.Context, source string, goal string, solution_mode string) {
	solution, duration := scrapper.BFS_interface(PREFIX+source, PREFIX+goal, solution_mode)
	sendRequest(context, wrapPackage(duration, solution))
}

// Serves as the entry point for requests
func main_router(context *gin.Context) {
	if isLinkValid(context) {
		source, goal, algorithm_mode, solution_mode := getRequests(context)
		if algorithm_mode == BFS_PARAM {
			bfs_router(context, source, goal, solution_mode)
		} else {
			ids_router(context, source, goal, solution_mode)
		}
		scrapper.Total_Visited_Link = 0
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Bad getaway, check your api."})
	}
}
