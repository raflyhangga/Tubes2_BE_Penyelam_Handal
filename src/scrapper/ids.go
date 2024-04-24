package scrapper

/**
 * Function to perform Depth First Search (DFS) algorithm
 *
 * @param current_node: the current node (current link)
 * @param link_tujuan: title of the destination link
 * @param depth: the maximum depth to search
 * @param hasil: a list of nodes that contain the path from the start node to the destination node
 */
 func depth_limited_search(current_node Node,link_tujuan string, depth int, hasil *[]Node){
	if(depth >= 0){
		if current_node.Current == link_tujuan{
			current_node.Paths = append(append(current_node.Paths, current_node.Current),link_tujuan)
			*hasil = append(*hasil, current_node)
		} else {
			tetangga := getAdjacentLinks(current_node)
			for _,node := range tetangga {
				if node.Current == link_tujuan {
					current_node.Paths = append(append(current_node.Paths, current_node.Current),link_tujuan)
					*hasil = append(*hasil, current_node)
				} else {
					depth_limited_search(node,link_tujuan,depth-1,hasil)
				}
			}
		}
	}
}

func iterative_deepening_search(link_awal string, link_tujuan string,hasil *[]Node){
	var initial Node
	initial.Current = link_awal

	var solutions []Node
	depth := 0

	for (len(solutions) == 0) {
		depth_limited_search(initial, link_tujuan, depth,&solutions)
		depth++
	}

	*hasil = solutions
}