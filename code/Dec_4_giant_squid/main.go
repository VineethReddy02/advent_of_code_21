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
	k, err = ioutil.ReadFile(mydir + "/code/Dec_4_giant_squid/input.txt")
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}

	// read the data into list of 2D arrays
	var keys []string
	var players [][][]string
	for ind, i := range strings.Split(string(k), "\n\n") {
		if ind == 0 {
			keys = strings.Split(i, ",")
			continue
		}
		rows := strings.Split(i, "\n")
		var k123 = make([][]string, 5)
		for m := 0; m < len(rows); m++ {
			nums := strings.Split(rows[m], " ")
			for ind1, num := range nums {
				if num == "" {
					nums = append(nums[:ind1], nums[ind1+1:]...)
				}
			}

			if len(nums) > 1 {
				k123[m] = nums
			}
		}

		players = append(players, k123)
	}

	// process the keys vs player cards
	var WinningPlayer, WinningKey int
	var resultSetTotals, resultSetKeys []int
	var wonPlayers = make(map[int]int)
	for ind, _ := range keys {
		var total int
		for p := 0; p < len(players); p++ {
			// once a player is won
			// ignore them in further processing
			if wonPlayers[p] > 0 {
				continue
			}

			// mark the keys
			for m := 0; m < 5; m++ {
				for n := 0; n < 5; n++ {
					if keys[ind] == players[p][m][n] {
						players[p][m][n] = "#"
					}
				}
			}
		}

		// Validate the players won from the key picked
		for p := 0; p < len(players); p++ {
			// once a player is won
			// ignore them in further processing
			if wonPlayers[p] > 0 {
				continue
			}

			// validating specific to a player
			var count, count1 int
			for m := 0; m < 5; m++ {
				count = 0
				count1 = 0
				for n := 0; n < 5; n++ {
					if players[p][m][n] == "#" {
						count++
					}
					if players[p][n][m] == "#" {
						count1++
					}
				}

				if count == 5 || count1 == 5 {
					WinningPlayer = p
					WinningKey, err = strconv.Atoi(keys[ind])
					if err != nil {
						log.Fatalf("unable to convert string to int %v", err)
					}
					// mark the player won
					wonPlayers[p] = wonPlayers[p] + 1
					// A player may have multiple rows & columns getting marked by a winning key
					// so once the row/column is found marked ignore other rows, columns.
					break
				}
			}
		}

		// calculate the sum for the player won
		for m := 0; m < 5; m++ {
			for n := 0; n < 5; n++ {
				if players[WinningPlayer][m][n] != "#" {
					num, err := strconv.Atoi(players[WinningPlayer][m][n])
					if err != nil {
						log.Fatalf("unable to convert string to int %v", err)
					}
					total += num
				}
			}
		}

		// store the sum and key values for the players won
		if total != 0 && WinningKey != 0 {
			resultSetTotals = append(resultSetTotals, total)
			resultSetKeys = append(resultSetKeys, WinningKey)
		}
	}

	// calculate the final output i.e. total*key
	var results []int
	for ik := 0; ik < len(resultSetTotals); ik++ {
		results = append(results, resultSetTotals[ik]*resultSetKeys[ik])
	}

	fmt.Println("Result 1: ", results[0])
	fmt.Println("Result 2: ", results[len(resultSetTotals)-1])
}

// Cheers!
