package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}

	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, v := range s {
		s[i] = fn(v)
	}
	return s
}

func (s IntList) Reverse() IntList {
	result := IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return result
}

func (s IntList) Append(lst IntList) IntList {
	s = append(s, lst...)
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, lst := range lists {
		s = append(s, lst...)
	}
	return s
}
