package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input := "2"
	input := "0023"
	// input := "00000"
	res := addBracket(input)
	fmt.Println(res)
}

func addBracket(in string) string {
	var nums []int
	var result string
	var current int
	splits := strings.Split(in, "")
	for _, s := range splits {
		v, _ := strconv.Atoi(s)
		nums = append(nums, v)
	}

	for i, n := range nums {
		if n == 0 {
			result += fmt.Sprintf("%d", n)
		}
		// Insert open brackets
		for j := 1; j <= n; j++ {
			if n > current {
				fmt.Printf("n: %d, current: %d \n", n, current)
				result += "("
			}
			if j == n {
				result += fmt.Sprintf("%d", n)
				current++
			}
		}

		// Insert closing brackets
		if i < (len(nums) - 1) {
			diff := nums[i] - nums[i+1]
			if diff > 0 && current >= diff {
				for j := 1; j <= diff; j++ {
					current = n
					result += ")"
				}
			}
		}

		// final closing
		if i == len(nums)-1 {
			fmt.Println("here")
			for j := 0; j <= current; j++ {
				result += ")"
			}
		}
	}

	return result
}
