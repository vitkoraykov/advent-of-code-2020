package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open the file
	passwordList, err := os.Open("day2-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer passwordList.Close()
	// read from file
	scanner := bufio.NewScanner(passwordList)

	var acceptedPasswords []string

	for scanner.Scan() {
		// break the line down
		parts := strings.Fields(scanner.Text())

		// 1st part - parameters
		parameters := strings.Split(parts[0], "-")

		least, err := strconv.Atoi(parameters[0])
		if err != nil {
			panic(err)
		}
		most, err := strconv.Atoi(parameters[1])
		if err != nil {
			panic(err)
		}

		// 2nd part - letter
		letter := parts[1]
		actualLetter := string(letter[0])

		// 3rd part - password
		password := parts[2]

		// do some checking
		if string(password[least-1]) == actualLetter && string(password[most-1]) != actualLetter {
			acceptedPasswords = append(acceptedPasswords, password)
		}

		if string(password[least-1]) != actualLetter && string(password[most-1]) == actualLetter {
			acceptedPasswords = append(acceptedPasswords, password)
		}
	}
	fmt.Println(len(acceptedPasswords))
}
