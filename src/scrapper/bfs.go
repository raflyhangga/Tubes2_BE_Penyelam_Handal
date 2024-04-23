package scrapper

import (
	"sync"
)

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
		isDestinationLinkExist := false
		for _, current_node := range current_queue {
			visitedNode[current_node.Current] = true
			if current_node.Current == destination_link {
				*hasil = append(*hasil, current_node)
				isDestinationLinkExist = true
			}
		}

		if isDestinationLinkExist {
			// solution is found, return
			return
		} else {
			// if the current queue does not have the destination node
			// create a goroutine for each adjacent link to concurrently get the adjacent links
			var wg sync.WaitGroup

			// make a new queue to store the adjacent links from the current queue
			var new_queue []Node

			// process the first 100 links in the current queue
			for len(current_queue) > THREADS {
				// limit the number of goroutines to 100 to avoid spam http request
				wg.Add(THREADS)
				queue_THREADS_elmt := append([]Node{}, current_queue[:THREADS]...)
				current_queue = current_queue[THREADS:]
				for _, current_node := range queue_THREADS_elmt {
					go func(current_node Node) {
						defer wg.Done()
						// get the adjacent links from the current node
						adjacent_links := getAdjacentLinks(current_node)
						new_queue = append(new_queue, adjacent_links...)
					}(current_node)
				}
				wg.Wait()
			}

			// process the remaining links
			wg.Add(len(current_queue))
			for _, current_node := range current_queue {
				go func(current_node Node) {
					defer wg.Done()
					// get the adjacent links from the current node
					adjacent_links := getAdjacentLinks(current_node)
					new_queue = append(new_queue, adjacent_links...)
				}(current_node)
			}

			// wait for all goroutines to finish
			wg.Wait()

			// recursively call the breadth_first_search function with the new queue
			breadth_first_search(new_queue, destination_link, hasil)
		}
	}
}
