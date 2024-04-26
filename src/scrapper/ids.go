package scrapper

import (
	"fmt"
)

/**
 * Procedure to perform Depth Limited Search (DLS) algorithm with one solution
 *
 * @param currentNode: the current node (current link)
 * @param destinationLink: destination link
 * @param depth: the maximum depth to search
 * @param path: a path from the start node to the current node
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func depth_limited_search_one_solution(currentNode Node, destinationLink string, depth int, hasil *[]Node, path []string) {
	if len(*hasil) > 0 {
		return
	} else if depth == 0 { // if the current node is the ddepest node
		if currentNode.Current == destinationLink {
			currentNode.Paths = path
			*hasil = append(*hasil, currentNode)
		}
	} else {
		path = append(path, currentNode.Current)
		var tetangga []Node

		// get the adjacent links from the current node
		if isInCache(currentNode.Current) {
			tetangga = Cache[currentNode.Current]
		} else {
			tetangga = getAdjacentLinks(currentNode, "ids")
			Cache[currentNode.Current] = tetangga
		}

		// go through all the adjacent links
		for _, node := range tetangga {
			Total_Visited_Link++
			depth_limited_search_one_solution(node, destinationLink, depth-1, hasil, path)
			if len(*hasil) > 0 {
				return
			}
		}
	}
}

/**
 * Procedure to perform Depth Limited Search (DLS) algorithm with many solutions
 *
 * @param currentNode: the current node (current link)
 * @param destinationLink: destination link
 * @param depth: the maximum depth to search
 * @param path: a path from the start node to the current node
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func depth_limited_search_many_solution(currentNode Node, destinationLink string, depth int, hasil *[]Node, path []string) {
	if depth == 0 { // if the current node is the ddepest node
		if currentNode.Current == destinationLink {
			currentNode.Paths = path
			*hasil = append(*hasil, currentNode)
		}
	} else {
		path = append(path, currentNode.Current)
		var tetangga []Node

		// get the adjacent links from the current node
		if isInCache(currentNode.Current) {
			tetangga = Cache[currentNode.Current]
		} else {
			tetangga = getAdjacentLinks(currentNode, "ids")
			Cache[currentNode.Current] = tetangga
		}

		// go through all the adjacent links
		for _, node := range tetangga {
			Total_Visited_Link++
			depth_limited_search_many_solution(node, destinationLink, depth-1, hasil, path)
		}
	}
}

/**
 * Procedure to perform Iterative Deepening Search (IDS) algorithm with one solution
 *
 * @param startLink: the start link
 * @param destinationLink: destination link
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func iterative_deepening_search_single(startLink string, destinationLink string, hasil *[]Node) {
	var initial = Node{
		Current: startLink,
	}

	var solutions []Node

	// check if the start node is the destination node
	if initial.Current == destinationLink {
		solutions = append(solutions, initial)
	}

	// start the depth from 1
	depth := 1
	for len(solutions) == 0 {
		Total_Visited_Link = 0
		fmt.Println("Starting new depth..")
		depth_limited_search_one_solution(initial, destinationLink, depth, &solutions, []string{})
		depth++
	}

	*hasil = solutions
}

/**
 * Procedure to perform Iterative Deepening Search (IDS) algorithm with one solution
 *
 * @param startLink: the start link
 * @param destinationLink: destination link
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func iterative_deepening_search_many(startLink string, destinationLink string, hasil *[]Node) {
	var initial = Node{
		Current: startLink,
	}

	var solutions []Node

	// check if the start node is the destination node
	if initial.Current == destinationLink {
		solutions = append(solutions, initial)
	}

	// start the depth from 1
	depth := 1
	for len(solutions) == 0 {
		Total_Visited_Link = 0
		fmt.Println("Starting new depth..")
		depth_limited_search_many_solution(initial, destinationLink, depth, &solutions, []string{})
		depth++
	}

	*hasil = solutions
}
