package models

type SearchingMissing []int

func (el SearchingMissing) Solution() int {
	hash := map[int]bool{}
	for _, i := range el {
		hash[i] = true
	}
	lenA := len(el)
	for i := 1; i <= lenA; i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}
	return -1
}
