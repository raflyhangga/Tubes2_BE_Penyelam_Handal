package scrapper

/**
 * Function to perform Depth First Search (DFS) algorithm
 *
 * @param current_node: the current node (current link)
 * @param link_tujuan: title of the destination link
 * @param depth: the maximum depth to search
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
 func iterative_deepening_search(current_node Node,link_tujuan string, depth int, hasil *[]Node){
	// fmt.Println(current_node.Current)
	if(depth > 0){
		if current_node.Current == link_tujuan{
			*hasil = append(*hasil, current_node)
		} else {
			tetangga := getAdjacentLinks(current_node)
			for _,node := range tetangga {
				iterative_deepening_search(node,link_tujuan,depth-1,hasil)
				if(len(*hasil) != 0) {
					return
				}
			}
		}
	}
}





// func iterative_deepening_search(current_node Node,link_tujuan string, depth int, hasil *[]Node) {
// 	fmt.Println(current_node.Current);
// 	if(depth == 0){
// 		if current_node.Current == link_tujuan{
// 			*hasil = append(*hasil, current_node)
// 		}
// 	} else {
// 		if current_node.Current == link_tujuan{
// 			*hasil = append(*hasil, current_node)
// 		} else {
// 			for _,node := range current_node.Neighbours {
// 				if(len(cache[node.Current]) == 0) {
// 					getLinkz(&node)
// 					cache[node.Current] = node.Neighbours
// 				} else {
// 					node.Neighbours = cache[node.Current]
// 				}
// 				iterative_deepening_search(node,link_tujuan,depth-1,hasil)
// 			}
// 		}
// 	}
// }