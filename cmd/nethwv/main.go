package main

import (
	"fmt"
	"net/http"
	"nethwv-cli/pkg/github"
	"nethwv-cli/pkg/pdf"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		// fmt.Println("Usage: nethwv-cli <repo-url> <output-pdf>")
		// os.Exit(1)
		args = []string{"", "psf/requests", "default.pdf"}
	}

	repoURL, outputPDF := args[1], args[2]

	// GitHubクライアントの初期化
	client := github.NewClient(http.DefaultClient)

	// GitHubリポジトリからファイルを取得
	fileUrls, err := client.RetrieveFiles(repoURL)
	if err != nil {
		fmt.Printf("Error retrieving files from GitHub: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("len(fileUrls) = ", len(fileUrls))
	// PDFを生成
	if err := pdf.GeneratePDF(fileUrls, outputPDF); err != nil {
		fmt.Printf("Error generating PDF: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully:", outputPDF)
}
