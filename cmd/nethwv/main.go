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
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "git":
		ignore := gitCmd.String("i", "", "Comma-separated list of ignore patterns")
		executeGitCommand(gitCmd, branch, tag, directory, ignore)
	case "local":
		localIgnore := localCmd.String("i", "", "Comma-separated list of ignore patterns")
		executeLocalCommand(localCmd, localPath, localIgnore)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: nethwv <command> [options]")
}

func executeGitCommand(gitCmd *flag.FlagSet, branch, tag, directory *string, ignore *string) {
	gitCmd.Parse(os.Args[2:])
	args := gitCmd.Args()
	if len(args) < 2 {
		fmt.Println("Usage: nethwv git [options] <user/repo> <output.pdf>")
		os.Exit(1)
	}

	repoURL := args[0]
	outputPDF := args[1]

	if err := createOutputDir(outputPDF); err != nil {
		fmt.Printf("Error creating output directory: %s\n", err)
		os.Exit(1)
	}

	client := github.NewClient(nil)

	branchOrTag := resolveBranchOrTag(branch, tag)
	if err := client.CloneRepo(repoURL, "tmp", branchOrTag); err != nil {
		fmt.Printf("Error cloning repository: %s\n", err)
		os.Exit(1)
	}

	var files []string
	var err error
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

	ignorePatterns := strings.Split(*ignore, ",")
	if err := pdf.GeneratePDF(files, outputPDF, ignorePatterns); err != nil {
		fmt.Printf("Error generating PDF: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully:", outputPDF)

	os.RemoveAll("tmp")
}

func resolveBranchOrTag(branch, tag *string) string {
	if *tag != "" {
		return *tag
	}
	return *branch
}

func executeLocalCommand(localCmd *flag.FlagSet, localPath *string, localIgnore *string) {
	localCmd.Parse(os.Args[2:])
	args := localCmd.Args()
	if len(args) < 1 {
		fmt.Println("Usage: nethwv local -p <path> <output.pdf>")
		os.Exit(1)
	}

	outputPDF := args[0]

	if *localPath == "" {
		fmt.Println("Error: Local path is required")
		os.Exit(1)
	}

	if err := createOutputDir(outputPDF); err != nil {
		fmt.Printf("Error creating output directory: %s\n", err)
		os.Exit(1)
	}

	files, err := retrieveFilesFromLocalPath(*localPath)
	if err != nil {
		fmt.Printf("Error retrieving files: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Number of files retrieved:", len(files))

	ignorePatterns := strings.Split(*localIgnore, ",")
	if err := pdf.GeneratePDF(files, outputPDF, ignorePatterns); err != nil {
		fmt.Printf("Error generating PDF: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully:", outputPDF)
}

func createOutputDir(outputPDF string) error {
	outputDir := filepath.Dir(outputPDF)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		return os.MkdirAll(outputDir, os.ModePerm)
	}
	return nil
}

func retrieveFilesFromLocalPath(localPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(localPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
