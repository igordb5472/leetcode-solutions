package main

type LRUCache struct {
	history []int
	data    map[int]int
}

func Constructor146(capacity int) LRUCache { // without 146; 146 here for solve name conflict
	return LRUCache{
		history: make([]int, 0, capacity),
		data:    make(map[int]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.data[key]
	if !ok {
		return -1
	}
	this.liftUp(key)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; !ok {
		if len(this.history) == cap(this.history) {
			delete(this.data, this.history[len(this.history)-1])
			this.history[len(this.history)-1] = key
		} else {
			this.history = append(this.history, key)
		}
	}
	this.liftUp(key)
	this.data[key] = value
}

func (this *LRUCache) liftUp(key int) {
	for i := len(this.history) - 1; i > 0; i-- {
		if this.history[i] == key {
			this.history[i-1] ^= this.history[i]
			this.history[i] ^= this.history[i-1]
			this.history[i-1] ^= this.history[i]
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
