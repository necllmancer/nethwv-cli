package main

import (
	"flag"
	"fmt"
	"nethwv-cli/pkg/github"
	"nethwv-cli/pkg/pdf"
	"os"
	"path/filepath"
)

func main() {
	// コマンドラインオプションの定義
	branch := flag.String("b", "", "Branch to clone")
	tag := flag.String("t", "", "Tag to clone")
	directory := flag.String("d", "", "Specific directory to retrieve files from")
	flag.Parse()

	// 残りの引数からリポジトリと出力PDFファイル名を取得
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: nethwv [options] <user/repo> <output.pdf>")
		os.Exit(1)
	}

	repoURL, outputPDF := args[0], args[1]

	client := github.NewClient(nil) // GitHubクライアントの初期化

	// リポジトリをクローン（ブランチまたはタグの指定がある場合はそれを使用）
	branchOrTag := *branch
	if *tag != "" {
		branchOrTag = *tag
	}
	err := client.CloneRepo(repoURL, "tmp", branchOrTag) // 一時ディレクトリにクローン
	if err != nil {
		fmt.Printf("Error cloning repository: %s\n", err)
		os.Exit(1)
	}

	// 特定のディレクトリが指定されている場合はそのディレクトリのみからファイルを取得
	var files []string
	if *directory != "" {
		files, err = client.RetrieveFiles("tmp", *directory)
	} else {
		files, err = client.RetrieveFiles("tmp", "")
	}
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
