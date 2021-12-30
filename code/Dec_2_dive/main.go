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
	k, err = ioutil.ReadFile(mydir + "/code/Dec_2_dive/input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}

	// Part 1
	var horizontalPosition, depth int
	for _, i := range strings.Split(string(k), "\n") {
		cmd := strings.Split(i, " ")
		num, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalf("unable to convert string to int %v", err)
		}

		if cmd[0] == "forward" {
			horizontalPosition += num
		} else if cmd[0] == "down" {
			depth += num
		} else {
			// up
			depth -= num
		}
	}

	fmt.Println("Result 1: ", horizontalPosition*depth)

	// Part 2
	var horizontalPosition1, depth1, aim int
	for _, i := range strings.Split(string(k), "\n") {
		cmd := strings.Split(i, " ")
		num, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalf("unable to convert string to int %v", err)
		}

		if cmd[0] == "forward" {
			horizontalPosition1 += num
			depth1 = depth1 + (aim * num)
		} else if cmd[0] == "down" {
			aim += num
		} else {
			// up
			aim -= num
		}
	}

	fmt.Println("Result 2: ", horizontalPosition1 * depth1)
}

// Cheers!
