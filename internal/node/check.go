package node

// check - methods for checking various conditions

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dojo-engineering/markdown-to-confluence/internal/confluence"
	"github.com/dojo-engineering/markdown-to-confluence/internal/markdown"
	"github.com/sirupsen/logrus"
)

// checkIfRootAlive method checks if root node is alive,
// and creates a subnode (alive = markdown files exist in folder)
// if root is alive then root is the parent node for subnode
// else the root root node is the parent node for subnode
// then it calls generateMaster method on subnode
func (node *Node) checkIfRootAlive(fpath string) {
	if node.path != fpath {
		subNode := newNode()
		subNode.treeLink = node.treeLink
		subNode.path = fpath
		subNode.images = node.images

		if node.alive {
			subNode.root = node.root
		} else {
			if node.root != nil {
				subNode.root = node.root.root
			} else {
				subNode.root = node.root
			}
		}

		node.branches = append(node.branches, subNode)

		subNode.generateMaster()
	}
}

// fileInDirectoryCheck method takes file path and two bools (checking / folders)
//
// if folders is false & checking is true then it returns true if it finds a markdown file in a folder
//
// if folders is false & checking is false then it processes markdown files via checkIfMarkDown method
//
// if folders is true & checking is true then it returns true if it finds an image file in a folder
//
// if folders is true and checking is false then it processes uploads images and processes
// other file types via checkOtherFileTypes method
func (node *Node) fileInDirectoryCheck(fpath string, checking, folders bool) bool {
	if folders {
		return node.checkOtherFileTypes(fpath, checking) // you can process other file types inside this method
	}

	validFile := node.checkIfMarkDown(fpath, checking) // for checking & processing markdown files

	return validFile && checking
}

// checkIfMarkDown method checks is a folder or not, and if not
// passes file to checkIfMarkDownFile method
func (node *Node) checkIfMarkDown(fpath string, checking bool) bool {
	if !isFolder(fpath) {
		if ok := node.checkIfMarkDownFile(checking, fpath); ok {
			if checking {
				node.alive = true
			}

			return true
		}
	}

	return false
}

// checkIfMarkDownFile method checks whether file is a markdown file or not
// checking bool is for whether we are just checking returning bool, or
// if we are doing work on file
func (node *Node) checkIfMarkDownFile(checking bool, name string) bool {
	fileName := filepath.Base(name)

	if strings.HasSuffix(strings.ToLower(fileName), ".md") {
		if !checking {
			if strings.ToLower(fileName) == indexName { // we don't want to process index.md here
				return true
			}

			err := node.processMarkDown(name, fileName)
			if err != nil {
				logrus.WithError(err).Debug("Failed to process markdown")
			}

			return true
		}

		return true
	}

	return false
}

// checkIfFolder method checks filepath is a folder or not
// and returns bool
func (node *Node) checkIfFolder(fpath string) bool {
	if isFolder(fpath) && !isVendorOrGit(fpath) {
		node.checkIfRootAlive(fpath)

		return true
	}

	return false
}

// checkOtherFileTypes method checks if file is a folder
// and if not, checks for if it is a go or image file
func (node *Node) checkOtherFileTypes(fpath string, checking bool) bool {
	if !node.checkIfFolder(fpath) {
		node.checkIfGoFile(fpath)
		return node.checkForImages(fpath, checking)
	}

	return false
}

// checkIfGoFile method checks to see if the file is
// a golang file
func (node *Node) checkIfGoFile(name string) {
	if !strings.HasSuffix(name, ".go") {
		return
	}

	err := node.processGoFile(name)
	if err != nil {
		logrus.WithError(err).Debug("Failed to process go file")
	}
}

// checkForImages method checks to see if the file is an image file
func (node *Node) checkForImages(name string, checking bool) bool {
	validFiles := []string{".png", ".jpg", ".jpeg", ".gif", ".svg"}

	for index := range validFiles {
		if strings.HasSuffix(strings.ToLower(name), validFiles[index]) {
			if checking {
				node.alive = true
			} else {
				node.checkNodeRootIsNil(name)
			}

			return true
		}
	}

	return false
}

// checkNodeRootIsNil method checks whether the
// node root is nil before calling uploadFile method
func (node *Node) checkNodeRootIsNil(name string) {
	if node.root != nil {
		if node.imageToBeUploaded(name) {
			node.uploadFile(name, node.indexPage)
		}
	}
}

// imageToBeUploaded method checks if this imag ewas previously uploaded
func (node *Node) imageToBeUploaded(name string) bool {
	if _, exists := node.images[name]; exists {
		return false
	}

	node.images[name] = true

	return true
}

// checkConfluencePages method runs through the CRUD operations for confluence
func (node *Node) checkConfluencePages(pageContents *markdown.FileContents, filepath string) error {
	_, nodeAbsolutePath := node.currentPath()

	if pageContents == nil {
		return fmt.Errorf("checkConfluencePages error for folder path [%s]: the markdown file was nil", nodeAbsolutePath)
	}

	pageTitle := pageContents.MetaData["title"].(string)

	var parentPageId int64
	if node.root != nil {
		parentPageId = int64(node.root.id)
	} else {
		parentPageId = int64(node.parentPageID)
	}

	pages, r, err := confluenceAPIClient.ChildrenApi.GetChildPages(confluenceAuthContext, parentPageId).Execute()
	if err != nil && r.StatusCode != http.StatusNotFound {
		return err
	}
	var page *confluence.ChildPage
	if r.StatusCode == http.StatusOK {
		for _, p := range pages.Results {
			p := p
			if strings.EqualFold(pageTitle, *p.Title) {
				page = &p
			}
		}
	}

	if page == nil {
		err := node.newPage(pageContents)
		if err != nil {
			return fmt.Errorf("create page error for folder path [%s] - page title [%s]: %w", nodeAbsolutePath, pageTitle, err)
		}

		mapSem <- struct{}{}

		id := strconv.Itoa(node.id)

		node.treeLink.branches[filepath] = id

		logrus.Debugf("processed file - id: [%d]", node.id)
		<-mapSem

		return nil
	}

	err = node.updatePage(pageContents, page)
	if err != nil {
		return fmt.Errorf("create/update page error for folder path [%s] - page title [%s]: %w", nodeAbsolutePath, pageTitle, err)
	}

	mapSem <- struct{}{}

	id := strconv.Itoa(node.id)

	node.treeLink.branches[filepath] = id

	logrus.Debugf("processed file - id: [%d]", node.id)
	<-mapSem

	return nil
}

func (node *Node) newPage(pageContents *markdown.FileContents) error {
	_, nodeAbsolutePath := node.currentPath()

	if pageContents == nil {
		return fmt.Errorf("newPage error for folder path [%s]: the markdown file was nil", nodeAbsolutePath)
	}

	err := node.generatePage(pageContents)
	if err != nil {
		return err
	}

	node.addContents(pageContents)

	return nil
}

func (node *Node) updatePage(pageContents *markdown.FileContents, page *confluence.ChildPage) error {
	_, nodeAbsolutePath := node.currentPath()

	if pageContents == nil {
		return fmt.Errorf("createOrUpdatePage error for folder path [%s]: the newPageContents param was nil", nodeAbsolutePath)
	}

	if page == nil {
		return fmt.Errorf("createOrUpdatePage pageResult error for folder path [%s]: the pageResult param was nil", nodeAbsolutePath)
	}

	node.checkPageID(*page)

	existingPage, _, err := confluenceAPIClient.PageApi.GetPageById(confluenceAuthContext, *page.Id.Int64).BodyFormat(confluence.PRIMARYBODYREPRESENTATION_STORAGE).Execute()
	if err != nil {
		return err
	}
	pageBody := string(pageContents.Body)
	if *existingPage.Body.Storage.Value == pageBody {
		logrus.Debug("No changes to file %s", nodeAbsolutePath)
		return nil
	}

	versions, _, err := confluenceAPIClient.VersionApi.GetPageVersions(confluenceAuthContext, *page.Id.Int64).Sort(confluence.VERSIONSORTORDER_MODIFIED_DATE_DESC).Limit(1).Execute()
	if err != nil {
		return err
	}
	var version int32
	if len(versions.Results) > 0 {
		version = *versions.Results[0].Number + 1
	}
	updatePageRequest := &confluence.UpdatePageRequest{
		Id:      strconv.FormatInt(*page.Id.Int64, 10),
		Status:  string(*page.Status),
		Title:   *page.Title,
		SpaceId: strconv.FormatInt(*page.SpaceId.Int64, 10),
		Version: confluence.UpdateBlogPostRequestVersion{
			Number: &version,
		},
		Body: confluence.CreatePageRequestBody{
			PageBodyWrite: &confluence.PageBodyWrite{
				Representation: &defaultPageBodyType,
				Value:          &pageBody,
			},
		},
	}

	updatedPage, _, err := confluenceAPIClient.PageApi.UpdatePage(confluenceAuthContext, *page.Id.Int64).UpdatePageRequest(*updatePageRequest).Execute()
	if err != nil {
		return err
	}

	if updatedPage != nil {
		node.addContents(pageContents)
	}

	return nil
}

// addContents adds the page title to either the parent page titles slice, or the node slice
// multiple goroutines could access same titles (or node.root.titles) slice so locking is required
func (node *Node) addContents(pageContents *markdown.FileContents) {
	_, nodeAbsolutePath := node.currentPath()

	if pageContents.MetaData == nil {
		logrus.Debugf("createOrUpdatePage warning for folder path [%s]: metadata was nil", nodeAbsolutePath)

		return
	}

	if node.root != nil {
		node.root.mu.Lock()

		defer node.root.mu.Unlock()

		node.root.titles = append(node.root.titles, pageContents.MetaData["title"].(string))
	}

	node.mu.Lock()

	defer node.mu.Unlock()

	node.titles = append(node.titles, pageContents.MetaData["title"].(string))
}

// checkPageID method checks the pageID of the page contents and
// sets the node id to the page id
// multiple goroutines could access same id field so locking is required
func (node *Node) checkPageID(page confluence.ChildPage) {
	node.mu.RLock()

	defer node.mu.RUnlock()

	node.id = int(*page.Id.Int64)
}
