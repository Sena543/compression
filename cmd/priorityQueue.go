package cmd

type Node struct {
	Data        string
	Count       int
	Left, Right *Node
}

type PriorityQueue struct {
	queueArray []Node
}

func (p *PriorityQueue) Insert(value Node) {
	p.queueArray = append(p.queueArray, value)
	lastIdx := len(p.queueArray) - 1
	p.heapifyUp(lastIdx)

}

func (p *PriorityQueue) Pop() *Node {
	if len(p.queueArray) == 0 {
		return nil
	}
	lastIdx := len(p.queueArray) - 1
	p.swap(0, lastIdx)
	value := p.queueArray[lastIdx]
	p.queueArray = p.queueArray[:lastIdx]
	p.heapifyDown(0)
	return &value
}

func (p *PriorityQueue) PQ() []Node {
	return p.queueArray
}

func (p *PriorityQueue) Peek() *Node {
	return &p.queueArray[0]
}

func (p *PriorityQueue) leftNode(index int) int {
	return 2*index + 1
}

func (p *PriorityQueue) rightNode(index int) int {
	return 2*index + 2
}

func (p *PriorityQueue) parentNode(index int) int {
	return (index - 1) / 2
}

func (p *PriorityQueue) swap(i, j int) {
	temp := p.queueArray[i]
	p.queueArray[i] = p.queueArray[j]
	p.queueArray[j] = temp
}

func (p *PriorityQueue) heapifyUp(index int) {
	if index == 0 {
		return
	}

	value := p.queueArray[index]
	parentIndex := p.parentNode(index)
	parentValue := p.queueArray[parentIndex]

	if parentValue.Count > value.Count {
		p.swap(parentIndex, index)
		p.heapifyUp(parentIndex)
	}

}

func (p *PriorityQueue) heapifyDown(index int) {
	leftIndex := p.leftNode(index)
	rightIndex := p.rightNode(index)

	if index >= len(p.queueArray)-1 || leftIndex >= len(p.queueArray)-1 {
		return
	}
	currentValue := p.queueArray[index]
	lValue := p.queueArray[leftIndex]
	rValue := p.queueArray[rightIndex]

	if lValue.Count < rValue.Count && lValue.Count < currentValue.Count {
		p.swap(leftIndex, index)
		p.heapifyDown(leftIndex)
	} else if rValue.Count < lValue.Count && rValue.Count < currentValue.Count {
		p.swap(rightIndex, index)
		p.heapifyDown(rightIndex)
	}

}
