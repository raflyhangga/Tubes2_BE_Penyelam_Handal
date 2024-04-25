package scrapper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	Paths   *Node
	// Neighbours []string
}

// Define a struct to represent a cache format in file
type CacheFile struct {
	Key        string   `json:"key"`
	Neighbours []string `json:"Node"`
}

/**
 *    GLOBAL VARIABLES
 *
 */
var DOMAIN_PREFIX string = "https://en.wikipedia.org"
var THREADS int = 100
var Total_Visited_Link int = 0

// Define a map to keep track of node (link) that has been added to the queue/stack
// if the node is added, the value is true, otherwise false
var Visited_Node = make(map[string]bool)

// Cache
var Cache = make(map[string][]string)

/**
 *    FUNCTION
 *
 */

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
func getAdjacentLinks(activeNode Node) []string {
	// Get the HTML document from the current link
	link := DOMAIN_PREFIX + activeNode.Current
	res := getResponse(link)
	doc := getDocument(res)

	// If the document is nil, return an empty list
	if doc != nil {
		unique := make(map[string]bool) // to make sure that the link is unique

		// Find all the links in the HTML document
		var link_list []string
		// Filter the links that start with "/wiki/" and do not contain "."
		doc.Find("div.mw-content-ltr.mw-parser-output").Find("a").FilterFunction(func(i int, s *goquery.Selection) bool {
			link_temp, _ := s.Attr("href")
			return strings.HasPrefix(link_temp, "/wiki") && !(strings.Contains(link_temp, "."))
		}).Each(func(i int, s *goquery.Selection) {
			link_temp, _ := s.Attr("href")
			// Create a new node for each link
			link_processed := removeHash(link_temp)

			// Append the new node to the list of links
			if !Visited_Node[link_processed] && !unique[link_processed] {
				link_list = append(link_list,  removeHash(link_processed))
				unique[link_processed] = true
			}
		})
		return link_list
	} else {
		return nil
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

/**
 * Function to write the cache to a JSON file
 */
func writeCacheToFile() {
	// write cache to struct
	var cacheFile []CacheFile
	for key, value := range Cache {
		var temp CacheFile
		temp.Key = key
		temp.Neighbours = append(temp.Neighbours, value...)
		cacheFile = append(cacheFile, temp)
	}

	// Marshal data slice into JSON
	jsonData, err := json.MarshalIndent(cacheFile, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// write cache to file
	file, err := os.Create("cache.json")
	if err != nil {
		fmt.Println("Failed to create cache file")
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing cache to file:", err)
		return
	}

	fmt.Println("Cache has been written to cache.json")
}

/**
 * Function to read the cache from a JSON file
 */
func readCacheFromFile() {
	// read cache from file
	file, err := os.Open("cache.json")
	if err != nil {
		fmt.Println("Failed to open cache file")
		return
	}
	defer file.Close()

	// Unmarshal JSON data into a slice of structs
	var cacheFile []CacheFile
	err = json.NewDecoder(file).Decode(&cacheFile)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// write cache to map
	for _, value := range cacheFile {
		var temp []string
		temp = append(temp,value.Neighbours...)
		Cache[value.Key] = temp
	}

	fmt.Println("Cache has been read from cache.json")
}



func AddDomainPrefix(list []string) []string {
	realLink := make([]string, len(list))
	for i, link := range list {
		realLink[i] = DOMAIN_PREFIX + link
	}
	return realLink
}

func reversePath(arr []string) []string {
	n := len(arr)
	reversed := make([]string, n)
	for i := 0; i < n; i++ {
		reversed[i] = arr[n-i-1]
	}
	return reversed
}