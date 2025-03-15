/*
Binary Search
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const LEN int = 100
	var arr [LEN]int

	// rand with different soruce every runtime
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	searchNum := rand.Intn(100)
	for i := 0; i < LEN; i++ {
		arr[i] = rand.Intn(100)
	}

	// sort - selection sort
	for i := 0; i < LEN-1; i++ {
		min := i
		for j := i + 1; j < LEN; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	// printing
	for i := 0; i < LEN; i++ {
		fmt.Printf("[i=%d]%d  ", i, arr[i])
	}

	// binary search
	var left, right, found = 0, LEN - 1, -1
	for left <= right {
		mid := left + (right-left)/2
		if searchNum < arr[mid] {
			right = mid - 1
		} else if searchNum > arr[mid] {
			left = mid + 1
		} else {
			found = mid
			// make sure you found the first item if duplicated items exist
			for mid--; mid >= 0 && searchNum == arr[mid]; mid-- {
				found = mid
			}
			break
		}
	}
	fmt.Printf("\nSearch Num:%d, found at Index %d\n", searchNum, found)
}
