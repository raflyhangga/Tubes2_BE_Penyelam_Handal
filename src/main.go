package main

import (
	"wikipedia-scraper-engine/scrapper"
)

func main(){
	start_link := "https://en.wikipedia.org/wiki/Adolf_Hitler"
	end_link := "https://en.wikipedia.org/wiki/Dinosaur"
	scrapper.BFS_interface(start_link,end_link)
}
