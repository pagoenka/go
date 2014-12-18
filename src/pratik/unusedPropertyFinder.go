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

	fileHandle_en, _ := os.Open("brand_en.properties")
	defer fileHandle_en.Close()
	fileScanner_en := bufio.NewScanner(fileHandle_en)

	for fileScanner_en.Scan() {
		if (len(fileScanner_en.Text()) > 0) && !strings.HasPrefix(fileScanner_en.Text(), "#") {
			strToBeSearched := strings.Split(fileScanner_en.Text(), "=")[0]

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

			for _, file := range fileList {
				if !strings.HasSuffix(file, ".properties") {
					fileHandle, _ := os.Open(file)
					fileScanner := bufio.NewScanner(fileHandle)
					for fileScanner.Scan() {
						if len(fileScanner.Text()) > 0 {
							if strings.Contains(fileScanner.Text(), strToBeSearched) {
								fmt.Println(strToBeSearched)
								fmt.Println(file)
								break
							}
						}
					}
					fileHandle.Close()
				}
			}
		}
	}

}
