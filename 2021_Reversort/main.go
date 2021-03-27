package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var numOfCases int
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	numOfCases, _ = strconv.Atoi(line)

	for i := 1; i <= numOfCases; i++ {

		fmt.Printf("Case #%d: ", i)
		scanner.Scan()
		_ = scanner.Text()

		scanner.Scan()
		line = scanner.Text()
		splits := strings.Split(line, " ")
		var in []int
		for _, s := range splits {
			v, _ := strconv.Atoi(s)
			in = append(in, v)
		}
		total := solve(&in)
		fmt.Println(total)
	}

}
func solve(in *[]int) int {
	var total int
	for i := 0; i < len(*in)-1; i++ {
		j := find(in, i, i+1)
		total += (j - i + 1)
		arrange(in, i, j)
	}
	return total
}

func arrange(s *[]int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func find(s *[]int, start, target int) int {
	for i := start; i < len(*s); i++ {
		if (*s)[i] == target {
			return i
		}
	}
	return 0
}
