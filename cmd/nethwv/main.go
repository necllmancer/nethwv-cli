package main

import (
	"fmt"
	"nethwv-cli/pkg/github"
	"nethwv-cli/pkg/pdf"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		args = []string{"", "psf/requests", "default.pdf"} // デフォルトの引数
	}

	repoURL, outputPDF := args[1], args[2]

	client := github.NewClient(nil) // GitHubクライアントの初期化

	// リポジトリをシャロークローン (--depth 1 を使用)
	err := client.CloneRepo(repoURL, "tmp") // 一時ディレクトリにクローン
	if err != nil {
		fmt.Printf("Error cloning repository: %s\n", err)
		os.Exit(1)
	}

	// ファイルの一覧を取得
	files, err := client.RetrieveFiles("tmp") // 一時ディレクトリからファイル一覧を取得
	if err != nil {
		fmt.Printf("Error retrieving files: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Number of files retrieved:", len(files))

	// PDFを生成
	if err := pdf.GeneratePDF(files, outputPDF); err != nil {
		fmt.Printf("Error generating PDF: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully:", outputPDF)

	// クリーンアップ: 一時ディレクトリを削除
	os.RemoveAll("tmp")
}

func printTree(path, indent string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	files, err := f.ReadDir(-1)
	if err != nil {
		return err
	}

	for i, file := range files {
		filePath := filepath.Join(path, file.Name())
		isLast := i == len(files)-1

		fmt.Print(indent)
		if isLast {
			fmt.Print("└── ")
		} else {
			fmt.Print("├── ")
		}

		if file.IsDir() && file.Name() != ".git" {
			fmt.Println(file.Name())
			var newIndent string
			if isLast {
				newIndent = indent + "    "
			} else {
				newIndent = indent + "│   "
			}
			err := printTree(filePath, newIndent)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(file.Name())
		}
	}

	return nil
}
