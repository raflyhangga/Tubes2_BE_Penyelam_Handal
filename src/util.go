package main

import (
	"net/http"
	"time"
	"wikipedia-scraper-engine/scrapper"

	"github.com/gin-gonic/gin"
)

const PREFIX = "/wiki/"
const ALGORITHM_PARAM = "algorithm"
const SOLUTIONS_PARAM = "solutions"
const BFS_PARAM = "bfs"
const IDS_PARAM = "ids"
const SOURCE_QUERY = "source"
const GOAL_QUERY = "goal"

type Package struct {
	Solutions [][]string `json:"solutions"`
	Duration  string     `json:"duration"`
	Total     int        `json:"total_visited"`
}

func getRequests(context *gin.Context) (source string, goal string, algorithm_mode string, solution_mode string) {
	queryParams := context.Request.URL.Query()
	algorithm_mode = context.Param("algorithm")
	source = queryParams.Get("source")
	goal = queryParams.Get("goal")
	solution_mode = context.Param("solutions")

	return source,goal,algorithm_mode,solution_mode
}

func isLinkValid(context *gin.Context) bool {
	queryParams := context.Request.URL.Query()
	algorithm_mode := context.Param(ALGORITHM_PARAM)
	source := queryParams.Get(SOURCE_QUERY)
	goal := queryParams.Get(GOAL_QUERY)
	solution_mode := context.Param(SOLUTIONS_PARAM)
	return  len(source) != 0 && len(goal) != 0 && len(solution_mode) != 0 && (algorithm_mode == BFS_PARAM || algorithm_mode == IDS_PARAM) && (solution_mode == scrapper.SINGLE_PARAM || solution_mode == scrapper.MANY_PARAM)
}

func wrapPackage(durasi time.Duration,  solusi []scrapper.Node)Package{
	var pack Package
	pack.Duration = durasi.String()
	pack.Total = scrapper.Total_Visited_Link
	for _, node := range solusi {
		path := scrapper.AddDomainPrefix(node.Paths)
		current := scrapper.DOMAIN_PREFIX + node.Current
		pack.Solutions = append(pack.Solutions, append(path, current))
	}
	return pack
}

func sendRequest(context *gin.Context, pack Package) {
	if len(pack.Solutions) != 0 {
		context.IndentedJSON(http.StatusFound, pack)
	} else {
		context.IndentedJSON(http.StatusFound, gin.H{"message": "No Solution"})
	}
}