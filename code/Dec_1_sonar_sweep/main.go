package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read the input from file
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	var k []byte
	k, err = ioutil.ReadFile(mydir + "/code/Dec_1_sonar_sweep/input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}

	var inputNumbers []int
	for _, i := range strings.Split(string(k), "\n") {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalf("unable to convert string to int %v", err)
		}
		inputNumbers = append(inputNumbers, num)
	}

	// Part 1
	var result1 int
	for i := 0; i < len(inputNumbers)-2; i++ {
		if inputNumbers[i] < inputNumbers[i+1] {
			result1 += 1
		}
	}
	fmt.Println("Result 1: ", result1)

	// Part 2
	var result2 int
	for i := 0; i < len(inputNumbers)-3; i++ {
		if inputNumbers[i]+inputNumbers[i+1]+inputNumbers[i+2] < inputNumbers[i+1]+inputNumbers[i+2]+inputNumbers[i+3] {
			result2 += 1
		}
	}
	fmt.Println("Result 2: ", result2)
}

// Cheers!