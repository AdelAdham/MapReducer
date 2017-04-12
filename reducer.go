// Works on part of dataset and saves result in file as a reducer.

package main

// Import necessary package to accomplish your work.
import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"sort"
	"io/ioutil"
	"strings"
	"log"
)

type SortWikiResult struct {
	searchWiki map[string]int
	Keys       []string
}

func (sw *SortWikiResult) Len() int {
	// TODO: Implement Len function.
	return len(sw.searchWiki)
}

func (sw *SortWikiResult) Less(i, j int) bool {
	// TODO: Implement Less function.
	return sw.searchWiki[sw.Keys[i]] > sw.searchWiki[sw.Keys[j]]

}

func (sw *SortWikiResult) Swap(i, j int) {
	// TODO: Implement Swap function.
	sw.Keys[i], sw.Keys[j] = sw.Keys[j], sw.Keys[i]

}

func sortArticles(searchWiki map[string]int) []string {
	// TODO: Implement sortKeys function.
	var Keys []string

	for key, _ := range searchWiki {
		Keys = append(Keys, key)
	}

	SWR := SortWikiResult{searchWiki, Keys}
	sort.Stable(&SWR)

	return Keys
}

func getMappersFiles(jobID int) []string {
	// TODO: get a list of path of files were produced
	// by mappers in path /mnt/datanode/tmp/jobID.
	path := "/mnt/datanode/tmp/" + strconv.Itoa(jobID)

	files, _ := ioutil.ReadDir(path)

	var arr []string
	for _, f := range files {
		arr = append(arr, f.Name())
	}

	return arr
}

func readFile(path string, searchWiki map[string]int) {
	// TODO: read file which was written by the mapper and
	// add data to searchWiki.

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		arr := strings.Fields(line)
		v, _ := strconv.Atoi(arr[1])
		searchWiki[arr[0]] = v
	}
}

func main() {
	// TODO: get one arguments in format jobID.
	strjobID := os.Args[1]

	jobID, _ := strconv.Atoi(strjobID)

	// TODO: call getMappersFiles and save result in array.
	mapFiles := getMappersFiles(jobID)

	// TODO: use readFile for files of mappers.
	myMap := make(map[string]int)
	for i := 0; i < len(mapFiles); i++ {
		readFile("/mnt/datanode/tmp/"+strconv.Itoa(jobID)+"/"+mapFiles[i], myMap)

	}

	// TODO: perform sort on results.
	sortedWiki := sortArticles(myMap)

	// TODO: save only results of top 100 in output file
	// in this path pattern /mnt/datanode/tmp/jobID/output.

	var top60 []string
	for i := 0; i < len(sortedWiki); i++ {
		top60 = append(top60, sortedWiki[i])
	}
	top := strings.Join(top60, "\n")

	err := ioutil.WriteFile("/mnt/datanode/tmp/"+strjobID+"/output", []byte(top), 0644)

	if err == nil {
		fmt.Println("Data has been written")
	}

}
