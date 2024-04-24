package content_extraction

import (
	"log"
	"os"
	"regexp"

	"github.com/unidoc/unidoc/pdf/extractor"
	"github.com/unidoc/unidoc/pdf/model"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToFile(filename string, content string) {
	f, err := os.Create(filename)
	HandleError(err)
	defer f.Close()

	_, err = f.WriteString(content)
	HandleError(err)
}

func CleanText(content string) string {
	re := regexp.MustCompile("[^\\x00-\\x7F]+|\\n|\\t")
	return re.ReplaceAllString(content, "")
}

func ExtractTextFromPDF(path string) (string, error) {
	f, err := os.Open(path)
	HandleError(err)

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	HandleError(err)

	var text string
	numPages, _ := pdfReader.GetNumPages()
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		HandleError(err)

		ex, err := extractor.New(page)
		HandleError(err)

		pageText, err := ex.ExtractText()
		HandleError(err)

		text += pageText
	}

	return text, nil
}

// func main() {

// 	// extract text from pdf
// 	text, err := ExtractTextFromPDF("sample.pdf")
// 	HandleError(err)

// 	// clean text
// 	cleaned := cleanText(text)

// 	// write cleaned text to new file
// 	WriteToFile("cleaned.txt", cleaned)

// 	fmt.Println("Text cleaned and written to cleaned.txt")

// }
