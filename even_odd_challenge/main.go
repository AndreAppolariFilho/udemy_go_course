package main

import "fmt"

func main() {
	var nums []int
	maxNum := 10

	for i := 0; i < maxNum; i += 1 {
		nums = append(nums, i)
	}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}
