package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{2, 4, 5, 6, 7, 8, 1, 9, 3}

	sort.Ints(a)

	for _, v := range a {
		fmt.Println(v)
	}
}
