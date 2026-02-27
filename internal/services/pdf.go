package services

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dslipak/pdf"
)

// ExtractTextFromPDF reads a PDF file from raw bytes and extracts all text content.
func ExtractTextFromPDF(pdfBytes []byte) (string, error) {
	reader := bytes.NewReader(pdfBytes)

	pdfReader, err := pdf.NewReader(reader, int64(len(pdfBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to create PDF reader: %w", err)
	}

	var textBuilder strings.Builder

	numPages := pdfReader.NumPage()
	if numPages == 0 {
		return "", fmt.Errorf("PDF has no pages")
	}

	for i := 1; i <= numPages; i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			// Skip pages that can't be read, continue with others
			continue
		}

		textBuilder.WriteString(text)
		textBuilder.WriteString("\n")
	}

	result := strings.TrimSpace(textBuilder.String())
	if result == "" {
		return "", fmt.Errorf("no text could be extracted from the PDF")
	}

	return result, nil
}
