package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	k, err = ioutil.ReadFile(mydir + "/code/Dec_3_binary_diagnostic/input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}

	// Part 1
	inputData := strings.Split(string(k), "\n")
	var inputNumbers = make([]int, len(inputData[0]))
	for _, i := range inputData {
		for ind, j := range i {
			num, err := strconv.Atoi(fmt.Sprintf("%c", j))
			if err != nil {
				log.Fatalf("unable to convert string to int %v", err)
			}

			if num == 1 {
				inputNumbers[ind] += 1
			}
		}
	}

	// most common bit & least common bit
	var mostCommonBits, leastCommonBits string
	for _, i := range inputNumbers {
		if i > len(inputData)/2 {
			mostCommonBits += "1"
			leastCommonBits += "0"
		} else {
			mostCommonBits += "0"
			leastCommonBits += "1"
		}
	}

	mostCommon, err := strconv.Atoi(mostCommonBits)
	if err != nil {
		log.Fatalf("unable to convert string to int %v", err)
	}

	leastCommon, err := strconv.Atoi(leastCommonBits)
	if err != nil {
		log.Fatalf("unable to convert string to int %v", err)
	}

	mostCommonDecimal := binaryToDecimal(mostCommon)
	leastCommonDecimal := binaryToDecimal(leastCommon)

	fmt.Println("Result 1: ", mostCommonDecimal*leastCommonDecimal)

	// Part 2
	// oxygen rating
	oxyRating := inputData
	for ind := 0; len(oxyRating) > 0 && ind < len(oxyRating[0]); ind++ {
		var one, two []string
		for _, i := range oxyRating {
			if fmt.Sprintf("%c", i[ind]) == "1" {
				one = append(one, i)
			} else {
				two = append(two, i)
			}
		}

		if len(one) < len(two) {
			one = two
		}
		oxyRating = one
	}

	// co2 rating
	co2Rating := inputData
	for ind := 0; len(co2Rating) > 1 && ind < len(co2Rating[0]); ind++ {
		var one, two []string
		for _, i := range co2Rating {
			if fmt.Sprintf("%c", i[ind]) == "1" {
				one = append(one, i)
			} else {
				two = append(two, i)
			}
		}

		if len(one) >= len(two) {
			one = two
		}
		co2Rating = one
	}

	o, err :=  strconv.Atoi(oxyRating[0])
	if err != nil {
		log.Fatalf("unable to convert string to int %v", err)
	}

	c, err :=  strconv.Atoi(co2Rating[0])
	if err != nil {
		log.Fatalf("unable to convert string to int %v", err)
	}

	oxyRatingDecimal := binaryToDecimal(o)
	co2RatingDecimal := binaryToDecimal(c)

	fmt.Println("Result 2: ", oxyRatingDecimal*co2RatingDecimal)
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}

// Cheers!