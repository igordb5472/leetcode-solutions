package main

type RecentCounter struct {
	requests []int
}

func Constructor933() RecentCounter { // 933 resolves name conflicts
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	firstNotExpired := 0
	for this.requests[firstNotExpired] < t-3000 {
		firstNotExpired++
	}
	this.requests = this.requests[firstNotExpired:]
	return len(this.requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
