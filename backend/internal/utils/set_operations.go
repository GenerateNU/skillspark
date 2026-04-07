package utils

func ListComparison(listOne []string, listTwo []string) int {

	counterOne := make(map[string]bool)
	counterTwo := make(map[string]bool)
	commonElements := 0

	for idx := range listOne {
		element := listOne[idx]
		counterOne[element] = true
	}

	for idx := range listTwo {
		element := listTwo[idx]
		counterTwo[element] = true
	}

	for element := range counterOne {
		if _, ok := counterTwo[element]; ok {
			commonElements += 1
		}
	}

	return commonElements
}
