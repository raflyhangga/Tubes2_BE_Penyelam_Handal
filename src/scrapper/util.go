package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	// "sync"

	"time"

	"github.com/PuerkitoBio/goquery"
)

var DOMAIN_PREFIX string = "https://en.wikipedia.org"
var THREADS int = 100

// Define a struct to represent a node in the graph
type Node struct {
	Current 	string		 
	Paths   	[]string	 
}

// Define a map to keep track of node (link) that has been added to the queue/stack
// if the node is added, the value is true, otherwise false
var visitedNode = make(map[string]bool)

// Cache
type LinkData struct {
	Link string
	Neighbours []string
}
var cache = make(map[string][]string)
var temp_cache []LinkData


func loadCache(){
	for i:=0; i<len(temp_cache);i++ {
		link := temp_cache[i].Link
		neighbour := temp_cache[i].Neighbours
		if(len(cache[link]) == 0) {
			cache[link] = neighbour
		}
	}
}

func writeCache(link string, neighbour []string){
	temp_cache = append(temp_cache, LinkData{
		Link: link,
		Neighbours: neighbour,
	})
}

func isInCache(link string) bool {
	return !(len(cache[link]) == 0)
}

func getCache(link string) []string {
	return cache[link]
}

func getCacheToNode(main_link Node) []Node {
	var list_node []Node
	for _,link := range getCache(main_link.Current) {
		var temp_Node = Node{
			Current: link,
			Paths: append(main_link.Paths,main_link.Current),
		}
		list_node = append(list_node, temp_Node)
	}
	return list_node
}

// Function to get the response from the link
func getResponse(link string) *http.Response {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal("Failed to connect to designated page", err)
	}

	if res.StatusCode != 200 && res.StatusCode != 429 {
		//log.Fatalf("HTTP Error %d: %s", res.StatusCode, res.Status)
		fmt.Println("HTTP Error", res.StatusCode, ":", res.Status)
		return nil
	} else if res.StatusCode == 429 {
		fmt.Println("Too many requests, please wait for a while")
		time.Sleep(2 * time.Second)
		return getResponse(link)
	} else {
		return res
	}
}

// Function to get HTML document from the response
func getDocument(resp *http.Response) *goquery.Document {
	if resp == nil {
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Failed to parse the HTML document", err)
		return nil
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
	if(!isInCache(active_node.Current)) {
		// Get the HTML document from the current link
		link := active_node.Current
		var links_string []string
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
			var tempNode =  Node{
				Current : removeHash(DOMAIN_PREFIX + linkTemp),
				Paths : append(active_node.Paths, removeHash(active_node.Current)),
			}
			// Append the new node to the list of links
			if !visitedNode[tempNode.Current] && !unique[tempNode.Current] {
				links = append(links, tempNode)
				links_string = append(links_string, tempNode.Current)
				unique[tempNode.Current] = true
			}
		})
		writeCache(active_node.Current,links_string)
		return links
	} else {
		return getCacheToNode(active_node)
	}
}

func removeHash(link string) string {
	lastIndex := strings.LastIndexByte(link, '#')
	if lastIndex < 0 {
		return link
	}
	return link[:lastIndex]
}
