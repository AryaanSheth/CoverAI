package main

import (
	"fmt"
	"os"
	"regexp"
	"log"
	
)

func main() {
	// open text file 
	f, err := os.Open("output.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	buffer := make([]byte, 1024)
	var content string
	for n, err := f.Read(buffer); n > 0; n, err = f.Read(buffer) {
		if err != nil {
			log.Fatal(err)
		}
		content += string(buffer[:n])
	}

	// clean text
	// remove all non-ascii characters and newlines/tabs
	re := regexp.MustCompile("[^\\x00-\\x7F]+|\\n|\\t")

	cleaned := re.ReplaceAllString(content, "")

	// write cleaned text to new file
	f, err = os.Create("cleaned.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(cleaned)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Text cleaned and written to cleaned.txt")
	


}
