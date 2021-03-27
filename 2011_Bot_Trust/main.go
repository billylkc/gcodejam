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
	fmt.Scanf("%d", &numOfCases)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i <= numOfCases; i++ {
		fmt.Printf("\n-----------------------------------\n")
		fmt.Printf("Case #%d:", i)

		scanner.Scan()
		line := scanner.Text()
		fmt.Println(line)

		splits := strings.Split(line, " ")
		length := splits[0]
		n, _ := strconv.Atoi(length)

		var nextRobot []string // e.g. BBBOOOB
		var nextRoom []int     // 2 -> 2

		// Construct list
		for j := 1; j <= n; j++ {
			robot := splits[2*j-1]
			room, _ := strconv.Atoi(splits[2*j])
			fmt.Printf("Robot: %s, Room: %d\n", robot, room)

			nextRobot = append(nextRobot, robot)
			nextRoom = append(nextRoom, room)
		}
		fmt.Println(nextRobot)
		fmt.Println(nextRoom)

		// step := 1
		// var currentBot, anotherRobot string
		// currentB := 1
		// currentO := 1

		// for r := 0; r < len(nextRoom); r++ {
		// 	currentRobot = nextRobot[r]
		// 	target = nextRoom[r]
		// 	if currentBot == "B" {
		// 		anotherRobot == "O"
		// 		current = currentB
		// 	} else {
		// 		anotherRobot == "B"
		// 		current = currentO
		// 	}

		// 	for {
		// 		if current != target {
		// 			current += 1
		// 			step++
		// 			break
		// 		} else {

		// 		}
		// 	}

		// }
	}

}

func findNextRoom(r string, cur int, robot []string, room []int) int {
	if cur == len(robot) {
		return 0
	}
	for i, v := range robot[cur+1:] {
		if v == r {
			return cur + i + 1
		}
	}
	return 0
}
