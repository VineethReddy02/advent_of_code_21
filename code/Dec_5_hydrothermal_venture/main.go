package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// read the input from file
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	var k []byte
	k, err = ioutil.ReadFile(mydir + "/code/Dec_5_hydrothermal_venture/input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}

	var inputData [][]string
	for _, i := range strings.Split(string(k), "\n") {
		i := strings.Replace(i, " -> ", ",", 1)
		i = strings.Replace(i, " ", "", -1)
		data := strings.Split(i, ",")
		inputData = append(inputData, data)
	}

	// TODO implementation
}
