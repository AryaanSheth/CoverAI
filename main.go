package main

import (
	llm_interface "github.com/AryaanSheth/coverlettergen/llm_interface"
	content_extraction "github.com/AryaanSheth/coverlettergen/pdf_to_text"
)

func main() {
	filename := "sample.pdf"

	text, err := content_extraction.ExtractTextFromPDF(filename)
	llm_interface.HandleError(err)

	cleaned := content_extraction.CleanText(text)

	content_extraction.WriteToFile("cleaned.txt", cleaned)

}
