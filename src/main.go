package main

import (
	"wikipedia-scraper-engine/scrapper"
)

func main(){
	start_link := "https://en.wikipedia.org/wiki/Indonesia"
	end_link := "https://en.wikipedia.org/wiki/Asia"
	scrapper.IDS_interface(start_link,end_link,5)
}
