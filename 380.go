package main

import "math/rand"

type RandomizedSet struct {
	idxByVal map[int]int
	valByIdx []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		idxByVal: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idxByVal[val]; ok {
		return false
	}
	this.idxByVal[val] = len(this.valByIdx)
	this.valByIdx = append(this.valByIdx, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.idxByVal[val]
	if !ok {
		return false
	}
	lastVal := this.valByIdx[len(this.valByIdx)-1]
	this.valByIdx[idx] = lastVal
	this.idxByVal[lastVal] = idx
	delete(this.idxByVal, val)
	this.valByIdx = this.valByIdx[:len(this.valByIdx)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.valByIdx[rand.Intn(len(this.valByIdx))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
