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
	outcome        string
	points         int32
}

func main() {
	fmt.Println("Which way do you want to play? Please enter 1 or 2: ")
	var input int
	fmt.Scanln(&input)

	//create list of rounds
	var roundList []Round
	//retrieve input and parse into Round structures list with appropriate player selections
	roundList = readInputValues(input)

	if input == 2 {
		roundList = calculateChoicePoints(roundList, 2)
		roundList = assignMyChoice(roundList) // if playing part 2 way assign my choice option first by looking at the outcome before calculating choice points
	}

	//calculate points based on your player selection
	roundList = calculateChoicePoints(roundList, 1)
	//calculate outcome of game
	roundList = roundOutcomeCalculator(roundList)

	//calculate total points
	var pointTotal = calculatePoints(roundList)
	//print output
	fmt.Println("Total points: ", pointTotal)

}

func readInputValues(style int) []Round {
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
		round := Round{}
		//get line of file and save into string
		input := scanner.Text()
		//split values of string by space and hold values in array
		inputArray := strings.Split(input, " ")
		//obtain opponents choice and other property either out choice or outcome
		v1 := inputArray[0]
		v2 := inputArray[1]
		//save values for reusability sake
		round = Round{v1, v2, "", 0}
		//append list with new object
		list = append(list, round)
	}

	//error handling on scanner
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return list
}

func calculateChoicePoints(list []Round, input int) []Round {

	for i, v := range list {
		switch v.myChoice {
		// if value is rock add 1 point
		case "X":

			if input == 1 {
				list[i].points = 1
			}
			if input == 2 {
				list[i].outcome = "lose" // if value is 'X' and style 2 game the outcome is lose
			}

		//if value is paper add 2 points
		case "Y":

			if input == 1 {
				list[i].points = 2
			}
			if input == 2 {
				list[i].outcome = "tie" // if value is 'Y' and style 2 game the outcome is tie
			}

		//if value is scissors add 3 points
		case "Z":

			if input == 1 {
				list[i].points = 3
			}

			if input == 2 {
				list[i].outcome = "win" // if value is 'Z' and style 2 game the outcome is win
			}
		}
	}

	return list
}

func assignMyChoice(list []Round) []Round {

	for i, v := range list {

		switch list[i].opponentChoice {

		case "A":
			if v.outcome == "win" {
				list[i].myChoice = "Y"
			}
			if v.outcome == "lose" {
				list[i].myChoice = "Z"
			}
			if v.outcome == "tie" {
				list[i].myChoice = "X"
			}
		case "B":
			if v.outcome == "win" {
				list[i].myChoice = "Z"
			}
			if v.outcome == "lose" {
				list[i].myChoice = "X"
			}
			if v.outcome == "tie" {
				list[i].myChoice = "Y"
			}
		case "C":
			if v.outcome == "win" {
				list[i].myChoice = "X"
			}
			if v.outcome == "lose" {
				list[i].myChoice = "Y"
			}
			if v.outcome == "tie" {
				list[i].myChoice = "Z"
			}
		}
	}
	return list
}

func roundOutcomeCalculator(list []Round) []Round {

	for i, v := range list {

		switch list[i].myChoice {

		// I choose rock
		case "X":

			if v.opponentChoice == "B" {
				list[i].outcome = "lose" //opponent chooses paper
			}
			if v.opponentChoice == "C" {
				list[i].outcome = "win" //opponent chooses scissors
			}
			if v.opponentChoice == "A" {
				list[i].outcome = "tie" //opponent chooses rock
			}

		//I choose paper
		case "Y":
			if v.opponentChoice == "A" {
				list[i].outcome = "win" //opponent chooses rock
			}
			if v.opponentChoice == "C" {
				list[i].outcome = "lose" //opponent chooses scissors
			}
			if v.opponentChoice == "B" {
				list[i].outcome = "tie" //opponent chooses paper
			}
		//I choose scissors
		case "Z":
			if v.opponentChoice == "A" {
				list[i].outcome = "lose" //opponent chooses rock
			}
			if v.opponentChoice == "B" {
				list[i].outcome = "win" //opponent chooses paper
			}
			if v.opponentChoice == "C" {
				list[i].outcome = "tie" //opponent chooses scissors
			}
		}
	}
	return list
}

func calculatePoints(list []Round) int32 {
	var pointSum int32

	for i := range list {
		switch list[i].outcome {
		case "tie":
			list[i].points = list[i].points + 3
		case "win":
			list[i].points = list[i].points + 6
		case "lose":
			list[i].points = list[i].points + 0
		}

		pointSum = pointSum + list[i].points
	}
	return pointSum
}
