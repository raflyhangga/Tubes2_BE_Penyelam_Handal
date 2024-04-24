package scrapper

/**
 * Function to perform Depth Limited Search (DLS) algorithm
 *
 * @param currentNode: the current node (current link)
 * @param destinationLink: destination link
 * @param depth: the maximum depth to search
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func depth_limited_search(currentNode Node, destinationLink string, depth int, hasil *[]Node) {
	Total_Visited_Link++

	if depth == 0 { // if the current node is the ddepest node
		if currentNode.Current == destinationLink {
			*hasil = append(*hasil, currentNode)
		}
	} else {
		if currentNode.Current == destinationLink { // check is the current node is destination node
			*hasil = append(*hasil, currentNode)
		} else { // if the current node is not the destination node
			tetangga := getAdjacentLinks(currentNode)
			for _, node := range tetangga {
				depth_limited_search(node, destinationLink, depth-1, hasil)
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
func iterative_deepening_search(startLink string, destinationLink string, hasil *[]Node) {
	var initial = Node{
		Current: startLink,
	}

	var solutions []Node
	depth := 0

	for len(solutions) == 0 {
		depth_limited_search(initial, destinationLink, depth, &solutions)
		depth++
	}

	*hasil = solutions
}
