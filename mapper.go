// Works on part of dataset and saves result in file as mapper.

package main

// Import necessary package to accomplish your work.
import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"io/ioutil"
	"log"
	"fmt"
	"math"
)

type Directories struct {
	Dirs []string
}

func getFilesList(searchDir string) []string {
	// TODO: Reads the file contains a list of paths in path
	// /mnt/dataset/fileList.

	file, err := os.Open(searchDir)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // excutes it the last statement till all other stats finish

	var arr []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)

	}

	return arr
}

func searchInFile(path string, keyword string, searchMap map[string]int) {
	// TODO: open File and read it line by line and count
	// number of keyword and save result in map.

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)

		if strings.Contains(line, keyword) {
			searchMap[path] += strings.Count(line, keyword)
		}
	}

}

func main() {
	// TODO: get three arguments in format mapperID jobID keyword.
	mapperID,_ := strconv.Atoi(string(os.Args[1]))
	jobID ,_ := strconv.Atoi(string(os.Args[2]))
	keyword := os.Args[3]
	keyword = strings.ToLower(keyword)

	// TODO: call getFilesList and save result in array
	// path for Wikipedia is at path /mnt/datanode/dataset.
	filelist := getFilesList("/mnt/datanode/dataset/fileList")

	// TODO: determine the start and the end of working data.
	var start, end float64
	if mapperID == 1 {
		start = 0
		end = math.Ceil(float64(len(filelist)) / float64(3))
	} else if mapperID == 2 {
		start = math.Ceil(float64(len(filelist)) / float64(3))
		end = float64(len(filelist)) * (float64(2) / float64(3))
	} else {
		start = float64(len(filelist)) * (float64(2) / float64(3))
		end = float64(len(filelist))
	}

	// TODO: create a map with string key for path and int for values.
	myMap := make(map[string]int)


	// TODO: loop on files and use call searchInFile.
	for i:= int(start); i < int(end); i++ {
		searchInFile(filelist[i],keyword,myMap)
	}

	// TODO: save results in any format in output file with name of
	// mapperID in this path pattern /mnt/datanode/tmp/jobID/mapperID.
	var data []string
	for key, value := range myMap {
		data = append(data, key + " " + strconv.Itoa(value))
	}

	data1 := strings.Join(data, "\n")

	err := ioutil.WriteFile("/mnt/datanode/tmp/" + strconv.Itoa(jobID) + "/" + strconv.Itoa(mapperID), []byte(data1), 0644)

	if err == nil {
		fmt.Println("Data has been written")
	}
}