package main

import (
	"fmt"
	"nethwv-cli/pkg/github"
	"nethwv-cli/pkg/pdf"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: nethwv-cli <repo-url> <output-pdf>")
		os.Exit(1)
	}

	repoURL := os.Args[1]
	outputPDF := os.Args[2]

	// Retrieve files from GitHub repository
	files, err := github.RetrieveFiles(repoURL)
	if err != nil {
		fmt.Printf("Error retrieving files: %s\n", err)
		os.Exit(1)
	}

	// Generate PDF
	if err := pdf.GeneratePDF(files, outputPDF); err != nil {
		fmt.Printf("Error generating PDF: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully:", outputPDF)
}
