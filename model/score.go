package model

type Score map[int]int

func NewScore() Score {
	return Score{}
}

func (s *Score) GetMaxScore() int {
	max := 0

	for _, v := range *s {
		if max < v {
			max = v
		}
	}

	return max
}

func (s *Score) GetMinScore() int {
	min := 1000

	for _, v := range *s {
		if min > v {
			min = v
		}
	}

	return min
}
