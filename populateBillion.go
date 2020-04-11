package main

// This file takes a single large file of pi and splits it
// into 90MB chunks that are subsequently placed into
// the ./billion-digits folder.

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

var digits []byte = []byte{}

func main() {
	const fileN = "./pi-billion.txt"

	dat, _ := ioutil.ReadFile(fileN)
	digits = dat
	fmt.Println("Done reading in pi")

	// break apart digits into 90MB chunks and create files for them in "billion digits"
	counter := 0

	const ninetyMillion = 90_000_000
	for i := range digits {
		if i%ninetyMillion == 0 && i > 0 {
			partNameFile := "./billion-digits/" + strconv.Itoa(counter) + ".txt"
			ioutil.WriteFile(partNameFile, digits[i-ninetyMillion:i], 0644)
			counter++
		}

	}

	// gather the last part of pi into its file
	partNameFile := "./billion-digits/" + strconv.Itoa(counter) + ".txt"
	startIndex := counter * ninetyMillion

	ioutil.WriteFile(partNameFile, digits[startIndex:], 0644)

}
