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
func depth_limited_search_many_solution(currentNode *Node, destinationLink string, depth int, hasil *[]Node) {
	Total_Visited_Link++

	if depth == 0 { // if the current node is the depest node
		if *currentNode.Current == destinationLink {
			*hasil = append(*hasil, *currentNode)
		}
	} else {
		if *currentNode.Current == destinationLink { // check is the current node is destination node
			*hasil = append(*hasil, *currentNode)
		} else { // if the current node is not the destination node
			if !isInCache(currentNode.Current) {
				Cache[*currentNode.Current] = getAdjacentLinks(*currentNode)
			}
			isFound := false 
			for _, link := range Cache[*currentNode.Current] {
				var tempNode = &Node{
					Current: &link,
					Paths: currentNode,
				}
				if(link == destinationLink){
					*hasil = append(*hasil, *tempNode)
					isFound = true
				} else if(!isFound) {
					depth_limited_search_many_solution(tempNode, destinationLink, depth-1, hasil)
				}
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
func depth_limited_search_one_solution(currentNode *Node, destinationLink string, depth int, hasil *[]Node) {
	if len(*hasil) > 0 {
		return
	} else if depth == 0 { // if the current node is the ddepest node
		Total_Visited_Link++
		if *currentNode.Current == destinationLink {
			*hasil = append(*hasil, *currentNode)
		}
	} else {
		Total_Visited_Link++
		if *currentNode.Current == destinationLink { // check is the current node is destination node
			*hasil = append(*hasil, *currentNode)
		} else { // if the current node is not the destination node
			if !isInCache(currentNode.Current) {
				Cache[*currentNode.Current] = getAdjacentLinks(*currentNode)
			} 
			
			for _, link := range Cache[*currentNode.Current] {
				depth_limited_search_many_solution(&Node{
					Current: &link,
					Paths: currentNode,
				}, destinationLink, depth-1, hasil)
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
		Current: &startLink,
		Paths: nil,
	}

	var solutions []Node
	depth := 0

	for len(solutions) == 0 {
		depth_limited_search_one_solution(&initial, destinationLink, depth, &solutions)
		fmt.Println("Starting new depth..")
		depth++
	}

	*hasil = solutions
}

func iterative_deepening_search_many(startLink string, destinationLink string, hasil *[]Node) {
	var initial = Node{
		Current: &startLink,
		Paths: nil,
	}

	var solutions []Node
	depth := 0

	for len(solutions) == 0 {
		depth_limited_search_many_solution(&initial, destinationLink, depth, &solutions)
		fmt.Println("Starting new depth..")
		depth++
	}

	*hasil = solutions
}
