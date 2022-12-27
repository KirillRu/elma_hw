package models

type CheckingSequence []int

func (el CheckingSequence) Solution() int {
	lenA := len(el)
	hash := map[int]bool{}
	for _, i := range el {
		if i > lenA || i < 1 {
			return 0
		}
		if _, ok := hash[i]; ok {
			return 0
		}

		hash[i] = false
	}
	return 1
}
