package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	adjacency map[string][]string
}

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

func (g Graph) CreatePath(firstNode, secondNode string) bool {
	visited := g.createVisited()
	var (
		path []string
		q    []string
	)
	q = append(q, firstNode)
	visited[firstNode] = true

	for len(q) > 0 {
		var currentNode string
		currentNode, q = q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adjacency[currentNode]
		if contains(edges, secondNode) {
			path = append(path, secondNode)
			fmt.Println(strings.Join(path, "->"))
			return true
		}

		for _, node := range g.adjacency[currentNode] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}
	fmt.Println("no link found")
	return false
}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (g Graph) DFS(startingNode string) {
	visited := g.createVisited()
	g.dfsRecursive(startingNode, visited)
}

func (g Graph) dfsRecursive(startingNode string, visited map[string]bool) {
	visited[startingNode] = true
	fmt.Println("Node: ", startingNode)
	for _, node := range g.adjacency[startingNode] {
		if !visited[node] {
			g.dfsRecursive(node, visited)
		}
	}
}

func main() {
	fmt.Println("--- Depth First Search ---")

	g := NewGraph()

	g.AddVertex("Dhaka")
	g.AddVertex("Savar")
	g.AddVertex("Cumilla")
	g.AddVertex("Narsingdi")
	g.AddVertex("Sylhet")

	g.AddEdge("Dhaka", "Savar")
	g.AddEdge("Savar", "Dhaka")
	g.AddEdge("Dhaka", "Narsingdi")
	g.AddEdge("Narsingdi", "Dhaka")
	g.AddEdge("Narsingdi", "Sylhet")
	g.AddEdge("Sylhet", "Narsingdi")
	g.AddEdge("Cumilla", "Dhaka")

	g.DFS("Savar")
	g.CreatePath("Savar", "Sylhet")
}
