package scrapper

import (
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

var DOMAIN_PREFIX string = "https://en.wikipedia.org"

type Node struct {
	Current string
	Paths []string
}

var visitedMap =  make(map[string]bool);


func getResponse(link string) *http.Response{
	res,err := http.Get(link)
	if(err != nil){
		log.Fatal("Failed to connect to designated page",err)
	}
	if(res.StatusCode != 200){
		log.Fatalf("HTTP Error %d: %s",res.StatusCode,res.Status)
	}
	return res;
}

func getDocument(resp *http.Response) *goquery.Document {
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	if(err != nil){
		log.Fatal("Failed to parse the HTML document",err)
	}
	return doc
}

func getAdjacentLinks(active_node Node) []Node {
	link := active_node.Current
	res := getResponse(link)
	doc := getDocument(res)
	var links []Node
	doc.Find("div.mw-content-ltr.mw-parser-output").Find("a").FilterFunction(func(i int, s *goquery.Selection) bool {
		linkTemp,_ := s.Attr("href")
		return strings.HasPrefix(linkTemp,"/wiki") && !(strings.Contains(linkTemp,":") || strings.Contains(linkTemp,"#"))
	}).Each(func(i int, s *goquery.Selection) {
		linkTemp,_ := s.Attr("href")

		var tempNode Node
		tempNode.Current = DOMAIN_PREFIX+linkTemp
		tempNode.Paths = append(active_node.Paths,active_node.Current)
		if(!visitedMap[tempNode.Current]){
			links = append(links,tempNode)
			visitedMap[tempNode.Current] = true;
		}
	})
	return links
}