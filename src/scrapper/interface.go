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
	fmt.Println(len(Cache))
	//	for key := range Cache {
	//		fmt.Println(Cache[key])
	//	}
}

func IDS_interface(link_awal string, link_tujuan string) ([]Node, time.Duration) {
	var solutions []Node
	// find the solution using Iterative Deepening Search (IDS) algorithm
	startTime := time.Now()
	iterative_deepening_search(link_awal, link_tujuan, &solutions)
	duration := time.Since(startTime)

	loadCache()

	printSolution(solutions, duration)

	return solutions, duration
}

func BFS_interface(link_awal string, link_tujuan string) ([]Node, time.Duration) {
	var initial = Node{
		Current: link_awal,
	}

	var solutions []Node
	// find the solution using Breadth First Search (BFS) algorithm
	startTime := time.Now()
	breadth_first_search([]Node{initial}, link_tujuan, &solutions)
	duration := time.Since(startTime)

	loadCache()

	printSolution(solutions, duration)

	return solutions, duration
}
