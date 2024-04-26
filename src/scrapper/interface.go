package scrapper

import (
	"fmt"
	"time"
)

/**
 *    GLOBAL VARIABLES
 *
 */
const SINGLE_PARAM = "single"
const MANY_PARAM = "many"

/**
 * Procedure to print the requested parameters
 *
 * @param link_awal: the source link
 * @param link_tujuan: the destination link
 */
func printRequestedParameters(link_awal string, link_tujuan string) {
	fmt.Println("\n<=== Requested Paramters ===>")
	fmt.Print("Source: ")
	fmt.Println(link_awal)
	fmt.Print("Goal: ")
	fmt.Println(link_tujuan)
}

/**
 * Procedure to print the solution
 *
 * @param solutions: a list of nodes that contain the path from the start node to the destination node
 * @param duration: the duration to find the solution
 */
func printSolution(solutions []Node, duration time.Duration) {
	fmt.Println("======================= SOLUTION =======================")
	fmt.Println("Found", len(solutions), "solutions in", duration)
	fmt.Println("Total of links visited:", Total_Visited_Link)
	fmt.Printf("Link rate: %.2f link/second\n",getLinkRate(duration))

	for i, value := range solutions {
		fmt.Println("Solution -", i+1)
		for k, path := range value.Paths {
			fmt.Println(k+1, path)
		}
		fmt.Println(len(solutions[i].Paths)+1, value.Current)
	}
}

/**
 * Function to do the Iterative Deepening Search (IDS) algorithm
 *
 * @param link_awal: the start link
 * @param link_tujuan: destination link
 * @param solution_mode: the solution mode (single or many)
 * @return a list of solution that contain the path from the start node to the destination node and the duration to find the solution
 */
func IDS_interface(link_awal string, link_tujuan string, solution_mode string) ([]Node, time.Duration) {
	printRequestedParameters(link_awal, link_tujuan)

	// read cache from file
	if len(Cache) == 0 {
		readCacheFromFile()
	}

	fmt.Print("Starting Iterative Deepening Search ")
	var solutions []Node

	// find the solution using Iterative Deepening Search (IDS) algorithm
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

/**
 * Function to do the Breadth First Search (BFS) algorithm
 *
 * @param link_awal: the start link
 * @param link_tujuan: destination link
 * @param solution_mode: the solution mode (single or many)
 * @return a list of solution that contain the path from the start node to the destination node and the duration to find the solution
 */
func BFS_interface(link_awal string, link_tujuan string, solution_mode string) ([]Node, time.Duration) {
	printRequestedParameters(link_awal, link_tujuan)
	fmt.Print("Starting Breadth First Search ")
	var initial = Node{
		Current: link_awal,
	}

	// find the solution using Breadth First Search (BFS) algorithm
	var solutions []Node
	startTime := time.Now()
	if solution_mode == SINGLE_PARAM {
		fmt.Println("single solution..")
		breadth_first_search_one_solution([]Node{initial}, link_tujuan, &solutions)
	} else if solution_mode == MANY_PARAM {
		fmt.Println("multiple solution..")
		breadth_first_search_many_solution([]Node{initial}, link_tujuan, &solutions)
	}
	duration := time.Since(startTime)

	printSolution(solutions, duration)

	// clean map Visited_Node
	Visited_Node = make(map[string]bool)

	return solutions, duration
}
