package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Round struct {
	opponentChoice string
	myChoice       string
	outCome        string
	points         int32
}

func main() {
	//create list of rounds
	//var roundList []Round
	//retrieve input and parse into Round structures list with appropriate player selections
	readInputValues()
	//calculate points based on your player selection

	//calculate outcome of game and adjust points accordingly
}

func readInputValues() []Round {
	var list []Round

	//error handling on opening file
	content, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	//creating scanner object to read file
	scanner := bufio.NewScanner(content)

	//scanning file for next line
	for scanner.Scan() {
		input := scanner.Text()
		v1, v2 := strings.Split(input, " ")
		fmt.Println("Opponent value: ", v1, " My Choice: ", v2)
		round := Round{v1, v2, "", 0}
		list = append(list, round)
	}

	//error handling on scanner
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return list
}
