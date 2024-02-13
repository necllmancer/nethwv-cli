package pdf

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"nethwv-cli/pkg/filefetcher"
)

func GeneratePDF(filePaths []string, outputDir string) error {
	for _, path := range filePaths {
		content, err := filefetcher.FetchFileContent(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			continue
		}
		if content != "" {
			pdf := gofpdf.New("P", "mm", "A4", "")
			pdf.SetFont("Arial", "", 12)
			pdf.AddPage()
			pdf.SetXY(10, 10)
			pdf.MultiCell(0, 10, fmt.Sprintf("File: %s\n\n%s", path, content), "", "", false)

			// Create a RAG-friendly filename with directory structure
			filename := createRAGFriendlyFilenameWithDir(path, outputDir)
			outputPath := filepath.Join(outputDir, filename)

			err := pdf.OutputFileAndClose(outputPath)
			if err != nil {
				fmt.Printf("Error saving PDF %s: %v\n", outputPath, err)
				return err
			}
		}
	}

	return nil
}

// createRAGFriendlyFilenameWithDir generates a filename that includes the directory structure.
func createRAGFriendlyFilenameWithDir(path, baseDir string) string {
	relativePath, err := filepath.Rel(baseDir, path)
	if err != nil {
		// If unable to get relative path, use the base name
		return filepath.Base(path)
	}

	// Remove "../" or "./" from the beginning of the relative path
	relativePath = strings.TrimPrefix(relativePath, "../")
	relativePath = strings.TrimPrefix(relativePath, "./")

	name := strings.TrimSuffix(relativePath, filepath.Ext(relativePath))
	name = strings.ReplaceAll(name, string(filepath.Separator), "_")
	return fmt.Sprintf("%s.pdf", name)
}


// createRAGFriendlyFilename generates a filename that is optimized for RAG.
func createRAGFriendlyFilename(path string) string {
	base := filepath.Base(path)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	return fmt.Sprintf("%s_RAG.pdf", name)
}
