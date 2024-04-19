package main
import (
	"fmt"
	"tes"
	
)

// // import "fmt"

// // // 	import "fmt"

// // func main() {
// // 	var x int
// // 	fmt.Print("Enter a number: ")
// // 	fmt.Scan(&x)
// // 	fmt.Println(x)

// // }

// run
// go run main.go
package main

import (
	"fmt"
	"net/html"
	"net/http"
	"strings"
)

func main() {
	// URL yang ingin di-scraping
	url := "https://en.wikipedia.org/wiki/Adolf_Hitler"

	// Lakukan HTTP GET request untuk mengambil halaman web
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error saat melakukan HTTP GET request:", err)
		return
	}
	defer response.Body.Close()

	// Parsing HTML dari body response
	doc, err := html.Parse(response.Body)
	if err != nil {
		fmt.Println("Error saat parsing HTML:", err)
		return
	}

	// Rekursif untuk mencari dan mengekstrak link
	var links []string
	// representasi node nya adalah left child, right sibling
	var extractLinks func(*html.Node)
	// Fungsi rekursif untuk mengekstrak link dari node HTML
	extractLinks = func(n *html.Node) {
		// Cari tag <a> dan ambil atribut href-nya
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				// Jika atribut href ditemukan, maka simpan ke dalam slice links
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		// Rekursif untuk mengecek child node dari node saat ini
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	// Panggil fungsi extractLinks untuk mencari link
	extractLinks(doc)

	// Cetak link yang ditemukan
	// var i int = 1
	for indeks, link := range links {
		// Filter link agar hanya menampilkan link yang valid
		//if strings.HasPrefix(link, "/wiki") {
		fmt.Print(indeks + 1)
		fmt.Print(" : ")
		if !strings.HasPrefix(link, "https://") {
			fmt.Print("https://en.wikipedia.org")
		}
		fmt.Println(link)

		if strings.Contains(link, "printable") {
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println("printable")
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
		}
		// if strings.Contains(link, "Toothbrush_moustache")
		// if indeks == 999 {
		// 	break
		// }

	}
}

// // package main

// // import (
// // 	"fmt"
// // 	"io/ioutil"
// // 	"net/http"
// // )

// // func main() {
// // 	// Tentukan URL yang akan diambil HTML-nya
// // 	url := "https://en.wikipedia.org/wiki/Adolf_Hitler#CITEREFMcMillan2012"

// // 	// Lakukan HTTP GET request ke URL
// // 	response, err := http.Get(url)
// // 	if err != nil {
// // 		fmt.Println("Error saat melakukan GET request:", err)
// // 		return
// // 	}
// // 	defer response.Body.Close()

// // 	// Baca isi dari response body
// // 	body, err := ioutil.ReadAll(response.Body)
// // 	if err != nil {
// // 		fmt.Println("Error saat membaca response body:", err)
// // 		return
// // 	}

// // 	// Tampilkan isi HTML
// // 	fmt.Println("Isi HTML dari", url, "adalah:")
// // 	fmt.Println(string(body))
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"io/ioutil"
// // 	"net/http"
// // )

// // func main() {
// // 	// Tentukan URL yang akan diambil HTML-nya
// // 	url := "https://en.wikipedia.org/wiki/Adolf_Hitler"

// // 	// Lakukan HTTP GET request ke URL
// // 	response, err := http.Get(url)
// // 	if err != nil {
// // 		fmt.Println("Error saat melakukan GET request:", err)
// // 		return
// // 	}
// // 	defer response.Body.Close()

// // 	// Baca isi dari response body
// // 	body, err := ioutil.ReadAll(response.Body)
// // 	if err != nil {
// // 		fmt.Println("Error saat membaca response body:", err)
// // 		return
// // 	}

// // 	// Simpan isi HTML ke dalam sebuah file
// // 	err = ioutil.WriteFile("contohwebsite.html", body, 0644)
// // 	if err != nil {
// // 		fmt.Println("Error saat menyimpan file:", err)
// // 		return
// // 	}

// // 	fmt.Println("File HTML telah berhasil disimpan sebagai contohwebsite.html")
// // }
