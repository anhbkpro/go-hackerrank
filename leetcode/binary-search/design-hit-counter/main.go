package design_hit_counter

type queue []int

func (s *queue) push(item int) {
	*s = append(*s, item)
}

func (s *queue) pop() int {
	returnVal := (*s)[0]
	*s = (*s)[1:]
	return returnVal
}

func (s *queue) front() int {
	return (*s)[0]
}

// HitCounter Runtime: 0 ms, faster than 100.00% of Go online submissions for Design Hit Counter.
//Memory Usage: 2.3 MB, less than 67.21% of Go online submissions for Design Hit Counter.
type HitCounter struct {
	hits queue
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (hc *HitCounter) Hit(timestamp int) {
	hc.hits.push(timestamp)
}

func (hc *HitCounter) GetHits(timestamp int) int {
	for len(hc.hits) > 0 {
		diff := timestamp - hc.hits.front()
		if diff >= 300 {
			hc.hits.pop()
		} else {
			break
		}
	}
	return len(hc.hits)
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
