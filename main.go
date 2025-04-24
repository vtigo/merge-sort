package main

import (
	"fmt"
)

func main() {
	c := []int{9, 3, 5, 7, 2, 8, 1, 4, 6}
	
	sortedC := MergeSort(c)

	fmt.Printf("unsorted: %v", c)
	fmt.Printf("sorted: %v", sortedC)
}

func MergeSort(c []int) []int {
	if len(c) < 2 {
		return c
	}
	
	a := MergeSort(c[:len(c)/2])
	b := MergeSort(c[len(c)/2:])
	return merge (a, b)
}

func merge(a []int, b []int) []int {
	c := make([]int, len(a) + len(b))
	i := 0
	j := 0

	for k := range c {
		if i >= len(a) {
			c[k] = b[j]
			j++
		} else if j >= len(b) {
			c[k] = a[i]
			i++
		} else if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}

	return c
}


