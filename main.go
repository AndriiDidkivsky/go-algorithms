package main

import (
	"./bsearch"
	"fmt"
)

func main() {
	array := []int{1, 4, 8, 9, 11, 15, 17, 22, 23, 26}
	fmt.Println(bsearch.BSearch(array, 7))
}
