# A* Pathfinding in Go

This is a simple A* pathfinding solution for Go.

# Installation

	go get github.com/massung/pathfind

# Usage

	import "github.com/massung/pathfind"

# Quickstart

The `pathfind` package works by having your data structure implement the `Graph` interface, which is comprised of two functions:

	type Graph interface {
		Neighbors(Node) []Edge
		H(Node, Node) float32
	}

The first of these is a function that - given a `Node` (an `interface{}`) - will return a list of `Edge` values that represent neighboring nodes that can be traversed to. The second is a heuristic function. Given a from `Node` and a goal `Node`, the `H` function should return an approximate cost to traverse that are.

Once you have implemented these two functions, then you should be able to call the `Search` function and get back a list of `Node` values that comprise the optimal path to take.
	
# Example (a 2D Map)

Since many times A* pathfinding is used for 2D games, this will be an example of using the `pathfind` module for walking a 2D map.

	// the world will be our node graph
	type World struct {
		Tiles [10][10]Tile
	}
	
	// a simple definition for a node in the graph
	type Tile struct {
		X, Y int
		Cost float32
	}
	
	// create a simple heuristic to estimate the cost from A to B
	func (w *World) H(from, goal pathfind.Node) float32 {
		dx := goal.(*Tile).X - from.(*Tile).X
		dy := goal.(*Tile).Y - from.(*Tile).Y
		
		// distance squared...
		return float32(dx * dx + dy * dy)
	}
	
	// return an edge set of neighbors for a given tile
	func (w *World) Neighbors(node pathfind.Node) []pathfind.Edge {
		edges := make([]pathfind.Edge, 0, 8)
		tile := node.(*Tile)
		
		// look in the 4 cardinal directions...
		if tile.X > 0 { 
			edge := pathfind.Edge{
				Node: &w.Tiles[tile.X - 1][tile.Y],
				Cost: w.Tiles[tile.X - 1][tile.Y].Cost,
			}
			
			edges = append(edges, edge)
		}
		
		/* TODO: look at the other 3 directions as well... */
		
		return edges
	}
	
Now that we have our basic node graph, node definition, a heuristic function, and a function that can return a list of neighbors for any given "node" on the graph, we can now perform simple searches.

	// get a starting and ending node for the path
	start := &world.Tile[0][0]
	goal := &world.Tile[9][9]
	
	// perform the search...
	path, found := pathfind.Search(world, start, goal)
	
	// make sure there was a path found
	if found {
		for _, node := range path {
			tile := node.(*Tile)
			
			/* TODO: traverse each tile in the world... */
		}
	}