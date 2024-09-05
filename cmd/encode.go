package cmd

import (
	"fmt"
	"io"
)

type HuffNode = Node

func NewFrequencyMap(input string) map[string]int {
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
	/* fmt.Fprintf(writer, "%s : %d : %d \n", node.Data, node.Count, node.Weight) */
	/* } */
	h.PrintTree(node.Left, writer)
	fmt.Fprintf(writer, "%s : %d : %s \n", node.Data, node.Count, node.Weight)
	h.PrintTree(node.Right, writer)
}

func (h *HuffNode) VariableEncoding(node *HuffNode) {
}

/*
	 func (h *HuffNode) getBinEncoding(node *HuffNode, frequencyMap map[string]int) {
	}
*/
func (h *HuffNode) EncodeStringMap(node *HuffNode, data map[string]int) map[string]string { //could be changed to a byte array
	var encodeMap = make(map[string]string)
	for key := range data {
		var output string = ""
		h.TraverseTree(node, key, &output, &encodeMap)

	}
	return encodeMap

}

func (h *HuffNode) EncodeString(data string, encodeKey map[string]string) string { //could be changed to a byte array
	var result string

	for _, v := range data {
		result += encodeKey[string(v)]
	}

	return result

}

func (h *HuffNode) DecodeString(node *HuffNode, encodedData string) string { //could be changed to a byte array
	var result string
	var pointer int

	for pointer < len(encodedData) {
		/* fmt.Println("this ran", pointer) */
		result += h.decodeTraversal(node, &encodedData, &pointer)
	}

	return result
}

func (h *HuffNode) decodeTraversal(node *HuffNode, encodedData *string, pointer *int) string {
	if node == nil {
		return ""
	}

	if *pointer > len(*encodedData) {
		return node.Data
	}

	if node.Left == nil && node.Right == nil {
		return node.Data
	}

	if string((*encodedData)[*pointer]) == "0" {
		*pointer++
		return h.decodeTraversal(node.Left, encodedData, pointer)
	}
	/* 	else {

	   		return h.decodeTraversal(node.Right, encodedData, &*pointer++)
	   	}
	*/
	/* return node.Data */
	*pointer++
	return h.decodeTraversal(node.Right, encodedData, pointer)
}

/* func (h *HuffNode) TraverseTree(node *HuffNode, variableCharacter string, output string) string { */
/* 	if node == nil {
   		return ""
   	}

   	// If the node is a leaf node and matches the variable character
   	if node.Left == nil && node.Right == nil {
   		if node.Data == variableCharacter {
   			fmt.Println(node.Data, "res:", output)
   			return output
   		}
   		return ""
   	}

   	// Traverse the left subtree
   	leftOutput := h.TraverseTree(node.Left, variableCharacter, output+"0")
   	if leftOutput != "" {
   		return leftOutput
   	}

   	// Traverse the right subtree
   	rightOutput := h.TraverseTree(node.Right, variableCharacter, output+"1")
   	if rightOutput != "" {
   		return rightOutput
   	}

   	// If neither left nor right subtree produced the result, return an empty string
   	return ""
} */
func (h *HuffNode) TraverseTree(node *HuffNode, variableCharacter string, output *string, encoding *map[string]string) {

	if node == nil {
		return
	}

	if node.Data == variableCharacter && node.Left == nil && node.Right == nil {
		(*encoding)[node.Data] = *output
		/* fmt.Println(node.Data, "res:", *output) */
		return
	}

	*output += "0"
	h.TraverseTree(node.Left, variableCharacter, output, encoding)
	*output = (*output)[:len(*output)-1]

	*output += "1"
	h.TraverseTree(node.Right, variableCharacter, output, encoding)
	*output = (*output)[:len(*output)-1]

}

func (h *HuffNode) AssignWeights(node *HuffNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		node.Weight = "0"
		/* node.Weight = 0 */
		h.AssignWeights(node.Left)
	}

	if node.Right != nil {
		node.Weight = "1"
		/* node.Weight = 1 */
		h.AssignWeights(node.Right)
	}
}
