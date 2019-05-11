package main

import "fmt"

func findMax(a []rune) string {
	var max int
	for _, chr := range a {
		if int(chr) > max {
			max = int(chr)
		}
	}
	return string(max)
}

func sort_array(data [][]rune) {
	n := len(data)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = findMax(data[i])
	}
	fmt.Println(result)
}

func main() {
	data := [][]rune{[]rune{'d', 'e', 'f'}, []rune{'g', 'a'}}
	sort_array(data)
}
