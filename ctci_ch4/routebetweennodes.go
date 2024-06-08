package ch4

// Route Between Nodes: Given a directed graph, design an algorithm
// to find out whether there is a route between two nodes

// GraphNode represents a node in the graph
type GraphNode struct {
	data	string
	children []*GraphNode
	visited bool
}


// Graph represents the datastructure as a whole
type Graph struct {
	nodes []*GraphNode
}


func (g *Graph) CreateGraphNode(data string) *GraphNode {

	n := &GraphNode{data: data}
	n.children = make([]*GraphNode, 0)

	g.nodes = append(g.nodes, n)

	return n

}


func addEdge(src, target *GraphNode) {
	src.children = append(src.children, target)
}


func (g *Graph) clearVisits() {

	for _, node := range g.nodes {
		node.visited = false
	}

}


// depth first search
func dfs(src, target *GraphNode) bool {
	
	if src == target {
		return true
	}

	src.visited = true

	// answer
	var ans bool

	for _, child := range src.children {
		if !child.visited {
			ans = dfs(child, target)
			if ans == true {
				return ans
			}
		}
	}

	return ans

}


// breadth first search
func bfs(src *GraphNode, dst *GraphNode) bool {

	if src == dst {
		return true
	}

	queue := []*GraphNode{src}
	src.visited = true

	for len(queue) != 0 {
		node := queue[0]
		if node == dst {
			return true
		}

		queue = queue[1:]

		for _, child := range node.children {
			if !child.visited {
				child.visited = true
				queue = append(queue, child)
			}
		}
	}
	
	return false
}






