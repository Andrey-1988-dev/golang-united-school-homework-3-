package homework

func reverse(input []int64) (result []int64) {
	var output []int64
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}
