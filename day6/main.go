package main

import (
	"fmt"
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

func main() {
	input := advent.ReadStringArrayInput()

	fmt.Printf("%d\n", totalOrbits(input))
}
