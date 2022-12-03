package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	id           int32
	calorieCount int32
}

func main() {

	//create slice to store elves
	var elfList []Elf

	//call read file function to obtain inputs and save elves in slice
	elfList = readFileToGenerateList(elfList)

	//calculate elf with the most calories
	var highestCalorieElf = findMostCalorieElf(elfList)

	//print output
	fmt.Println("Highest Calorie Elf is: ", highestCalorieElf.id, " Total calories is: ", highestCalorieElf.calorieCount)
}

func readFileToGenerateList(elfList []Elf) []Elf {
	//error handling on opening file
	content, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(content)
	var elfId int32
	elfId = 1

	//read from file
	for scanner.Scan() {
		//create elf object
		elf := Elf{}
		//add elf ID
		elf.id = elfId

		for scanner.Text() != "" {
			calories, _ := strconv.Atoi(scanner.Text())
			elf.calorieCount = elf.calorieCount + int32(calories)
			scanner.Scan()
		}

		//if blank line add elf to slice and increment elfId
		elfList = append(elfList, elf)
		elfId++

	}
	//error handling on scanner
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return elfList
}

func findMostCalorieElf(list []Elf) Elf {
	var mostCalorieElf = list[0]

	for i, v := range list {

		if v.calorieCount > mostCalorieElf.calorieCount {
			mostCalorieElf = list[i]
		}
	}

	return mostCalorieElf
}
