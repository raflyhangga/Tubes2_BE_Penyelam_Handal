package scrapper

import (
	"fmt"
	"time"
)

func printSolution(solutions []Node, duration time.Duration) {
	fmt.Println("======================= SOLUTION =======================")
	fmt.Println("Found", len(solutions), "solutions in", duration)
	fmt.Println("Total of links visited:", Total_Visited_Link)

	for i, value := range solutions {
		fmt.Println("Solution -", i+1)
		for k, path := range value.Paths {
			fmt.Println(k+1, path)
		}
		fmt.Println(len(solutions[i].Paths)+1, value.Current)
	}
}

func IDS_interface(link_awal string, link_tujuan string, sum_solution string) ([]Node, time.Duration) {
	// find the solution using Iterative Deepening Search (IDS) algorithm
	var solutions []Node

	if len(Cache) == 0 {
		readCacheFromFile()
	}
	startTime := time.Now()
	iterative_deepening_search(link_awal, link_tujuan, sum_solution, &solutions)
	duration := time.Since(startTime)

	printSolution(solutions, duration)
	writeCacheToFile()

	return solutions, duration
}

func BFS_interface(link_awal string, link_tujuan string, sum_solution string) ([]Node, time.Duration) {
	var initial = Node{
		Current: link_awal,
	}

	// find the solution using Breadth First Search (BFS) algorithm
	var solutions []Node
	startTime := time.Now()
	if sum_solution == "1" {
		breadth_first_search_one_solution([]Node{initial}, link_tujuan, &solutions)
	} else if sum_solution == "0" {
		breadth_first_search_many_solution([]Node{initial}, link_tujuan, &solutions)
	}
	duration := time.Since(startTime)

	printSolution(solutions, duration)

	// clean map Visited_Node
	Visited_Node = make(map[string]bool)

	return solutions, duration
}
