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
func depth_limited_search_many_solution(currentNode Node, destinationLink string, depth int, hasil *[]Node, path []string) {
	if depth == 0 { // if the current node is the ddepest node
		if currentNode.Current == destinationLink {
			currentNode.Paths = path
			*hasil = append(*hasil, currentNode)
		}
	} else {
		path = append(path, currentNode.Current)

		if currentNode.Current == destinationLink { // check is the current node is destination node
			currentNode.Paths = path
			*hasil = append(*hasil, currentNode)
		} else { // if the current node is not the destination node
			var tetangga []Node
			if isInCache(currentNode.Current) {
				tetangga = Cache[currentNode.Current]
			} else {
				tetangga = getAdjacentLinks(currentNode, "ids")
				Cache[currentNode.Current] = tetangga
			}
			for _, node := range tetangga {
				Total_Visited_Link++
				depth_limited_search_many_solution(node, destinationLink, depth-1, hasil, path)
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

		if currentNode.Current == destinationLink { // check is the current node is destination node
			currentNode.Paths = path
			*hasil = append(*hasil, currentNode)
		} else { // if the current node is not the destination node
			var tetangga []Node
			if isInCache(currentNode.Current) {
				tetangga = Cache[currentNode.Current]
			} else {
				tetangga = getAdjacentLinks(currentNode, "ids")
				Cache[currentNode.Current] = tetangga
			}
			for _, node := range tetangga {
				Total_Visited_Link++
				depth_limited_search_one_solution(node, destinationLink, depth-1, hasil, path)
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
func iterative_deepening_search_single(startLink string, destinationLink string, hasil *[]Node) {
	var initial = Node{
		Current: startLink,
	}

	var solutions []Node
	depth := 1

	for len(solutions) == 0 {
		Total_Visited_Link = 0
		fmt.Println("Starting new depth..")
		depth_limited_search_one_solution(initial, destinationLink, depth, &solutions, []string{})
		depth++
	}

	*hasil = solutions
}

func iterative_deepening_search_many(startLink string, destinationLink string, hasil *[]Node) {
	var initial = Node{
		Current: startLink,
	}

	var solutions []Node
	depth := 0

	for len(solutions) == 0 {
		Total_Visited_Link = 0
		fmt.Println("Starting new depth..")
		depth_limited_search_many_solution(initial, destinationLink, depth, &solutions, []string{})
		depth++
	}

	*hasil = solutions
}
