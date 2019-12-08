package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/alecwest/advent-of-code-2019/advent"
)

type node struct {
	name     string
	children []string
}

func buildGraph(input []string) []node {
	var nodes []node
	for _, entry := range input {
		orbit := strings.Split(entry, ")")
		nodeIndex := getNodeIndex(nodes, orbit[0])
		nodeIndex2 := getNodeIndex(nodes, orbit[1])
		if nodeIndex == -1 {
			nodes = append(nodes, node{orbit[0], nil})
			nodeIndex = len(nodes) - 1
		}
		if nodeIndex2 == -1 {
			nodes = append(nodes, node{orbit[1], nil})
		}
		nodes[nodeIndex].children = append(nodes[nodeIndex].children, orbit[1])
	}
	return nodes
}

func breadthFirstSearch(nodes []node, nodeIndex int) int {
	queue := []string{}
	queue = append(queue, nodes[nodeIndex].name)
	nodesVisited := 0

	for len(queue) != 0 {
		currNodeIndex := getNodeIndex(nodes, queue[0])
		queue = queue[1:]
		currNode := nodes[currNodeIndex]
		for _, child := range currNode.children {
			queue = append(queue, child)
			nodesVisited++
		}
	}
	return nodesVisited
}

// return depth of node
func depthFirstSearch(nodes []node, startNode, target string) int {
	return dfs(nodes, target, getNodeIndex(nodes, startNode), 0)
}

// return depth of node
func dfs(nodes []node, target string, currNodeIndex, depth int) int {
	currNode := nodes[currNodeIndex]
	if currNode.name == target {
		return depth
	}
	for _, child := range nodes[currNodeIndex].children {
		result := dfs(nodes, target, getNodeIndex(nodes, child), depth+1)
		if result != -1 {
			return result
		}
	}
	return -1
}

func getNodeIndex(nodes []node, target string) int {
	nodeIndex := -1
	for i, n := range nodes {
		if n.name == target {
			nodeIndex = i
		}
	}
	return nodeIndex
}

func totalOrbits(input []string) int {
	graph := buildGraph(input)
	sum := 0
	for i := range graph {
		sum += breadthFirstSearch(graph, i)
	}
	return sum
}

func orbitalTransfers(input []string, nodeA, nodeB string) int {
	graph := buildGraph(input)
	minimumTransfer := math.MaxInt64

	for _, node := range graph {
		a := depthFirstSearch(graph, node.name, nodeA)
		b := depthFirstSearch(graph, node.name, nodeB)
		if a > 0 && b > 0 && a+b < minimumTransfer {
			minimumTransfer = a + b
		}
	}

	return minimumTransfer - 2
}

func main() {
	input := advent.ReadStringArrayInput()

	// fmt.Printf("%d\n", totalOrbits(input))
	fmt.Printf("%d\n", orbitalTransfers(input, "YOU", "SAN"))
}
