package models

type WonderfulOccurrences []int

func (el WonderfulOccurrences) Solution() int {
	hash := map[int]bool{}
	for _, i := range el {
		if _, ok := hash[i]; ok {
			delete(hash, i)
		} else {
			hash[i] = false
		}
	}
	for i, exists := range hash {
		if !exists {
			return i
		}
	}
	return -1
}
