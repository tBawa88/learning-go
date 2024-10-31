package prioqueue

import (
	"fmt"
	"math"
)

type Node struct {
	val      interface{}
	priority int
}

type Queue struct {
	list []Node
}

// Prints the queue
func (q *Queue) Print() {
	fmt.Println("===== printing =====")
	for _, node := range q.list {
		fmt.Println(node.val, node.priority)
	}
}

// Enqueue(val, priority int) : takes in a value and it' priority
// that value is placed at a position according to it's priority
func (q *Queue) Enqueue(val interface{}, priority int) {
	q.list = append(q.list, Node{val, priority})
	q.bubbleUp()
}

// Returns the value with minimum Priority-value in the queue
// if the queue is empty, returns nil
func (q *Queue) Dequeue() interface{} {
	if len(q.list) == 0 {
		return nil
	}
	size := len(q.list)
	minNode := q.list[0]
	//update the first value to the last
	q.list[0] = q.list[size-1]
	// remove the last value from the list
	q.list = q.list[:size-1]
	//perform the sinkdown operation
	q.sinkDown()
	return minNode.val
}

func (q *Queue) bubbleUp() {
	currIdx := len(q.list) - 1
	currPrio := q.list[currIdx].priority
	parentIdx := (currIdx - 1) / 2

	for parentIdx >= 0 && currPrio < q.list[parentIdx].priority {
		// parentPrio := q.list[parentIdx].priority
		q.swap(currIdx, parentIdx)
		currIdx = parentIdx
		parentIdx = (currIdx - 1) / 2
	}
}

func (q *Queue) sinkDown() {
	if len(q.list) < 2 {
		return
	}
	currIdx := 0
	currPrio := q.list[0].priority
	leftchild := currIdx*2 + 1
	rightChild := int(math.Min(float64(len(q.list)-1), float64(leftchild+1)))

	for leftchild < len(q.list) {
		var smallerIdx int
		if q.list[leftchild].priority < q.list[rightChild].priority {
			smallerIdx = leftchild
		} else {
			smallerIdx = rightChild
		}
		// if current node's prio is smaller than both child nodes, then break out
		if currPrio < q.list[smallerIdx].priority {
			return
		}
		q.swap(currIdx, smallerIdx)
		currIdx = smallerIdx
		leftchild = currIdx*2 + 1
		rightChild = leftchild + 1
	}

}

func (q *Queue) swap(a, b int) {
	temp := q.list[a]
	q.list[a] = q.list[b]
	q.list[b] = temp
}
