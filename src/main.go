package main

import (
	"wikipedia-scraper-engine/scrapper"
)

func main() {
	start_link := "https://en.wikipedia.org/wiki/Wet_Lake_(Kuyavia-Pomerania_Voivodeship)"
	end_link := "https://en.wikipedia.org/wiki/England"
	scrapper.BFS_interface(start_link, end_link)
}
