package scrapper

import "fmt"

/**
 * Function to perform Depth First Search (DFS) algorithm
 *
 * @param current_node: the current node (current link)
 * @param link_tujuan: title of the destination link
 * @param depth: the maximum depth to search
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
var THREADS_AMOUNT = 2

func iterative_deepening_search(current_node Node, link_tujuan string, depth int, hasil *[]Node) {
	if depth == 0 {
		// check is the current node is the destination node
		if current_node.Current == link_tujuan {
			*hasil = append(*hasil, current_node)
		}
	} else {
		if current_node.Current == link_tujuan { // check is the current node is the destination node
			*hasil = append(*hasil, current_node)
		} else { // if the current node is not the destination node
			for _, node := range current_node.Neighbours {
				iterative_deepening_search(node, link_tujuan, depth-1, hasil)
			}
		}
	}
}

/**
 * Function to perform Breadth First Search (BFS) algorithm
 *
 * @param current_queue: a list of nodes (links) to be processed
 * @param destination_link: link of the destination link
 * @param hasil: a list of nodes that contain destination node
 */
func breadth_first_search(current_queue []Node, destination_link string, hasil *[]Node) {
	if len(current_queue) == 0 {
		return
	} else {
		// go through all the nodes in the current queue
		// check if the destination node is in the current queue
		var new_queue []Node
		isDestinationLinkExist := false
		for _, current_node := range current_queue {
			fmt.Println("BFS : ", current_node.Current)
			visitedNode[current_node.Current] = true
			total_nodes++
			new_queue = append(new_queue, current_node.Neighbours...)
			if current_node.Current == destination_link {
				*hasil = append(*hasil, current_node)
				isDestinationLinkExist = true
			}
		}

		if isDestinationLinkExist {
			// solution is found, return
			return
		} else {
			// recursively call the breadth_first_search function with the new queue
			breadth_first_search(new_queue, destination_link, hasil)
		}
	}
}
