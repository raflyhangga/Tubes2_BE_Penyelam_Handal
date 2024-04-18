package scrapper

import (
	"fmt"
	"time"
)

func printSolution(solutions []Node, duration time.Duration){
	fmt.Println("======================= SOLUTION =======================")
	fmt.Println("Found ",len(solutions)," solutions in" ,duration)
	for k,value :=range solutions {
		fmt.Println("Solution - ",k+1)
		for k,path := range value.Paths {
			fmt.Println(k+1,path)
		}
	}
}

func iterative_deepening_search(current_node Node,link_tujuan string, depth int, hasil *[]Node){
	if(depth >= 0){
		fmt.Println();
		if current_node.Current == link_tujuan{
			*hasil = append(*hasil, current_node)
		}
		tetangga := getAdjacentLinks(current_node)
		for _,node := range tetangga {
			iterative_deepening_search(node,link_tujuan,depth-1,hasil)
		}
	}
}

func breadth_first_search(current_node Node,link2 string, hasil *[]Node){
	queue := getAdjacentLinks(current_node)

	for len(queue) != 0 {
		current_node := queue[0]
		queue = queue[1:]

		fmt.Println(current_node.Current)
		fmt.Println(current_node.Paths)
		if(current_node.Current == link2){
			*hasil = append(*hasil, current_node)
			break;
		} else {
			queue =append(queue, getAdjacentLinks(current_node)...)
		}
	}
}

func IDS_interface(link_awal string, link_tujuan string, depth int){
	var initial Node
	initial.Current = link_awal

	var solutions []Node
	startTime:=time.Now()
	iterative_deepening_search(initial,link_tujuan,depth,&solutions)
	duration:=time.Since(startTime)

	printSolution(solutions,duration)

}

func BFS_interface(link_awal string, link_tujuan string){
	var initial Node
	initial.Current = link_awal

	var solutions []Node
	startTime:=time.Now()
	breadth_first_search(initial,link_tujuan,&solutions)
	duration:=time.Since(startTime)

	printSolution(solutions,duration)

}