package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, r.Intn(10))
	}

	for i, v := range numbers {

		if v%2 == 0 {
			fmt.Printf("index %v number %v is even\n", i, v)
		} else if v%2 != 0 {
			fmt.Printf("index %v number %v is odd\n", i, v)
		}

	}

}
