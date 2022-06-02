package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var output []string
	var keys []int

	for index := range input {
		keys = append(keys, index)
	}

	sort.Ints(keys)

	for _, value := range keys {
		output = append(output, input[value])
	}

	return output
}
