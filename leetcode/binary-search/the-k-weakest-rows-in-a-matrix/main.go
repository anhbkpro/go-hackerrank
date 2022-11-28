package the_k_weakest_rows_in_a_matrix

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type row struct {
	soldiers int
	index    int
}

//Runtime: 28 ms, faster than 19.79% of Go online submissions for The K Weakest Rows in a Matrix.
//Memory Usage: 5 MB, less than 41.18% of Go online submissions for The K Weakest Rows in a Matrix.
//Sort: O(nlog(n))
//Using Binary Search to count 1s: O(log(n))
func kWeakestRows(mat [][]int, k int) []int {
	var rows []row
	for i, r := range mat {
		rows = append(rows, row{
			soldiers: countOnes(r), // O(log(n))
			index:    i,
		})
	}
	//O(nlog(n))
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].soldiers > rows[j].soldiers {
			return false
		} else if rows[i].soldiers == rows[j].soldiers {
			if rows[i].index > rows[j].index {
				return false
			}
		}
		return true
	})
	fmt.Println(rows)
	ans := make([]int, 0)
	i := 0
	for k > 0 && i < len(rows) {
		ans = append(ans, rows[i].index)
		k--
		i++
	}
	return ans
}

func countOnes(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if arr[mid] == 1 {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi + 1
}

// An Item is something we manage in a priority queue.
type Item struct {
	soldiers int // The soldiers of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index         int // The index of the item in the heap.
	originalIndex int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and soldiers of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.soldiers = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

//Runtime: 14 ms, faster than 83.42% of Go online submissions for The K Weakest Rows in a Matrix.
//Memory Usage: 5 MB, less than 41.18% of Go online submissions for The K Weakest Rows in a Matrix.
func kWeakestRowsPriorityQueue(mat [][]int, k int) []int {
	// Some items and their priorities.
	var items []row
	for i, r := range mat {
		items = append(items, row{
			soldiers: countOnes(r), // O(log(n))
			index:    i,
		})
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	indexRange := int(math.Max(float64(len(mat)), float64(len(mat[0]))))
	for idx, item := range items {
		pq[i] = &Item{
			soldiers:      item.soldiers,
			priority:      item.soldiers*indexRange + idx,
			index:         item.index,
			originalIndex: idx,
		}
		i++
	}
	heap.Init(&pq)

	// Take the items out; they arrive in decreasing priority order.
	ans := make([]int, 0)
	for pq.Len() > 0 && k > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%.2d:%d ", item.soldiers, item.priority, item.index)
		ans = append(ans, item.originalIndex)
		k--
	}
	return ans
}
