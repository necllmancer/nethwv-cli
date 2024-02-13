package main

import (
	"flag"
	"fmt"
	"nethwv-cli/pkg/github"
	"nethwv-cli/pkg/pdf"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	gitCmd := flag.NewFlagSet("git", flag.ExitOnError)
	localCmd := flag.NewFlagSet("local", flag.ExitOnError)

	// Git subcommand flags
	branch := gitCmd.String("b", "", "Branch to clone")
	tag := gitCmd.String("t", "", "Tag to clone")
	directory := gitCmd.String("d", "", "Specific directory to retrieve files from")

	// Local subcommand flags
	localPath := localCmd.String("p", "", "Path to local directory")

	if len(os.Args) < 2 {
		fmt.Println("Usage: nethwv <command> [options]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "git":
		handleGitCommand(gitCmd, branch, tag, directory)
	case "local":
		handleLocalCommand(localCmd, localPath)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func handleGitCommand(gitCmd *flag.FlagSet, branch, tag, directory *string) {
	gitCmd.Parse(os.Args[2:])
	args := gitCmd.Args()
	if len(args) < 2 {
		fmt.Println("Usage: nethwv git [options] <user/repo> <output.pdf>")
		os.Exit(1)
	}

	// Existing implementation for git subcommand
	// コマンドラインオプションの定義
	branch = flag.String("b", "", "Branch to clone")
	tag = flag.String("t", "", "Tag to clone")
	directory = flag.String("d", "", "Specific directory to retrieve files from")
	flag.Parse()

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

func handleLocalCommand(localCmd *flag.FlagSet, localPath *string) {
	localCmd.Parse(os.Args[2:])
	args := localCmd.Args()
	if len(args) < 1 {
		fmt.Println("Usage: nethwv local -p <path> <output.pdf>")
		os.Exit(1)
	}

	outputPDF := args[0]

	// Validate local path
	if *localPath == "" {
		fmt.Println("Error: Local path is required")
		os.Exit(1)
	}

	// Check if the output directory exists and create if not
	outputDir := filepath.Dir(outputPDF)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating output directory: %s\n", err)
			os.Exit(1)
		}
	}

	var files []string
	err := filepath.Walk(*localPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error retrieving files: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Number of files retrieved:", len(files))

	// PDF generation
	if err := pdf.GeneratePDF(files, outputPDF); err != nil {
		fmt.Printf("Error generating PDF: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully:", outputPDF)
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
