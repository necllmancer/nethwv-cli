package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (c *Client) RetrieveFiles(repoURL string) ([]string, error) {
	return c.fetchFilesRecursively(repoURL, "")
}

func (c *Client) fetchFilesRecursively(repoURL, path string) ([]string, error) {
	apiEndpoint := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repoURL, path)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}
	// HTTPリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(resp.Body)", err)
		return nil, err
	}

	var files []GitHubFile
	err = json.Unmarshal(body, &files)
	if err != nil {
		fmt.Println("json.Unmarshal(body, &files)", err, string(body))
		return nil, err
	}

	var fileUrls []string
	for _, file := range files {
		if file.Type == "file" {
			fileUrls = append(fileUrls, file.Url)
		} else if file.Type == "dir" {
			subFiles, err := c.fetchFilesRecursively(repoURL, file.Path)
			if err != nil {
				fmt.Println("c.fetchFilesRecursively(repoURL, file.Path)", err)
				return nil, err
			}
			fileUrls = append(fileUrls, subFiles...)
		}
	}

	return fileUrls, nil
}
