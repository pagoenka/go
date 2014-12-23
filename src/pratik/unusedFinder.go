package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	searchDir := "/Volumes/Data/lrs-amit/lrsadmin/lrs-admin-web/src/js"

	fileHandle_en, _ := os.Open("brand_en.properties")
	defer fileHandle_en.Close()
	fileScanner_en := bufio.NewScanner(fileHandle_en)

	fileList := []string{}
	strToBeSearched := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("number of files", len(fileList))

	for fileScanner_en.Scan() {
		if (len(fileScanner_en.Text()) > 0) && !strings.HasPrefix(fileScanner_en.Text(), "#") {
			strToBeSearched = append(strToBeSearched, strings.Split(fileScanner_en.Text(), "=")[0])
		}
	}
	var c chan map[string]int = make(chan map[string]int, len(fileList))
	//var lok chan int = make(chan int, len(fileList))
	// var wg sync.WaitGroup
	// wg.Add(len(fileList))
	for _, file := range fileList {
		go finder(file, strToBeSearched, c)
	}
	// for _ := range fileList {
	// 	msg := <-lok
	// }
	finalSummary := make(map[string]int)
	var summ map[string]int
	for _ = range fileList {
		summ = <-c
		for key, value := range summ {
			finalSummary[key] = finalSummary[key] + value
		}
	}
	for key, value := range finalSummary {
		fmt.Println(key, value)
	}
}

func finder(file string, strToBeSearched []string, c chan map[string]int) {
	fmt.Print(".")
	data, _ := ioutil.ReadFile(file)
	summary := make(map[string]int)

	for _, value := range strToBeSearched {
		if strings.Contains(string(data), value) {
			summary[value] = summary[value] + 1
		}
	}
	c <- summary
	//lok <- 1
	//wg.Done()
}
