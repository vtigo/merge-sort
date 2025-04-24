package main

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
	"sort"
)

// Quick sort implementation for comparison
func quickSort(arr []int) []int {
	// Make a copy to avoid modifying the original
	result := make([]int, len(arr))
	copy(result, arr)
	
	quickSortInPlace(result, 0, len(result)-1)
	return result
}

func quickSortInPlace(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSortInPlace(arr, low, pivot-1)
		quickSortInPlace(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Go's built-in sort implementation
func goSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	
	sort.Ints(result)
	return result
}

// Generate test data with different characteristics
func generateTestData(size int, scenario string) []int {
	result := make([]int, size)
	
	switch scenario {
	case "sorted":
		for i := range result {
			result[i] = i
		}
	case "reverse":
		for i := range result {
			result[i] = size - i
		}
	case "random":
		rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := range result {
			result[i] = rand.Intn(size * 10)
		}
	case "few_unique":
		rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := range result {
			result[i] = rand.Intn(10) // Only 10 unique values
		}
	}
	return result
}

// Benchmarks
func BenchmarkMergeSort(b *testing.B) {
	scenarios := []string{"sorted", "reverse", "random", "few_unique"}
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		for _, scenario := range scenarios {
			b.Run(fmt.Sprintf("Size=%d/Scenario=%s", size, scenario), func(b *testing.B) {
				data := generateTestData(size, scenario)
				b.ResetTimer()
				
				for b.Loop() {
					MergeSort(data)
				}
			})
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	scenarios := []string{"sorted", "reverse", "random", "few_unique"}
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		for _, scenario := range scenarios {
			b.Run(fmt.Sprintf("Size=%d/Scenario=%s", size, scenario), func(b *testing.B) {
				data := generateTestData(size, scenario)
				b.ResetTimer()
				
				for b.Loop() {
					quickSort(data)
				}
			})
		}
	}
}

func BenchmarkGoSort(b *testing.B) {
	scenarios := []string{"sorted", "reverse", "random", "few_unique"}
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		for _, scenario := range scenarios {
			b.Run(fmt.Sprintf("Size=%d/Scenario=%s", size, scenario), func(b *testing.B) {
				data := generateTestData(size, scenario)
				b.ResetTimer()
				
				for b.Loop() {
					goSort(data)
				}
			})
		}
	}
}
