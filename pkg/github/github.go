package github

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type GitHubFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

// HttpClient interface defines the methods that our HTTP client must satisfy.
type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// Client struct holds the HttpClient.
type Client struct {
	HttpClient HttpClient
}

// NewClient creates a new GitHub client with the provided HTTP client.
func NewClient(httpClient HttpClient) *Client {
	return &Client{
		HttpClient: httpClient,
	}
}

func (c *Client) CloneRepo(orgAndRepository, localPath string) error {
	repoURL := fmt.Sprintf("https://github.com/%s.git", orgAndRepository)
	cmd := exec.Command("git", "clone", "--depth", "1", repoURL, localPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) RetrieveFiles(localPath string) ([]string, error) {
	return c.fetchFilesRecursively(localPath, "")
}

func (c *Client) fetchFilesRecursively(basePath, path string) ([]string, error) {
	fullPath := filepath.Join(basePath, path)
	files, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var fileFullPaths []string
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		fileFullPath := filepath.Join(basePath, filePath) // フルパスを生成
		if file.IsDir() {
			subFiles, err := c.fetchFilesRecursively(basePath, filePath)
			if err != nil {
				return nil, err
			}
			fileFullPaths = append(fileFullPaths, subFiles...)
		} else {
			fileFullPaths = append(fileFullPaths, fileFullPath)
		}
	}

	return fileFullPaths, nil
}
