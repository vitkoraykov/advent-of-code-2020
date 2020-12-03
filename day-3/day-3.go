package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open file
	areaMap, err := os.Open("day3-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer areaMap.Close()

	// read from file
	scanner := bufio.NewScanner(areaMap)

	currentPos := 0
	stepRight := 3
	trees := 0

	for scanner.Scan() {
		// check for a tree at starting pos
		if string(scanner.Text()[currentPos]) == "#" {
			trees++
		}
		// move 3 and finish row
		currentPos = currentPos + stepRight
		// if its the end of the row
		if currentPos >= len(scanner.Text()) {
			currentPos = currentPos - len(scanner.Text())
		}
	}
	fmt.Println(trees)
}
