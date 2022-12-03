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

	//print output of highest calorie elf
	fmt.Println("Highest Calorie Elf is: ", highestCalorieElf.id, " Total calories is: ", highestCalorieElf.calorieCount)

	//print output for part 2 solution of top 3 elves and total calories
	var topElvesList = findTopElves(elfList, 3)
	calculateTotalTopElves(topElvesList)
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

func findElfIndex(list []Elf, elf Elf) int {

	var index = -1

	for i, v := range list {
		if elf == v {
			index = i
		}
	}
	return index
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

func findTopElves(list []Elf, count int) []Elf {
	//make a copy of original list
	var elfListCopy = list

	//create new list to hold top elves
	var topElves []Elf

	//iterate through as many times specified
	for i := 0; i < count; i++ {
		//find the highest calorie elf of list
		var highestCalorieElf = findMostCalorieElf(elfListCopy)

		//add elf to the topElves slice
		topElves = append(topElves, highestCalorieElf)

		//find elf index
		var index = findElfIndex(elfListCopy, highestCalorieElf)

		//remove elf from copy of elf list to iterate through again finding new highest value
		elfListCopy = append(elfListCopy[:index], elfListCopy[index+1:]...)
	}
	return topElves
}

func calculateTotalTopElves(list []Elf) {
	var totalSum int32
	fmt.Println("Highest Calorie Elves Are :")

	for i, v := range list {
		totalSum = totalSum + v.calorieCount
		fmt.Println("In spot: ", i+1, " is elf ID: ", v.id, "with a total of: ", v.calorieCount, "calories!")
	}

	fmt.Println("For a combined total of: ", totalSum)
}
