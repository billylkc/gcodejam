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

	for c := 1; c <= numOfCases; c++ {
		var (
			cx, cy int    // cost x, cost y
			line   string // original
			text   string
		)
		scanner.Scan()
		line = scanner.Text()
		splits := strings.Split(line, " ")
		cx, _ = strconv.Atoi(splits[0])
		cy, _ = strconv.Atoi(splits[1])
		text = splits[2]

		res := final(text, cx, cy)
		fmt.Printf("Case #%d: %d\n", c, res)
	}

}

func final(text string, cx, cy int) int {
	text, ca, cb := convert(text, cx, cy)
	return (solve(text, ca, cb))
}

func convert(text string, cx, cy int) (string, int, int) {
	var ca, cb int
	if cx <= cy {
		ca = cx
		cb = cy
		text = strings.Replace(text, "C", "A", -1)
		text = strings.Replace(text, "J", "B", -1)
	} else {
		ca = cy
		cb = cx
		text = strings.Replace(text, "J", "A", -1)
		text = strings.Replace(text, "C", "B", -1)
	}
	return text, ca, cb
}

func solve(s string, ca, cb int) int {
	var (
		idx     int
		cost    int
		seen    string
		last    string
		current string
	)

	idx = len(s) - 1
	last = string(s[idx])
	seen = string(s[idx])
	for idx > 0 {
		idx--
		current = string(s[idx])

		if seen == "?" {
			last = current
			seen = current
		} else {
			if current == "?" {
			Loop:
				for idx > 0 {
					idx--
					current = string(s[idx])

					if current == "?" || current == seen {
						last = current
					} else if current != seen {
						if current == "A" { // Low cost
							cost += ca
							seen = current
						} else { // High cost
							cost += cb
							seen = current
						}
						last = current
						break Loop
					} else {

					}
					if idx == 0 {
						return cost
					}
				}

			} else {
				if current != last {
					if current == "A" { // Low cost
						cost += ca
						seen = current
					} else { // High cost
						cost += cb
						seen = current
					}
					last = current
				}
			}
		}
	}
	return cost
}
