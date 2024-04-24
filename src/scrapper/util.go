package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/**
 *    DATA STRUCTURE
 *
 */
// Define a struct to represent a node in the graph
type Node struct {
	Current string
	Paths   []string
}

// Cache
type LinkData struct {
	Link       string
	Neighbours []string
}

/**
 *    GLOBAL VARIABLES
 *
 */
var DOMAIN_PREFIX string = "https://en.wikipedia.org"
var THREADS int = 100
var Total_Visited_Link int = 0
var LOCK = false

// Define a map to keep track of node (link) that has been added to the queue/stack
// if the node is added, the value is true, otherwise false
var Visited_Node = make(map[string]bool)

// Cache
var Cache = make(map[string][]Node)

/**
 * Function to check if the link is in the Cache
 *
 * @param link: the link
 * @return true if the link is in the Cache, otherwise false
 */
func isInCache(link string) bool {
	return len(Cache[link]) != 0
}

/**
 * Function to get the response from the link
 *
 * @param link: the link
 * @return the response
 */
func getResponse(link string) *http.Response {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal("Failed to connect to designated page", err)
	}

	if res.StatusCode == 200 { // if the request is successful
		return res
	} else if res.StatusCode == 429 { // if the request is too many
		fmt.Println("Too many requests, please wait for a while")
		time.Sleep(2 * time.Second)
		return getResponse(link)
	} else { // if the request is failed
		fmt.Println("HTTP Error", res.StatusCode, ":", res.Status)
		return nil
	}
}

/**
 * Function to get the HTML document from the response
 *
 * @param resp: the response
 * @return the HTML document
 */
func getDocument(resp *http.Response) *goquery.Document {
	if resp == nil {
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Failed to parse the HTML document,", err)
		return nil
	}
	return doc
}

/**
 * Function to get the adjacent nodes (links) from the current node (current link)
 *
 * @param activeNode: the current node
 * @return links: a list of adjacent nodes
 */
func getAdjacentLinks(activeNode Node) []Node {
	// Get the HTML document from the current link
	link := DOMAIN_PREFIX + activeNode.Current
	res := getResponse(link)
	doc := getDocument(res)

	// If the document is nil, return an empty list
	if doc == nil {
		return []Node{}
	} else {
		unique := make(map[string]bool) // to make sure that the link is unique

		// Find all the links in the HTML document
		var linkNodes []Node
		// Filter the links that start with "/wiki/" and do not contain "."
		doc.Find("div.mw-content-ltr.mw-parser-output").Find("a").FilterFunction(func(i int, s *goquery.Selection) bool {
			linkTemp, _ := s.Attr("href")
			return strings.HasPrefix(linkTemp, "/wiki") && !(strings.Contains(linkTemp, "."))
		}).Each(func(i int, s *goquery.Selection) {
			linkTemp, _ := s.Attr("href")
			// Create a new node for each link
			var tempNode = Node{
				Current: removeHash(linkTemp),
				Paths:   append(activeNode.Paths, removeHash(activeNode.Current)),
			}

			// Append the new node to the list of links
			if !Visited_Node[tempNode.Current] && !unique[tempNode.Current] {
				linkNodes = append(linkNodes, tempNode)
				unique[tempNode.Current] = true
			}
		})
		return linkNodes
	}
}

/**
 * Function to remove the substring after the hash
 *
 * @param link: the link
 * @return the link without the hash
 */
func removeHash(link string) string {
	lastIndex := strings.LastIndexByte(link, '#')
	if lastIndex < 0 {
		return link
	}
	return link[:lastIndex]
}
