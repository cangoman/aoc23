package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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
	re := regexp.MustCompile(`\d`)

	var result int

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		result += calibrationValue(matches)
	}

	fmt.Printf("result: %d\n", result)
}


func calibrationValue(matches []string) int {
	resultString := matches[0] + matches[len(matches) - 1]

	result, err := strconv.Atoi(resultString)
	check(err)
	return result
}