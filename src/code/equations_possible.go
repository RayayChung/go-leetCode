package code

func equationsPossible(equations []string) bool {
	equalSet := make([]int, 26)
	unequalSet := make([]int, 26)

	for i := range equalSet {
		equalSet[i] = -1
		unequalSet[i] = -1
	}

	for _, equation := range equations {
		idxFirst := equation[0] - 'a'
		equalSet[int(idxFirst)] = int(idxFirst)

		idxSecond := equation[3] - 'a'
		equalSet[int(idxSecond)] = int(idxSecond)
	}

	for _, equation := range equations {
		ops := equation[1:3]
		if "==" == ops {
			first := findRoot(equalSet, int(equation[0]-'a'))
			second := findRoot(equalSet, int(equation[3]-'a'))
			if first != second {
				equalSet[second] = first
			}
		}
	}

	for _, equation := range equations {
		ops := equation[1:3]
		if "!=" == ops {
			first := findRoot(equalSet, int(equation[0]-'a'))
			second := findRoot(equalSet, int(equation[3]-'a'))
			if first == second {
				return false
			}
		}
	}

	return true
}

func findRoot(unionSet []int, idx int) int {
	for unionSet[idx] != idx {
		idx = unionSet[idx]
	}
	return idx
}
