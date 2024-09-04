package main

import (
	"fmt"

	"github.com/Sena543/compression/cmd"
)

func main() {
	fmt.Println("Compression tool")
	testString := "BCAADDDCCACACAC"
	res := cmd.NewFrequencyMap(testString)
	var pq cmd.PriorityQueue
	for k, v := range res {
		node := cmd.Node{Data: k, Count: v}
		pq.Insert(node)
	}
	var hTree cmd.HuffNode
	rootNode := hTree.BuildTree(pq)
	/* rootNode := hTree.BuildTree(pq.PQ(), 0) */
	hTree.AssignWeights(rootNode)
	/* hTree.PrintTree(rootNode, os.Stdout) */

	encodeRes := hTree.EncodeStringMap(rootNode, res)
	fmt.Println("encoded string: ", hTree.EncodeString(testString, encodeRes), "mapEncoding: ", encodeRes)

}
