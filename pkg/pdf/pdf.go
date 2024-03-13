package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"nethwv-cli/pkg/filefetcher"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(filePaths []string, outputPath string, ignorePatterns []string) error {
	err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error creating output directory: %v", err)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)

	for _, path := range filePaths {
		if isIgnored(path, ignorePatterns) {
			continue
		}

		content, err := filefetcher.FetchFileContent(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			continue
		}

		if content != "" {
			// 内容を最初の1000文字に制限する
			if len(content) > 1000 {
				content = content[:1000] + "..."
			}

			pdf.AddPage()
			pdf.SetXY(10, 10)
			pdf.MultiCell(0, 10, fmt.Sprintf("File: %s\n\n%s", path, content), "", "LT", false)
		}
	}

	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		fmt.Printf("Error saving PDF %s: %v\n", outputPath, err)
		return err
	}

	return nil
}

func isIgnored(path string, ignorePatterns []string) bool {
	for _, pattern := range ignorePatterns {
		matched, err := filepath.Match(pattern, path)
		if err != nil {
			fmt.Printf("Invalid ignore pattern: %s\n", pattern)
			continue
		}
		if matched {
			return true
		}
	}
	return false
}
