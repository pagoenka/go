package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	searchDir := "/Volumes/Data/lrs-amit/lrsadmin/lrs-admin-web/src/js"
	properyFoundMap := make(map[string]bool)

	fileHandle_en, _ := os.Open("brand_en.properties")
	defer fileHandle_en.Close()
	fileScanner_en := bufio.NewScanner(fileHandle_en)

	fileList := []string{}
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

	for fileScanner_en.Scan() {
		if (len(fileScanner_en.Text()) > 0) && !strings.HasPrefix(fileScanner_en.Text(), "#") {
			strToBeSearched := strings.Split(fileScanner_en.Text(), "=")[0]

			properyFoundMap[strToBeSearched] = false
			for _, file := range fileList {
				if !strings.HasSuffix(file, ".properties") {
					go finder(file, strToBeSearched, properyFoundMap)
				}
			}
		}
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println("Unused keys ----------------->")
	for k, _ := range properyFoundMap {
		if !properyFoundMap[k] {
			fmt.Println(k)
		}
	}
}

func finder(file string, strToBeSearched string, properyFoundMap map[string]bool) {
	if !properyFoundMap[strToBeSearched] {
		fileHandle, _ := os.Open(file)
		fileScanner := bufio.NewScanner(fileHandle)
		for fileScanner.Scan() {
			if len(fileScanner.Text()) > 0 {
				if strings.Contains(fileScanner.Text(), strToBeSearched) {
					//fmt.Println(strToBeSearched)
					//fmt.Println(file)
					fmt.Print("...")
					properyFoundMap[strToBeSearched] = true
					break
				}
			}
		}
		fileHandle.Close()
	}
}
