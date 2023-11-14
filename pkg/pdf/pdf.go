package pdf

import (
	"fmt"
	"nethwv-cli/pkg/filefetcher"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(filePaths []string, output string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)

	for _, path := range filePaths {
		content, err := filefetcher.FetchFileContent(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			continue // Skip this file and continue with the next
		}
		if content != "" {
			pdf.AddPage()
			pdf.SetXY(10, 10) // Set text start position
			pdf.MultiCell(0, 10, fmt.Sprintf("File: %s\n\n%s", path, content), "", "", false)
		}
	}

	err := pdf.OutputFileAndClose(output)
	if err != nil {
		fmt.Printf("Error saving PDF: %v\n", err)
		return err
	}

	return nil
}
