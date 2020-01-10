package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"
)

// Creation of the struct representing the job and the result
type Job struct {
	id         int
	textString []string
}

type Result struct {
	job     Job
	wordOcc map[string]int
}

// Jobs and results buffered channel to receiving the jobs and writes the output
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// Count the word occurrence
func wordCount(inputString []string) map[string]int {
	counts := make(map[string]int)

	for _, value := range inputString {
		wordList := strings.Fields(value)
		for _, word := range wordList {
			cleanedWord := cleanString(word)
			_, ok := counts[cleanedWord]
			if ok {
				counts[cleanedWord]++
			} else {
				counts[cleanedWord] = 1
			}
		}

	}
	return counts
}

//Create a worker read from the jobs channel and creates a Result struct using the current job
func worker(wg *sync.WaitGroup) {

	for job := range jobs {

		output := Result{job, wordCount(job.textString)}
		results <- output
	}
	wg.Done()
}

// Create a pool of worker Goroutines. It takes the number of workers inited in main func
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(results)
}

// The function below read the file content then set the job for the workers
func allocate(listDir []string) {

	for index, value := range listDir {
		textString := readFile(value)
		job := Job{index, textString}
		jobs <- job
	}
	fmt.Println("Test")
	close(jobs)
}

/* This function read the results channel and print the output */
func result(done chan bool) {

	for result := range results {

		for index, element := range result.wordOcc {
			fmt.Println(index, "=>", element)
		}
	}

	done <- true
}

/*
	This function list all the path of the file and file in the sub-folder( if available)
	dirPath : Take the user input to scan for files
	newFilePath : Output the path of all file in the user's path
*/
func listDir(dirPath string) (newFilePath []string) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			// Call this function again if the child is a folder
			listDir(dirPath + "/" + f.Name())
		} else {
			fullPath := dirPath + "/" + f.Name()
			newFilePath = append(newFilePath, fullPath)
		}
	}
	fmt.Println(newFilePath)
	return newFilePath
}

/* Remove all the non-alphanumeric characters from a string */
func cleanString(inputString string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(inputString, "")
	return processedString

}

/* Get the content of the file read from file Path */
func readFile(filePath string) []string {

	fileText, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
	}

	fileContent := strings.Fields(string(fileText))

	return fileContent

}
func main() {

	dirPath := flag.String("p", "", "Please input the directory path")
	flag.Parse()
	listDir := listDir(*dirPath)

	go allocate(listDir)
	done := make(chan bool)
	go result(done)
	workersAmount := 100
	createWorkerPool(workersAmount)
	<-done

}
