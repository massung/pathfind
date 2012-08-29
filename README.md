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

*TODO: create the example*