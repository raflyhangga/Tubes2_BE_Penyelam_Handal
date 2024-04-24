package scrapper

import (
	"fmt"
	"time"
)

func printSolution(solutions []Node, duration time.Duration) {
	fmt.Println("======================= SOLUTION =======================")
	fmt.Println("Found", len(solutions), "solutions in", duration)
	fmt.Println("Total link visited:", totalVisitedLink)
	for k, value := range solutions {
		fmt.Println("Solution - ", k+1)
		for k, path := range value.Paths {
			fmt.Println(k+1, path)
		}
		fmt.Println(len(value.Paths)+1, value.Current)
	}
}

func IDS_interface(link_awal string, link_tujuan string) ([]Node, time.Duration) {
	var solutions []Node

	startTime := time.Now()
	iterative_deepening_search(link_awal, link_tujuan, &solutions)
	duration := time.Since(startTime)

	printSolution(solutions, duration)
	return solutions, duration
}

func BFS_interface(link_awal string, link_tujuan string) ([]Node, time.Duration) {
	var initial Node
	initial.Current = link_awal

	var solutions []Node
	startTime := time.Now()
	breadth_first_search([]Node{initial}, link_tujuan, &solutions)
	duration := time.Since(startTime)

	printSolution(solutions, duration)
	return solutions, duration
}
