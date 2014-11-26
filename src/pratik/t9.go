package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	t9map := map[string]int{
		"a": 2,
		"b": 22,
		"c": 222,
		"d": 3,
		"e": 33,
		"f": 333,
		"g": 4,
		"h": 44,
		"i": 444,
		"j": 5,
		"k": 55,
		"l": 555,
		"m": 6,
		"n": 66,
		"o": 666,
		"p": 7,
		"q": 77,
		"r": 777,
		"s": 7777,
		"t": 8,
		"u": 88,
		"v": 888,
		"w": 9,
		"x": 99,
		"y": 999,
		"z": 9999,
	}

	fileHandle, _ := os.Open("cs.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	fileScanner.Scan()

	numCases, _ := strconv.ParseInt(fileScanner.Text(), 10, strconv.IntSize)
	var itr int64

	for itr = 1; itr <= numCases; itr++ {
		fmt.Print("Case #", itr, ": ")
		fileScanner.Scan()
		arrStrings := strings.Split(fileScanner.Text(), "")
		for str := range arrStrings {
			fmt.Print(t9map[arrStrings[str]])
		}
		fmt.Println("")
	}
}
