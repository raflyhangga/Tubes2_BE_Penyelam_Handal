package main

import (
	"wikipedia-scraper-engine/scrapper"
)

func main(){
	start_link := "https://en.wikipedia.org/wiki/Indonesia"
	end_link := "https://en.wikipedia.org/wiki/Asia"
	scrapper.BFS_interface(start_link,end_link)
}