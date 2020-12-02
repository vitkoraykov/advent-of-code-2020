package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// open file
	expenseList, err := os.Open("day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer expenseList.Close()
	// create an array of int
	scanner := bufio.NewScanner(expenseList)
	var values []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		values = append(values, num)
	}
	// do some math
	var result int
	for a := 0; a < len(values); a++ {
		for b := a; b < len(values); b++ {
			if values[a]+values[b] == 2020 {
				result = values[a] * values[b]
				fmt.Printf("Found a match: %d + %d = %d", values[a], values[b], result)
			}
			// part 2
			for c := b; c < len(values); c++ {
				if values[a]+values[b]+values[c] == 2020 {
					result = values[a] * values[b] * values[c]
					fmt.Printf("Found tripple match: %d + %d + %d = %d", values[a], values[b], values[c], result)
				}
			}
		}
	}
}
