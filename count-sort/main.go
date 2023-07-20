package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 7, 8, 0, 2, 1}
	fmt.Println(arr)
	sort(arr)
	fmt.Println(arr)
}
func sort(arr []int) {
	// find max
	var max int = -1
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	count := make([]int, max+1)
	for _, val := range arr {
		count[val]++
	}
	// fmt.Println(count)
	// find comulative sum
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	// fmt.Println(count)
	finalArr := make([]int, len(arr))
	// find position of arr elements in count
	for i := len(arr) - 1; i >= 0; i-- {
		finalArr[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	copy(arr, finalArr)
}
