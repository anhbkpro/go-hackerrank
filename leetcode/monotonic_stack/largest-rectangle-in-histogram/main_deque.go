package largest_rectangle_in_histogram

import "fmt"

func LargestRectangleArea(heights []int) int {
	n := len(heights)
	maxArea := 0

	// two increasing MQs, one for left nearest values and one for right ones
	leftRight := NewMq(heights, -1, n)
	for i := 0; i < n; i++ {
		leftRight.push(Item{heights[i], i}) // deque nearest leftRight: [-1 -1 1 2 1 4]
	}
	fmt.Println(leftRight.nearest)

	rightLeft := NewMq(heights, n, n)
	for i := n - 1; i >= 0; i-- {
		rightLeft.push(Item{heights[i], i}) // deque nearest rightLeft: [1 -1 4 4 -1 -1]
	}
	fmt.Println(rightLeft.nearest)

	for i := 0; i < n; i++ {
		width := rightLeft.nearest[i] - leftRight.nearest[i] - 1
		currentArea := heights[i] * width
		if currentArea > maxArea {
			maxArea = currentArea
		}
	}

	return maxArea
}

type Item struct {
	value int
	index int
}

type Mq struct {
	heights             []int
	q                   Deque
	defaultNearestValue int
	nearest             []int
}

func NewMq(heights []int, defaultNearestValue, n int) Mq {
	return Mq{
		heights:             heights,
		defaultNearestValue: defaultNearestValue,
		q:                   NewDeque(),
		nearest:             make([]int, n),
	}
}

func (m *Mq) push(newItem Item) {
	for !m.q.empty() && newItem.value <= m.heights[m.q.getLast()] {
		m.q.popLast()
	}
	if m.q.empty() {
		m.nearest[newItem.index] = m.defaultNearestValue
	} else {
		m.nearest[newItem.index] = m.q.getLast()
	}

	m.q.push(newItem.index)
}

type Deque struct {
	indexes []int
}

func NewDeque() Deque {
	return Deque{
		indexes: make([]int, 0),
	}
}

func (d *Deque) push(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *Deque) getFirst() int {
	return d.indexes[0]
}

func (d *Deque) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *Deque) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *Deque) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *Deque) empty() bool {
	return 0 == len(d.indexes)
}
