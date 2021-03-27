package main

import (
	"bufio"
	"fmt"
	"log"
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
		var err error
		var nrows int
		var t int  // trace
		var dr int // duplicate rows
		var dc int // duplicate columns

		scanner.Scan()
		line = scanner.Text()
		nrows, _ = strconv.Atoi(line)
		mc := make(map[int]map[int]bool) // map for column
		mcc := make(map[int]bool)        // map for column
		for j := 0; j < nrows; j++ {
			mr := make(map[int]bool) // map for row

			scanner.Scan()
			line := scanner.Text()
			splits := strings.Split(line, " ")
			nums := make([]int, len(splits))

			for i, s := range splits {
				if nums[i], err = strconv.Atoi(s); err != nil {
					log.Fatal(err)
				}
			}
			var bR bool // bool to skip row

			for k := 0; k < len(nums); k++ {
				n := nums[k]

				// row
				if bR == false {
					if _, ok := mr[n]; ok {
						dr++
						bR = true
					} else {
						mr[n] = true
					}
				}

				// column
				if v, ok := mc[k]; ok {
					if _, ook := v[n]; ook {
						mcc[k] = true
					} else {
						mc[k][n] = true
					}
				} else {
					mc[k] = make(map[int]bool)
					mc[k][n] = true
				}

			}
			t += nums[j]
		}
		for _, c := range mcc {
			if c == true {
				dc++
			}
		}
		// Case #1: 4 0 0
		fmt.Printf("Case #%d: %d %d %d\n", i, t, dr, dc)

	}

}
