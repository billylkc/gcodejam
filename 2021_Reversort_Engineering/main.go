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
		line = scanner.Text()
		splits := strings.Split(line, " ")
		var in []int
		for _, s := range splits {
			v, _ := strconv.Atoi(s)
			in = append(in, v)
		}
		num := in[0]
		target := in[1]

		res, err := final(num, target)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		output := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), " "), "[]")
		fmt.Println(output)
	}
}

func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + 1 + start
	}
	return xs
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
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

func final(num, target int) ([]int, error) {
	ll := permutation(rangeSlice(0, num))
	result := make([]int, num)
	for i := 0; i < len(ll); i++ {
		s := ll[i]
		copy(result, s)

		steps := solve(&s)
		if steps == target {
			return result, nil
		}
	}
	return result, fmt.Errorf("IMPOSSIBLE")
}
