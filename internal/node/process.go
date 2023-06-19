package node

// methods for processing/reading/uploading files & iterating through folders

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dojo-engineering/markdown-to-confluence/internal/flags"
	"github.com/dojo-engineering/markdown-to-confluence/internal/markdown"
	"github.com/dojo-engineering/markdown-to-confluence/internal/todo"
	"github.com/sirupsen/logrus"
)

// processGoFile method takes in a go file contents and
// calls method todo.ParseGo on the file contents with the
// file path
func (node *Node) processGoFile(fpath string) error {
	_, abs := node.currentPath()

	contents, err := os.ReadFile(filepath.Clean(fpath))
	if err != nil {
		return fmt.Errorf("absolute path [%s] - file [%s] - read file error: %w",
			abs, fpath, err)
	}

	fullpath := strings.Replace(fpath, ".", "", 2) //nolint:gomnd // only want to replace max of first 2

	fullpath = strings.TrimPrefix(fullpath, "/")

	todo.ParseGo(contents, fullpath) // not used atm but will be maybe in future

	return nil
}

// processMarkDownIndex method takes in index file contents
// and parses the markdown file
func (node *Node) processMarkDownIndex(path string) (*markdown.FileContents, error) {
	_, nodeAbsolutePath := node.currentPath()

	contents, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("absolute path [%s] - file [%s] - read file error: %w", nodeAbsolutePath, path, err)
	}

	mapSem <- struct{}{}

	parsedContents, err := markdown.ParseMarkdown(func() int {
		if node.root == nil {
			return 0
		}

		return node.root.id
	}(), contents, node.indexPage, node.treeLink.branches, node.path, nodeAbsolutePath, node.indexName)
	if err != nil {
		<-mapSem

		return nil, fmt.Errorf("absolute path [%s] - file [%s] - parse markdown error: %w", nodeAbsolutePath, path, err)
	}

	<-mapSem

	parsedContents.MetaData["title"] = parsedContents.MetaData["title"].(string) //TODO: + " (" + nodeAbsolutePath + ")"

	return parsedContents, nil
}

// processMarkDown method takes in file contents
// and parses the markdown file before calling
// checkConfluencePages method
func (node *Node) processMarkDown(path, fileName string) error {
	_, abs := node.currentPath()

	contents, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("absolute path [%s] - file [%s] - read file error: %w",
			abs, path, err)
	}

	mapSem <- struct{}{}

	parsedContents, err := markdown.ParseMarkdown(func() int {
		if node.root == nil {
			return 0
		}

		return node.root.id
	}(), contents, node.indexPage,
		node.treeLink.branches, node.path, abs, fileName)
	if err != nil {
		<-mapSem
		return fmt.Errorf("absolute path [%s] - file [%s] - parse markdown error: %w",
			abs, path, err)
	}

	<-mapSem

	//parsedContents.MetaData["title"] = parsedContents.MetaData["title"].(string) //TODO: + " (" + abs + ")"

	err = node.checkConfluencePages(parsedContents, path)
	if err != nil {
		return fmt.Errorf("absolute path [%s] - file [%s] - confluence check error: %w",
			abs, path, err)
	}

	return nil
}

// uploadFile method takes in file and
// uploads the file to a page by parent page ID (node.root.id)
func (node *Node) uploadFile(path string, isIndexPage bool) {
	_, abs := node.currentPath()

	var pageId int

	if isIndexPage {
		pageId = node.id
	} else {
		pageId = node.root.id
	}

	err := node.UploadAttachment(filepath.Clean(path), pageId)
	if err != nil {
		logrus.Debugf("absolute path [%s] - local path [%s] - file upload error: %v",
			path, abs, err)
	}
}

func newfileUploadRequest(uri string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("file close error: %w", err))
		}
	}()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, fmt.Errorf("create form file error: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("io copy error: %w", err)
	}

	params := map[string]string{
		"comment":   "file uploaded using markdown-github-action",
		"minorEdit": "true",
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	req, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return nil, fmt.Errorf("writer close error: %w", err)
	}

	data := []byte{}

	_, err = file.Read(data)
	if err != nil {
		log.Println(err)
	}

	sEnc := base64.StdEncoding.EncodeToString(data)
	req.ContentLength = int64(len(sEnc))

	req.Header = http.Header{
		"Content-Type": []string{writer.FormDataContentType()},
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("writer close error: %w", err)
	}

	return req, err
}

// UploadAttachment to a page identify by page ID
// you need the page ID to upload the attachment(file path)
func (node *Node) UploadAttachment(filename string, id int) error {
	targetURL := fmt.Sprintf(flags.ConfluenceBaseURL+"/wiki/rest/api/content/%d/child/attachment", id)

	req, err := newfileUploadRequest(targetURL, "file", filename)
	if err != nil {
		return fmt.Errorf("file upload error: %w", err)
	}

	req.SetBasicAuth(flags.ConfluenceUsername, flags.ConfluenceAPIKey)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("upload attachment response error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload attachment response issue: %s", resp.Status)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(fmt.Errorf("body close error: %w", err))
		}
	}()

	return nil
}
