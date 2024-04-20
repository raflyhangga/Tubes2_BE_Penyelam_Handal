package main

import (
	"fmt"
	"time"
)

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for idx := 1; idx <= total; idx++ {
// 		fmt.Printf("sending %d to channel\n", idx)
// 		ch <- idx
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for num := range ch {
// 		fmt.Printf("read %d from channel\n", num)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numberChan := make(chan int)

// 	wg.Add(2)
// 	go printNumbers(numberChan, &wg)

// 	generateNumbers(3, numberChan, &wg)

// 	close(numberChan)

// 	fmt.Println("Waiting for goroutines to finish...")
// 	wg.Wait()
// 	fmt.Println("Done!")
// }

// package main

func main() {
	a := "https://en.wikipedia.org/wiki/Duck"
	b := "https://en.wikipedia.org/wiki/Chicken"
	// c := "https://en.wikipedia.org/wiki/Pan-Germanism"
	// d := "https://en.wikipedia.org/wiki/Antisemitism"
	// e := "https://en.wikipedia.org/wiki/Anti-communism"
	// f := "https://en.wikipedia.org/wiki/Charismatic_authority"
	NodeA := Node{Current: a, Paths: []string{a}}
	// NodeC := Node{Current: c, Paths: []string{c}}
	// NodeD := Node{Current: d, Paths: []string{d}}
	// NodeE := Node{Current: e, Paths: []string{e}}
	// NodeF := Node{Current: f, Paths: []string{f}}
	// NodeB := Node{Current: b, Paths: []string{b}}
	hasil := []Node{}
	// breadth_first_search([]Node{NodeA}, b, &hasil)
	// fmt.Println(hasil)
	fmt.Println(len(hasil))
	start := time.Now()
	breadth_first_search([]Node{NodeA}, b, &hasil)
	end := time.Since(start)

	for k, value := range hasil {
		fmt.Println("Solution - ", k+1)
		// for j, path := range value.Paths {
		// 	fmt.Println(j+1, path)
		// }
		// fmt.Println(value.Current)
		fmt.Println(value.Paths[2])
	}
	fmt.Println("Total waktu: ", end)
	fmt.Println("Total nodes: ", total_nodes)

	// wg.close()
}
