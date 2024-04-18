package src

import (
	"fmt"
)

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

		if(current_node.Current == link2){
			fmt.Println("Masuk if")
			*hasil = append(*hasil, current_node)
			break;
		} else {
			fmt.Println("Masuk Else")
			queue =append(queue, getAdjacentLinks(current_node)...)
		}
	}
}