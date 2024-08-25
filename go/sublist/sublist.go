package sublist

func Sublist(subset, superset []int) Relation {
	if len(subset) > len(superset) {
		return invert(Sublist(superset, subset))
	}

	if len(subset) == 0 && len(subset) == len(superset) {
		return RelationEqual
	}

	if len(subset) == 0 {
		return RelationSublist
	}

	if len(superset) == 0 {
		return RelationSuperlist
	}

	containsList := contains(subset, superset)
	isSublist := containsList && len(subset) != len(superset)
	isEqual := containsList && len(subset) == len(superset)
	if isEqual {
		return RelationEqual
	}

	if isSublist {
		return RelationSublist
	}

	return RelationUnequal
}

// contains checks if list contains chunk
func contains(chunk, list []int) bool {
	s := 0
	e := len(chunk) - 1

	for e < len(list) {
		chunk2 := list[s : e+1]
		if chunksAreIdentical(chunk, chunk2) {
			return true
		}
		s++
		e++
	}

	return false
}

// chunksAreIdentical checks if two chunks are identical considering the order and values
func chunksAreIdentical(chunk1, chunk2 []int) bool {
	for i := range chunk1 {
		if chunk1[i] != chunk2[i] {
			return false
		}
	}

	return true
}

// invert inverts the relation
func invert(r Relation) Relation {
	if r == RelationSublist {
		return RelationSuperlist
	}
	if r == RelationSuperlist {
		return RelationSublist
	}

	return r
}
