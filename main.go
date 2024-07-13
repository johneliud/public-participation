package main

import (
	"fmt"
	"log"
	"os"
	"public-participation/apirequest"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	// Read the content of the PDF file
	content, err := readPdf("file/ictbill.pdf")
	if err != nil {
		log.Fatal("Error reading PDF file:", err)
	}

	// Call the function to get API response with file content
	simplifiedContent, err := apirequest.GetApiResponse(content)
	if err != nil {
		log.Fatal("Error getting API response:", err)
	}

	// Read the HTML file
	htmlContent, err := os.ReadFile("billinfo.html")
	if err != nil {
		log.Fatal("Error reading HTML file:", err)
	}

	// Replace the summary content in the HTML
	updatedHTML := strings.Replace(string(htmlContent),
		"<!-- Summary content goes here -->",
		simplifiedContent, 1)

	// Write the updated HTML back to the file
	err = os.WriteFile("billinfo.html", []byte(updatedHTML), 0644)
	if err != nil {
		log.Fatal("Error writing updated HTML file:", err)
	}

	fmt.Println("HTML file updated successfully.")
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var content string
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		text, err := p.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		content += text
	}

	return content, nil
}
