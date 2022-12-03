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
	var roundList []Round
	//retrieve input and parse into Round structures list with appropriate player selections
	roundList = readInputValues()
	//calculate points based on your player selection
	roundList = calculateChoicePoints(roundList)
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
		//get line of file and save into string
		input := scanner.Text()
		//split values of string by space and hold values in array
		inputArray := strings.Split(input, " ")
		//obtain individual values from array
		v1 := inputArray[0]
		v2 := inputArray[1]
		round := Round{v1, v2, "", 0}
		list = append(list, round)
	}

	//error handling on scanner
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return list
}

func calculateChoicePoints(list []Round) []Round {

	for i, v := range list {
		switch v.myChoice {
		// if value is rock add 1 point
		case "X":
			list[i].points = 1
		//if value is paper add 2 points
		case "Y":
			list[i].points = 2
		//if value is scissors add 3 points
		case "Z":
			list[i].points = 3
		}
	}

	return list
}
