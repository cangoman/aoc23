package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var firstString map[string]string = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
	"eightwo": "8",
	"nineight": "9",
	"threeight": "3",
	"twone": "2",
	"eighthree": "8",
	"oneight": "1",
	"fiveight": "5",
	"sevenine": "7",
}

var lastString map[string]string = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
	"eightwo": "2",
	"eighthree": "3",
	"nineight": "8",
	"threeight": "8",
	"twone": "1",
	"oneight": "8",
	"fiveight": "8",
	"sevenine": "9",
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	
	f, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	re := regexp.MustCompile(`\d|one(ight)?|two(ne)?|three(ight)?|four|five(ight)?|six|seven(ine)?|eight(hree)?(wo)?|nine(ight)?`)

	var result int = 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		result += calibrationValue(matches)
	}

	fmt.Printf("result: %d\n", result)
}


func calibrationValue(matches []string) int {
	msd := firstString[matches[0]]
	if (msd == "") {
		msd = matches[0]
	}

	lsd := lastString[matches[len(matches) - 1]]
	if (lsd == "") {
		lsd = matches[len(matches) - 1]
	}

	resultString := msd + lsd
	result, err := strconv.Atoi(resultString)
	check(err)
	return result
}

