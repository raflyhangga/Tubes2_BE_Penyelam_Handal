package scrapper

import (
	"fmt"
	"sync"
	"time"
)

func printSolution(solutions []Node, duration time.Duration) {
	fmt.Println("======================= SOLUTION =======================")
	fmt.Println("Found ", len(solutions), " solutions in", duration)
	for k, value := range solutions {
		fmt.Println("Solution - ", k+1)
		for k, path := range value.Paths {
			fmt.Println(k+1, path)
		}
		fmt.Println(len(value.Paths)+1, value.Current)
	}
}

func IDS_interface(link_awal string, link_tujuan string) {
	var wg sync.WaitGroup
	graph := constructGraph([]string{}, link_awal, MAX_DEPTH, &wg)

	var solutions []Node
	startTime := time.Now()
	iterative_deepening_search(graph, link_tujuan, MAX_DEPTH, &solutions)
	duration := time.Since(startTime)

	printSolution(solutions, duration)

}

func BFS_interface(link_awal string, link_tujuan string) {
	var wg sync.WaitGroup
	graph := constructGraph([]string{}, link_awal, MAX_DEPTH, &wg)

	initialQ := []Node{graph}
	solutions := []Node{}
	startTime := time.Now()
	breadth_first_search(initialQ, link_tujuan, &solutions)
	duration := time.Since(startTime)

	printSolution(solutions, duration)

}
