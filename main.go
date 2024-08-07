package main

import (
	"fmt"
	"os"

	"github.com/Sena543/compression/cmd"
)

func main() {
	fmt.Println("Compression tool")
	testString := "BCAADDDCCACACAC"
	res := cmd.NewEncode(testString)
	var pq cmd.PriorityQueue
	for k, v := range res {
		node := cmd.Node{Data: k, Count: v}
		pq.Insert(node)
	}
	var hTree cmd.HuffNode
	rootNode := hTree.BuildTree(pq)
	/* rootNode := hTree.BuildTree(pq.PQ(), 0) */
	hTree.AssignWeights(rootNode)
	hTree.PrintTree(rootNode, os.Stdout)

}
