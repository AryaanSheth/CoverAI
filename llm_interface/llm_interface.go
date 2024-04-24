package llm_interface

import (
	// "fmt"
	// "os"
	"log"
	// "regexp"
	// "github.com/AryaanSheth/coverlettergen"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
