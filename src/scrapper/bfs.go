package scrapper

import (
	"fmt"
	"sync"
)

/**
 * Function to perform Breadth First Search (BFS) algorithm
 *
 * @param currentQueue: a list of nodes (links) to be processed
 * @param destinationLink: link of the destination link
 * @param hasil: a list of nodes that contain destination node
 */
func breadth_first_search(currentQueue []Node, destinationLink string, hasil *[]Node) {
	if len(currentQueue) == 0 {
		return
	} else {
		// go through all the nodes in the current queue
		// check if the destination node is in the current queue
		isDestinationLinkExist := false
		for _, current_node := range currentQueue {
			visitedNode[current_node.Current] = true
			if current_node.Current == destination_link {
				current_node.Paths = append(current_node.Paths, current_node.Current)
				*hasil = append(*hasil, current_node)
				isDestinationLinkExist = true
			}
		}
		totalVisitedLink += len(currentQueue)

		if isDestinationLinkExist {
			// solution is found, return
			return
		} else {
			// if the current queue does not have the destination node
			// create a goroutine for each link to concurrently get the adjacent links
			var wg sync.WaitGroup

			// make a new queue to store the adjacent links from the current queue
			var newQueue []Node

			// process queue in blocks of THREADS to avoid spam http request
			startIndeks := 0
			endIndeks := THREADS
			for endIndeks < len(currentQueue) {
				// limit the number of goroutines to THREADS to avoid spam http request
				wg.Add(THREADS)

				for i := startIndeks; i < endIndeks; i++ {
					go func(current_node Node) {
						defer wg.Done()
						// get the adjacent links from the current node
						adjacent_links := getAdjacentLinks(current_node)
						newQueue = append(newQueue, adjacent_links...)
					}(currentQueue[i])
				}

				wg.Wait()
				startIndeks += THREADS
				endIndeks += THREADS
			}

			// process the remaining links
			wg.Add(len(currentQueue) - startIndeks)
			for i := startIndeks; i < len(currentQueue); i++ {
				go func(current_node Node) {
					defer wg.Done()
					// get the adjacent links from the current node
					adjacent_links := getAdjacentLinks(current_node)
					newQueue = append(newQueue, adjacent_links...)
				}(currentQueue[i])
			}
			// wait for all goroutines to finish
			wg.Wait()

			// recursively call the breadth_first_search function with the new queue
			breadth_first_search(newQueue, destinationLink, hasil)
		}
	}
}
