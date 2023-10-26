package main

import "fmt"

func main() {
	dn := []int{1, 1, 3, 3, 4, 4, 5, 5, 5}
	fmt.Println(removeDuplicates(dn))
}

func removeDuplicates(dn []int) []int {

	dm := make(map[int]int)
	var result []int

	for _, n := range dn {
		if _, exists := dm[n]; !exists {
			result = append(result, n)
			dm[n] = n
		}
	}

	return result
}
