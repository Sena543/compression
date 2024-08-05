package cmd

import "fmt"

type HuffmanNode struct {
	NodeData    Node
	Left, Right *HuffmanNode
}

func NewEncode(input string) map[string]int {
	var result = make(map[string]int)
	for _, val := range input {
		result[string(val)] += 1
	}
	return result
}

var rootNode *HuffmanNode

func (h *HuffmanNode) HuffTree(pq PriorityQueue) {
	/* fmt.Println(pq.PQ()) */
	pqLen := len(pq.queueArray)
	for i := 1; i < pqLen; i++ {
		fmt.Println(i)
		nodeValue := pq.Pop()
		h.Insert(nodeValue)
	}
	h.Traverse(rootNode)
}

func (h *HuffmanNode) Insert(newNode *Node) {

	if rootNode == nil {
		rootNode = &HuffmanNode{NodeData: *newNode}
		return
	}
	currNode := rootNode
	newHuffNode := &HuffmanNode{NodeData: *newNode}

	for currNode != nil {
		if currNode.NodeData.Count > newHuffNode.NodeData.Count { //go left
			if currNode.Left == nil {
				currNode.Left = newHuffNode
				/* return */
				break
			} else {
				fmt.Println(5)
				currNode = currNode.Left
			}

		} else { //go right
			if currNode.Right == nil {
				currNode.Right = newHuffNode
				/* return */
				break
			} else {
				currNode = currNode.Right
			}
		}
	}

}

func (h *HuffmanNode) Traverse(node *HuffmanNode) {
	if node == nil {
		return
	}
	h.Traverse(node.Left)
	fmt.Println(node.NodeData)
	h.Traverse(node.Right)
}
