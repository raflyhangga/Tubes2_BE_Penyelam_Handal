package scrapper

import (
	"fmt"
)

/**
 * Function to perform Depth Limited Search (DLS) algorithm
 *
 * @param currentNode: the current node (current link)
 * @param destinationLink: destination link
 * @param depth: the maximum depth to search
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func depth_limited_search_many_solution(currentNode Node, destinationLink string, depth int, hasil *[]Node) {
	Total_Visited_Link++

	if depth == 0 { // if the current node is the ddepest node
		if currentNode.Current == destinationLink {
			*hasil = append(*hasil, currentNode)
		}
	} else {
		if currentNode.Current == destinationLink { // check is the current node is destination node
			*hasil = append(*hasil, currentNode)
		} else { // if the current node is not the destination node
			var tetangga []Node
			if isInCache(currentNode.Current) {
				tetanggaCache := Cache[currentNode.Current]
				for i := 0; i < len(tetanggaCache); i++ {
					tetangga = append(tetangga, tetanggaCache[i])
					tetangga[i].Paths = append(currentNode.Paths, currentNode.Current)
				}
			} else {
				tetangga = getAdjacentLinks(currentNode)
				cacheTetangga := make([]Node, len(tetangga))
				for i := 0; i < len(tetangga); i++ {
					cacheTetangga[i].Current = tetangga[i].Current
				}
				Cache[currentNode.Current] = cacheTetangga
			}
			for _, node := range tetangga {
				depth_limited_search_many_solution(node, destinationLink, depth-1, hasil)
			}
		}
	}
}

/**
 * Function to perform Depth Limited Search (DLS) algorithm
 *
 * @param currentNode: the current node (current link)
 * @param destinationLink: destination link
 * @param depth: the maximum depth to search
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func depth_limited_search_one_solution(currentNode Node, destinationLink string, depth int, hasil *[]Node) {
	if len(*hasil) > 0 {
		return
	} else if depth == 0 { // if the current node is the ddepest node
		Total_Visited_Link++
		if currentNode.Current == destinationLink {
			*hasil = append(*hasil, currentNode)
		}
	} else {
		Total_Visited_Link++
		if currentNode.Current == destinationLink { // check is the current node is destination node
			*hasil = append(*hasil, currentNode)
		} else { // if the current node is not the destination node
			var tetangga []Node
			if isInCache(currentNode.Current) {
				tetanggaCache := Cache[currentNode.Current]
				for i := 0; i < len(tetanggaCache); i++ {
					tetangga = append(tetangga, tetanggaCache[i])
					tetangga[i].Paths = append(currentNode.Paths, currentNode.Current)
				}
			} else {
				tetangga = getAdjacentLinks(currentNode)
				cacheTetangga := make([]Node, len(tetangga))
				for i := 0; i < len(tetangga); i++ {
					cacheTetangga[i].Current = tetangga[i].Current
				}
				Cache[currentNode.Current] = cacheTetangga
			}
			for _, node := range tetangga {
				depth_limited_search_one_solution(node, destinationLink, depth-1, hasil)
				if len(*hasil) > 0 {
					return
				}
			}
		}
	}
}

/**
 * Function to perform Iterative Deepening Search (IDS) algorithm
 *
 * @param startLink: the start link
 * @param destinationLink: destination link
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func iterative_deepening_search(startLink string, destinationLink string, sumSolution string, hasil *[]Node) {
	var initial = Node{
		Current: startLink,
	}

	var solutions []Node
	depth := 0

	if sumSolution == "1" {
		for len(solutions) == 0 {
			fmt.Println("Starting new depth..")
			depth_limited_search_one_solution(initial, destinationLink, depth, &solutions)
			depth++
		}
	} else if sumSolution == "0" {
		for len(solutions) == 0 {
			fmt.Println("Starting new depth..")
			depth_limited_search_many_solution(initial, destinationLink, depth, &solutions)
			depth++
		}
	}

	*hasil = solutions
}
