package scrapper

import (
	"fmt" // only for DEBUGGING purposes, delete when it's done
)

/**
 * Function to perform Depth First Search (DFS) algorithm
 *
 * @param current_node: the current node (current link)
 * @param link_tujuan: title of the destination link
 * @param depth: the maximum depth to search
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func iterative_deepening_search(current_node Node, link_tujuan string, depth int, hasil *[]Node) {
	if depth == 0 {
		// check is the current node is the destination node
		fmt.Println() // for Debugging, delete when it's done
		if current_node.Current == link_tujuan {
			*hasil = append(*hasil, current_node)
		}
	} else {
		fmt.Println()                            // for Debugging, delete when it's done
		if current_node.Current == link_tujuan { // check is the current node is the destination node
			*hasil = append(*hasil, current_node)
		} else { // if the current node is not the destination node
			tetangga := getAdjacentLinks(current_node)
			for _, node := range tetangga {
				iterative_deepening_search(node, link_tujuan, depth-1, hasil)
			}
		}
	}
}

/*
 * Function to perform Breadth First Search (BFS) algorithm with one solution
 *
 * @param current_node: the current node (current link)
 * @param link_tujuan: title of the destination link
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
func breadth_first_search_one_solution(current_node Node, link_tujuan string, hasil *[]Node) {
	// queue to store the adjacent links
	queue := getAdjacentLinks(current_node)

	for len(queue) != 0 {
		// get the first element in the queue and remove it from the queue
		current_node := queue[0]
		queue = queue[1:]

		// For Debugging, delete when it's done
		fmt.Println(current_node.Current)
		fmt.Println(current_node.Paths)

		// check if the current node is the destination node
		if current_node.Current == link_tujuan {
			*hasil = append(*hasil, current_node)
			break
		} else { // if the current node is not the destination node
			queue = append(queue, getAdjacentLinks(current_node)...)
		}
	}
}

/*
 *
 *Function to perform Breadth First Search (BFS) algorithm with multiple solutions
 *
 * @param current_node: the current node (current link)
 * @param link_tujuan: title of the destination link
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 * @param minimum_lengt_solution: the minimum length of the solution
 */
func breadth_first_search_mant_solution(current_node Node, link_tujuan string, hasil *[]Node, minimum_lengt_solution *int) {
	// queue to store the adjacent links
	queue := getAdjacentLinks(current_node)

	for len(queue) != 0 {
		// get the first element in the queue and remove it from the queue
		current_node := queue[0]
		queue = queue[1:]

		// For Debugging, delete when it's done
		fmt.Println(current_node.Current)
		// fmt.Println(current_node.Paths)

		// check if the current node is the destination node
		if current_node.Current == link_tujuan {
			*hasil = append(*hasil, current_node)
			if len(current_node.Paths) < *minimum_lengt_solution {
				*minimum_lengt_solution = len(current_node.Paths)
			}
		} else if len(current_node.Paths) > *minimum_lengt_solution { // if the current path is longer than the minimum length of the solution
			return
		} else { // if the current node is not the destination node
			queue = append(queue, getAdjacentLinks(current_node)...)
		}
	}
}
