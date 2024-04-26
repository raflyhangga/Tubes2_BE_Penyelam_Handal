package scrapper

import (
	"fmt"
	"time"
)

const SINGLE_PARAM = "single"
const MANY_PARAM = "many"

func printRequestedParameters(link_awal string, link_tujuan string) {
	fmt.Println("Requested Paramters:")
	fmt.Print("Source: ")
	fmt.Println(link_awal)
	fmt.Print("Goal: ")
	fmt.Println(link_tujuan)
}

func GetPath(node *Node) []string {
	var path []string

	var itr_node *Node = node
	for itr_node != nil {
		path = append(path, *itr_node.Current)
		itr_node = itr_node.Paths 
	}
	return reversePath(path)
}

func printSolution(solutions []Node, duration time.Duration) {
	fmt.Println("======================= SOLUTION =======================")
	fmt.Println("Found", len(solutions), "solutions in", duration)
	fmt.Println("Total of links visited:", Total_Visited_Link)

	for i, Node := range solutions {
		fmt.Println("Solution -", i+1)
		for k, path := range GetPath(&Node) {
			fmt.Println(k+1, path)
		}
	}
}

func IDS_interface(link_awal string, link_tujuan string, solution_mode string) ([]Node, time.Duration) {
	printRequestedParameters(link_awal,link_tujuan)
	if(len(Cache) == 0) {
		readCacheFromFile()
	}
	fmt.Print("Starting Iterative Deepening Search ")
	var solutions []Node

	startTime := time.Now()
	if solution_mode == SINGLE_PARAM {
		fmt.Println("single solution..")
		iterative_deepening_search_single(link_awal, link_tujuan, &solutions)
	} else if solution_mode == MANY_PARAM {
		fmt.Println("multiple solution..")
		iterative_deepening_search_many(link_awal, link_tujuan, &solutions)
	}
	duration := time.Since(startTime)

	printSolution(solutions, duration)
	writeCacheToFile()

	return solutions, duration
}

func BFS_interface(link_awal string, link_tujuan string, solution_mode string) ([]Node, time.Duration) {
	printRequestedParameters(link_awal,link_tujuan)
	fmt.Print("Starting Breadth First Search ")
	var initial []*Node
	initial = append(initial, &Node{
		Current: &link_awal,
		Paths: nil,
	})

	// find the solution using Breadth First Search (BFS) algorithm
	var solutions []Node
	startTime := time.Now()
	if solution_mode == SINGLE_PARAM {
		fmt.Println("single solution..")
		breadth_first_search_one_solution(initial, &link_tujuan, &solutions)
	} else if solution_mode == MANY_PARAM {
		fmt.Println("multiple solution..")
		breadth_first_search_many_solution(initial, &link_tujuan, &solutions)
	}
	duration := time.Since(startTime)

	printSolution(solutions, duration)

	// clean map Visited_Node
	Visited_Node = make(map[string]bool)

	return solutions, duration
}