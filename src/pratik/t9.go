package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var t9map map[string]int
	t9map["a"] = 2
	t9map["b"] = 22
	t9map["c"] = 222
	t9map["d"] = 3
	t9map["e"] = 33
	t9map["f"] = 333
	t9map["g"] = 4
	t9map["h"] = 44
	t9map["i"] = 444
	t9map["j"] = 5
	t9map["k"] = 55
	t9map["l"] = 555
	t9map["m"] = 6
	t9map["n"] = 66
	t9map["o"] = 666
	t9map["p"] = 7
	t9map["q"] = 77
	t9map["r"] = 777
	t9map["s"] = 7777
	t9map["t"] = 8
	t9map["u"] = 88
	t9map["v"] = 888
	t9map["w"] = 9
	t9map["x"] = 99
	t9map["y"] = 999
	t9map["z"] = 9999

	fileHandle, _ := os.Open("cs.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	fileScanner.Scan()

	numCases, _ := strconv.ParseInt(fileScanner.Text(), 10, strconv.IntSize)
	var itr int64

	for itr = 1; itr <= numCases; itr++ {
		fileScanner.Scan()
		byteArray := []byte(fileScanner.Text())
		for idx, str := range byteArray {
			fmt.Printf(string(str))
			fmt.Printf(string(idx))
		}
	}
}
