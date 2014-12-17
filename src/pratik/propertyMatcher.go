package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fileHandle_en, _ := os.Open("brand_en.properties")
	fileHandle_es, _ := os.Open("brand_es.properties")

	defer fileHandle_en.Close()
	defer fileHandle_es.Close()

	fileScanner_en := bufio.NewScanner(fileHandle_en)
	fileScanner_es := bufio.NewScanner(fileHandle_es)

	var counter int64
	for counter = 1; fileScanner_en.Scan(); counter++ {
		fileScanner_es.Scan()

		if !(strings.EqualFold(strings.Split(fileScanner_en.Text(), "=")[0], strings.Split(fileScanner_es.Text(), "=")[0])) {
			fmt.Print(counter)
			fmt.Println(" " + fileScanner_en.Text())
			counter++
			break
		}
	}
}
