package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var DOMAIN_PREFIX string = "https://en.wikipedia.org"

// Define a struct to represent a node in the graph
type Node struct {
	Current string   // The current link (current node)
	Paths   []string // The path from the start node to the current node
}

// Define a map to keep track of node (link) that has been added to the queue/stack
// if the node is added, the value is true, otherwise false
var visitedNode = make(map[string]bool)

// Function to get the response from the link
func getResponse(link string) *http.Response {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal("Failed to connect to designated page", err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("HTTP Error %d: %s", res.StatusCode, res.Status)
	}

	return res
}

// Function to get HTML document from the response
func getDocument(resp *http.Response) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Failed to parse the HTML document", err)
	}
	return doc
}

/**
 * Function to get the adjacent nodes (links) from the current node (current link)
 *
 * @param active_node: the current node
 * @return links: a list of adjacent nodes
 */
func getAdjacentLinks(active_node Node) []Node {
	// Get the HTML document from the current link
	link := active_node.Current
	res := getResponse(link)
	doc := getDocument(res)
	unique := make(map[string]bool)
	// Find all the links in the HTML document
	var links []Node
	// Filter the links that start with "/wiki" and do not contain ":" or "#"
	doc.Find("a").FilterFunction(func(i int, s *goquery.Selection) bool {
		linkTemp, _ := s.Attr("href")
		return strings.HasPrefix(linkTemp, "/wiki") //&& !(strings.Contains(linkTemp, ":") || strings.Contains(linkTemp, "#"))
	}).Each(func(i int, s *goquery.Selection) {
		linkTemp, _ := s.Attr("href")
		// fmt.Println(DOMAIN_PREFIX + linkTemp)
		// Create a new node for each link
		var tempNode Node
		tempNode.Current = DOMAIN_PREFIX + linkTemp
		tempNode.Paths = append(active_node.Paths, active_node.Current)
		// Append the new node to the list of links
		if !visitedNode[tempNode.Current] && !unique[tempNode.Current] {
			links = append(links, tempNode)
			unique[tempNode.Current] = true
		}
	})
	return links
}

var total_nodes int = 0

func breadth_first_search(current_queue []Node, destination_link string, hasil *[]Node) {
	if len(current_queue) == 0 {
		return
	} else {
		// go through all the nodes in the current queue
		// check if the destination node is in the current queue
		isDestinationLinkExist := false
		for _, current_node := range current_queue {
			visitedNode[current_node.Current] = true
			total_nodes++
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
			for len(current_queue) > 100 {
				wg.Add(100)
				queue_100_elmt := append([]Node{}, current_queue[:100]...)
				current_queue = current_queue[100:]
				for _, current_node := range queue_100_elmt {
					go func(current_node Node) {
						defer wg.Done()
						// get the adjacent links from the current node
						adjacent_links := getAdjacentLinks(current_node)
						new_queue = append(new_queue, adjacent_links...)
						fmt.Println(new_queue[0].Current)
						fmt.Println(new_queue[len(new_queue)-1].Current)
						fmt.Println(len(new_queue))
					}(current_node)
				}
				wg.Wait()
			}

			// process the remaining links in the current queue
			wg.Add(len(current_queue))
			for _, current_node := range current_queue {
				go func(current_node Node) {
					defer wg.Done()
					// get the adjacent links from the current node
					adjacent_links := getAdjacentLinks(current_node)
					new_queue = append(new_queue, adjacent_links...)
					fmt.Println(new_queue[0].Current)
					fmt.Println(new_queue[len(new_queue)-1].Current)
					fmt.Println(len(new_queue))
				}(current_node)
			}

			// wait for all goroutines to finish
			wg.Wait()

			// recursively call the breadth_first_search function with the new queue
			breadth_first_search(new_queue, destination_link, hasil)
		}
	}
}
