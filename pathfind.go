package pathfind

// a graph node
type Node interface{}

// a pathfinding graph
type Graph interface {
	// return a list of neighboring nodes to a given node
	Neighbors(Node) []Edge

	// the heuristic function to approximate the cost between nodes
	H(Node, Node) float32
}

// connections between nodes
type Edge struct {
	Node Node
	Cost float32
}

// 2D A* pathfinding node
type Path struct {
	Parent *Path
	Node Node
	f, g float32
}

// traverse a node graph for the optimal path
func Search(g Graph, start, goal Node) ([]Node, bool) {
	var closedSet []*Path

	// all the path nodes being searched
	openSet := []*Path{&Path{Node: start, g: 0, f: g.H(start, goal)}}

	// search the open set nodes for the best path
	for len(openSet) > 0 {
		current, path := bestNode(openSet)

		// have we reached the goal?
		if path.Node == goal {
			return nodeListOfPath(path), true
		}

		// remove current from the open set
		openSet[current] = openSet[len(openSet) - 1]
		openSet = openSet[:len(openSet) - 1]

		// add it to the closed set
		closedSet = append(closedSet, path)

		// get all the neighboring nodes to this one
		for _, e := range(g.Neighbors(path.Node)) {
			if nodeInSet(e.Node, closedSet) != nil {
				continue
			}

			// calculate the tentative g score for this node
			score := path.g + e.Cost
			h := g.H(e.Node, goal)
			p := nodeInSet(e.Node, openSet)

			// add the node to the final path
			if p == nil {
				p = &Path{
					Node: e.Node,
					Parent: path,
					g: score,
					f: score + h,
				}

				// add the new path node to the open set
				openSet = append(openSet, p)
			} else {
				// check to see if this path is better
				if score < p.g {
					p.Parent = path
					p.g = score
					p.f = score + h
				}
			}
		}
	}

	// no path found
	return []Node{}, false
}

// return the list of nodes for given path (from -> goal)
func nodeListOfPath(path *Path) []Node {
	if path.Parent == nil {
		return []Node{path.Node}
	}

	// recursively build the list
	return append(nodeListOfPath(path.Parent), path.Node)
}

// return a path node with the lowest f score from a set
func bestNode(set []*Path) (int, *Path) {
	best := 0

	// loop over the entire set of nodes
	for i, path := range(set[1:]) {
		if path.f < set[best].f {
			best = i
		}
	}

	return best, set[best]
}

// true if the node is in a path set
func nodeInSet(node Node, set []*Path) *Path {
	for _, path := range(set) {
		if path.Node == node {
			return path
		}
	}

	return nil
}
