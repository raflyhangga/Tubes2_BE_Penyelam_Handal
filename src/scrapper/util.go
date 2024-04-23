package scrapper

import (
	// "fmt"
	// "fmt"
	"log"
	"net/http"
	"strings"
	// "sync"
	
	"github.com/PuerkitoBio/goquery"
)

var DOMAIN_PREFIX string = "https://en.wikipedia.org"
var total_nodes int = 0
var THREADS int = 5

// Define a struct to represent a node in the graph
type Node struct {
	Current 	string		`json:"current"`   
	Paths   	[]string	`json:"paths"` 
	Depth 		int			`json:"depth"`
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
	doc.Find("div.mw-content-ltr.mw-parser-output").Find("a").FilterFunction(func(i int, s *goquery.Selection) bool {
		linkTemp, _ := s.Attr("href")
		return strings.HasPrefix(linkTemp, "/wiki")  && !(strings.Contains(linkTemp, "."))
	}).Each(func(i int, s *goquery.Selection) {
		linkTemp, _ := s.Attr("href")
		// fmt.Println(DOMAIN_PREFIX + linkTemp)
		// Create a new node for each link
		var tempNode Node
		tempNode.Current = removeHash(DOMAIN_PREFIX + linkTemp)
		tempNode.Paths = append(active_node.Paths, removeHash(active_node.Current))
		// Append the new node to the list of links
		if !visitedNode[tempNode.Current] && !unique[tempNode.Current] {
			links = append(links, tempNode)
			unique[tempNode.Current] = true
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
