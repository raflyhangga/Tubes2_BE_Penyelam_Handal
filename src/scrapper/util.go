package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	// "sync"

	"github.com/PuerkitoBio/goquery"
)

var DOMAIN_PREFIX string = "https://en.wikipedia.org"
var total_nodes int = 0

// Define a struct to represent a node in the graph
type Node struct {
	Current    string // The current link (current node)
	Neighbours []Node // The path from the start node to the current node
	Paths      []string
}

const (
	MAX_THREAD int = 100
	MAX_DEPTH  int = 3
)

// func constructGraph(before []string, link string, depth int) Node {
// 	fmt.Println("construct : ", link)
// 	if depth == 0 {
// 		return Node{
// 			Current: link,
// 			Paths:   before,
// 		}
// 	} else {
// 		tetangga := getAdjacentLinks(link)
// 		var arrayNode []Node

// 		// var wg sync.WaitGroup

// 		// for len(tetangga) > MAX_THREAD {
// 		// 	wg.Add(MAX_THREAD)
// 		// 	queue_100_elmt := append([]string{}, tetangga[:MAX_THREAD]...)
// 		// 	tetangga = tetangga[MAX_THREAD:]
// 		// 	for _,v := range queue_100_elmt {
// 		// 		go func(v string){
// 		// 			defer wg.Done()
// 		// 			arrayNode = append(arrayNode, constructGraph(append(before,link),v,depth-1))
// 		// 		}(v)
// 		// 	}
// 		// 	wg.Wait()
// 		// }

// 		// wg.Add(len(tetangga))
// 		// for _,v := range tetangga {
// 		// 	go func(v string){
// 		// 		defer wg.Done()
// 		// 		arrayNode = append(arrayNode, constructGraph(append(before,link),v,depth-1))
// 		// 	}(v)
// 		// }
// 		// wg.Wait()

// 		for _, v := range tetangga {
// 			arrayNode = append(arrayNode, constructGraph(append(before, link), v, depth-1))
// 		}

// 		return Node{
// 			Current:    link,
// 			Neighbours: arrayNode,
// 			Paths:      before,
// 		}
// 	}
// }

func constructGraph(before []string, link string, depth int, wg *sync.WaitGroup) Node {
	defer wg.Done()
	fmt.Println("construct : ", link)
	if depth == 0 {
		return Node{Current: link, Paths: before}
	}

	tetangga := getAdjacentLinks(link)
	var childWG sync.WaitGroup
	var arrayNode []Node

	for _, v := range tetangga {
		childWG.Add(1)
		go func(v string) {
			defer childWG.Done()
			arrayNode = append(arrayNode, constructGraph(append(before, link), v, depth-1, &childWG))
		}(v)
	}

	childWG.Wait()
	return Node{Current: link, Neighbours: arrayNode, Paths: before}
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
func getAdjacentLinks(link string) []string {
	// Get the HTML document from the current link
	res := getResponse(link)
	doc := getDocument(res)
	unique := make(map[string]bool)
	// Find all the links in the HTML document
	var links []string
	// Filter the links that start with "/wiki" and do not contain ":" or "#"
	doc.Find("div.mw-content-ltr.mw-parser-output").Find("a").FilterFunction(func(i int, s *goquery.Selection) bool {
		linkTemp, _ := s.Attr("href")
		return strings.HasPrefix(linkTemp, "/wiki")
	}).Each(func(i int, s *goquery.Selection) {
		linkTemp, _ := s.Attr("href")
		// Create a new node for each link
		tempLink := removeHash(DOMAIN_PREFIX + linkTemp)
		// Append the new node to the list of links
		if !visitedNode[tempLink] && !unique[tempLink] {
			links = append(links, tempLink)
			unique[tempLink] = true
		}
	})
	return links
}

func removeHash(link string) string {
	lastIndex := strings.LastIndexByte(link, '#')
	if lastIndex < 0 {
		return link
	}
	return link[:lastIndex]
}
