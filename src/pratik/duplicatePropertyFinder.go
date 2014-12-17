package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fileHandle_en, _ := os.Open("brand_en.properties")
	duplicateKeymap := make(map[string]bool)

	defer fileHandle_en.Close()

	fileScanner_en := bufio.NewScanner(fileHandle_en)

	var duplicateCtr int64
	duplicateCtr = 0
	for fileScanner_en.Scan() {
		if (len(fileScanner_en.Text()) > 0) && !strings.HasPrefix(fileScanner_en.Text(), "#") && !duplicateKeymap[strings.Split(fileScanner_en.Text(), "=")[0]] {
			strToBeComapred := strings.Split(fileScanner_en.Text(), "=")[1]
			strKey := strings.Split(fileScanner_en.Text(), "=")[0]

			fileHandle_es, _ := os.Open("brand_en.properties")
			fileScanner_es := bufio.NewScanner(fileHandle_es)
			for fileScanner_es.Scan() {
				if (len(fileScanner_es.Text()) > 0) && !strings.HasPrefix(fileScanner_es.Text(), "#") {
					if strings.EqualFold(strToBeComapred, strings.Split(fileScanner_es.Text(), "=")[1]) && !strings.EqualFold(strKey, strings.Split(fileScanner_es.Text(), "=")[0]) {
						duplicateKeymap[strings.Split(fileScanner_es.Text(), "=")[0]] = true
						duplicateCtr++
						fmt.Print(duplicateCtr)
						fmt.Println(" " + fileScanner_es.Text())

					}
				}
			}
			fileHandle_es.Close()
		}

	}
}
