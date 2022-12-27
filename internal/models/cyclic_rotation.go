package models

type CyclicRotation struct {
	ArrayIn        []int
	RotationCounts int
}

func (el CyclicRotation) Solution() []int {
	aSlice := el.ArrayIn[:]
	lenA := len(el.ArrayIn)
	result := make([]int, 0, lenA)
	if el.RotationCounts > 0 && lenA > 0 {
		realK := el.RotationCounts % lenA
		if realK > 0 {
			result = aSlice[lenA-realK:]
			result = append(result, aSlice[:lenA-realK]...)
		}
	}
	if len(result) == 0 {
		result = aSlice
	}
	return result
}
