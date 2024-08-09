package cmd

import (
	"fmt"
	"io"
)

type HuffNode = Node

func NewEncode(input string) map[string]int {
	var result = make(map[string]int)
	for _, val := range input {
		result[string(val)] += 1
	}
	return result
}

func (h *HuffNode) BuildTree(pq PriorityQueue) *HuffNode {

	for len(pq.queueArray) > 1 {
		left := pq.Pop()
		right := pq.Pop()

		newNode := Node{
			Count: left.Count + right.Count,
			Left:  left,
			Right: right,
		}

		pq.Insert(newNode)
	}
	return pq.Peek()
}

// Helper function to print the tree
func (h *HuffNode) PrintTree(node *HuffNode, writer io.Writer) {
	if node == nil {
		return
	}

	/* if node.Left == nil && node.Right == nil { */
	fmt.Fprintf(writer, "%s : %d : %d \n", node.Data, node.Count, node.Weight)
	/* } */
	h.PrintTree(node.Left, writer)
	h.PrintTree(node.Right, writer)
}

func (h *HuffNode) AssignWeights(node *HuffNode) {
}

func (h *HuffNode) AssignWeights(node *HuffNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		node.Weight = 0
		h.AssignWeights(node.Left)
	}

	if node.Right != nil {
		node.Weight = 1
		h.AssignWeights(node.Right)
	}
}
