package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var digits = ""

const digitFolderName = "./billion-digits"

// fields must be uppercase to pass through response
type SearchRes struct {
	Index int `json:"index"`
}

func main() {
	loadPi()

	fmt.Printf("Length: %d\n", len(digits))
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/search", search)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Println("Listening on :" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	q, ok := r.URL.Query()["q"]

	if !ok || len(q[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// the query
	query := q[0]

	w.Header().Set("Content-Type", "application/json")

	// create json result object
	result := SearchRes{
		Index: (strings.Index(digits, query)) - 1,
	}

	json.NewEncoder(w).Encode(result)

	end := time.Now()
	fmt.Printf("\n----------------\nThis took %v ms\n----------------\n", end.Sub(start).Seconds())
	log.Println("Url Param 'q' is: " + string(query))
}

// Loads the entire contents of the /billion-digits
// files into the 'digits' global variable
// tldr; this loads pi into a string global called 'digits'
func loadPi() {
	defer fmt.Println("Done reading in pi")

	files, err := getDigitFileChunkNames()
	// fmt.Printf("FIles: %v", files)

	if err != nil {
		log.Fatal(err)
	}

	for _, fileN := range files {

		filePath := digitFolderName + "/" + fileN
		dat, _ := ioutil.ReadFile(filePath)
		digits += strings.TrimSpace(string(dat))
		fmt.Printf("Loaded %v...\n" + fileN)
	}

}

// Returns a list of the "*.txt" files in
// the digit directory
// For this to work, there must exist files in this numbering system in the 'digitFolderName' folder:
/**
0.txt
1.txt
2.txt
3.txt
...
*/
func getDigitFileChunkNames() ([]string, error) {
	files, err := ioutil.ReadDir(digitFolderName)

	fileNames := []string{}

	if err != nil {
		return nil, errors.New("Error in reading file.")
	}

	for _, f := range files {
		if f.Name()[len(f.Name())-4:] == ".txt" {
			fileNames = append(fileNames, f.Name())
		}
	}

	numberOfFiles := len(fileNames)
	sortedFileNames := []string{}

	for i := 0; i < numberOfFiles; i++ {
		sortedFileNames = append(sortedFileNames, strconv.Itoa(i)+".txt")
	}

	// MUST be sorted
	return sortedFileNames, nil
}
