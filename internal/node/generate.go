package node

// generate - methods where pages/content/nodes are being created

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dojo-engineering/markdown-to-confluence/internal/confluence"
	"github.com/dojo-engineering/markdown-to-confluence/internal/flags"
	"github.com/dojo-engineering/markdown-to-confluence/internal/markdown"
	goplantuml "github.com/jfeliu007/goplantuml/parser"
	"github.com/sirupsen/logrus"
)

// generateMaster method checks if the OnlyDocs flag is true and if so then checks
// if the folder is in the /repo/docs folder,
// whether the folder is alive (has markdown files in it) and if it is, creates a page for the folder.
// Also, then checks whether the folder has subfolders in it,
// and then begins the process of checking those folders (recursively)
func (node *Node) generateMaster() {
	if flags.OnlyDocs {
		listOfFolders := strings.Split(node.path, "/")

		// we only want to read from the repo/docs folder
		// so if there is at least 1 sub folder and it is not /docs ignore
		if len(listOfFolders) > 1 && !strings.Contains(node.path, "/docs") {
			logrus.Debugf("skipping this folder [%s] because not in the /docs folder", node.path)

			return
		}
	}

	// these constants are to aid navigation of iterate method lower down
	const (
		checking   = true
		processing = false
		folders    = true
		files      = false
	)

	subNode := newNode()
	subNode.path = node.path
	subNode.root = node
	subNode.treeLink = node.treeLink
	subNode.images = node.images
	node.branches = append(node.branches, subNode)

	thereAreValidMDFiles := subNode.iterate(checking, files)
	thereAreValidImageFiles := subNode.iterate(checking, folders)

	if !thereAreValidMDFiles && !thereAreValidImageFiles {
		logrus.Debugf("no valid files here [%s]", node.path)

		return
	}

	logrus.Debugf("generating a folder page here [%s] this page is alive", node.path)

	err := node.generateFolderPage(subNode.hasIndex)
	if err != nil {
		logrus.WithError(err).Debug("Generate folder page error")

		return
	}

	node.alive = true

	subNode.generateChildPages()

	node.images = subNode.images
}

// generateChildPages method generates all children pages for all parent pages
// can be run concurrently as they all have a parent page to attach to
// so there's no need to order their generation
func (node *Node) generateChildPages() {
	const processing = false

	const files = false

	const folders = true

	wg.Add()

	go func() {
		defer wg.Done()
		node.generatePlantuml(node.path)  // generate plantuml in folders with markdown in it only
		node.iterate(processing, files)   // generate child pages for any valid files in parent page
		node.iterate(processing, folders) // attach any image files for any valid files in parent page
	}()
}

// generateFolderPage method creates a folder page in confluence for a folder
func (node *Node) generateFolderPage(hasIndex bool) error {
	nodeDirectoryName, nodeAbsolutePath := node.currentPath()
	logrus.Debugf("START processing file [%s]", nodeAbsolutePath)

	if hasIndex {
		logrus.Debugf("this location [%s] has a [%s] file so will use that as index at this location", node.path, node.indexName)

		node.indexPage = true

		pageContents, err := node.processMarkDownIndex(filepath.Join(node.path, node.indexName))
		if err != nil {
			return err
		}

		err = node.checkConfluencePages(pageContents, node.path)
		if err != nil {
			logrus.Debugf("[generate folderpage] generation error for path [%s]: %v", node.path, err)
			return err
		}
		// have to do it twice...

		pageContents, err = node.processMarkDownIndex(filepath.Join(node.path, node.indexName))
		if err != nil {
			return err
		}

		err = node.checkConfluencePages(pageContents, filepath.Join(node.path, node.indexName))
		if err != nil {
			logrus.Debugf("[generate folderpage] generation error for path [%s]: %v", node.path, err)
			return err
		}

		logrus.Debugf("processed bespoke index file - id: [%d]", node.id)

		return nil
	}

	node.indexPage = false
	logrus.Debugf("no [%s] located here [%s], will generate a generic folderpage", indexName, node.path)

	pageContents := &markdown.FileContents{
		MetaData: map[string]interface{}{
			"title": nodeDirectoryName,
		},
		Body: []byte(`<p>Welcome to the '<b>` + nodeDirectoryName + `</b>' folder of this code repo.</p>
		<p>This folder full path in the repo is: ` + nodeAbsolutePath + `</p>
<p>You will find attachments/images for this folder via the ellipsis at the top right.</p>
<p>Any markdown or subfolders is available in children pages under this page.</p>`),
	}

	err := node.checkConfluencePages(pageContents, node.path)
	if err != nil {
		logrus.Debugf("[generate folderpage] generation error for path [%s]: %v", node.path, err)
		return err
	}

	logrus.Debugf("processed generic index file - id: [%d]", node.id)

	return nil
}

// currentPath returns two strings
// string 1 - folder of the node
// string 2 - the absolute filepath to the node dir from root dir
func (node *Node) currentPath() (string, string) {
	fullDir := strings.ReplaceAll(node.path, "/github/workspace/", "")
	//TODO: Skip replacing local directory
	// fullDir = strings.ReplaceAll(fullDir, ".", "")
	// fullDir = strings.TrimPrefix(fullDir, "/")

	dirList := strings.Split(fullDir, "/")
	dir := dirList[len(dirList)-1]

	return dir, fullDir
}

// generatePlantuml takes in a folder path and
// generates a .puml file of the go code in the folder
// then calls generatePlantumlImage method to create a picture
// then creates a page for the image to be uploaded to and displayed
func (node *Node) generatePlantuml(fpath string) {
	const minimumPageSize = 20 // plantuml that is generated <= 20 chars long is too small

	const iterateThroughSubFolders = false // we want to just generate plantuml for the folder

	path, nodeAbsolutePath := node.currentPath()

	if node.root.root == nil {
		path = strings.ReplaceAll(strings.ReplaceAll(nodeAbsolutePath, ".", ""), "/", "")
	}

	logrus.Debugf("generating plantuml text for %s", path)

	result, err := goplantuml.NewClassDiagram([]string{fpath}, []string{}, iterateThroughSubFolders)
	if err != nil {
		logrus.Debugf("[generate diagram] plantuml file generation error for path [%s]: %v", nodeAbsolutePath, err)
		return
	}

	rendered := result.Render()
	if len(rendered) > minimumPageSize {
		var filename = path + "-pumldiagram"

		var buf bytes.Buffer

		var writer io.Writer

		writer, err = os.Create(node.path + "/" + filename + ".puml") //nolint:gosec // file created
		if err != nil {
			logrus.Debugf("[create file] plantuml file generation error for path [%s]: %v", nodeAbsolutePath, err)
			return
		}

		fmt.Fprint(writer, rendered)

		err := node.generatePlantumlImage(node.path + "/" + filename + ".puml")
		if err != nil {
			logrus.Debugf("generatePlantumlImage error for path [%s]: %v", nodeAbsolutePath, err)
			return
		}

		logrus.Debugf("uploading generated png file [%s] to page id of [%d]", node.path+"/"+filename+".png", node.id)

		node.uploadFile(node.path+"/"+filename+".png", node.indexPage)

		masterpagecontents := markdown.FileContents{
			MetaData: map[string]interface{}{
				"title": "plantuml-" + path + " (" + nodeAbsolutePath + ")",
			},
			Body: buf.Bytes(),
		}

		err = node.checkConfluencePages(&masterpagecontents, node.path+"/"+filename+".png")
		if err != nil {
			logrus.Debugf("check confluence page error for path [%s]: %v", nodeAbsolutePath, err)
		}

		url := flags.ConfluenceBaseURL + "/wiki/spaces/" + flags.ConfluenceSpaceName + "/pages/" + func() string {
			return strconv.Itoa(node.id)
		}()

		logrus.Debugf("a plantuml image was generated for location [%s] & is available at [%s]", fpath, url)
	}
}

// generatePlantumlImage method calls external application (plantuml.jar)
// in the docker container to generate the plantuml image (as a .png)
func (node *Node) generatePlantumlImage(fpath string) error {
	logrus.Debugf("generating plantuml png from plantuml context provided by go code...")

	convertPlantuml := exec.Command("java", "-jar", "/app/plantuml.jar", "-tpng", fpath) // #nosec - pumlimage
	convertPlantuml.Stdout = os.Stdout

	err := convertPlantuml.Run()
	if err != nil {
		return fmt.Errorf("generatePlantumlImage error: %w", err)
	}

	return nil
}

// generatePage method creates a new page for node.
// and sets the parent page as the node root id.
// unless the node.root is nil in which case it is the root page
func (node *Node) generatePage(pageContents *markdown.FileContents) error {
	_, nodeAbsolutePath := node.currentPath()

	var isParentPage = true

	if confluenceAPIClient == nil {
		return fmt.Errorf("error: confluence API client is nil")
	}

	var err error

	if node.root == nil {
		if node.parentPageID != 0 { // if the master ID is not 0 then this is a child page
			isParentPage = false
		}
		node.id, err = node.createPage(node.parentPageID, pageContents, isParentPage)
		if err != nil {
			return fmt.Errorf("create page error for folder path [%s]: %w", nodeAbsolutePath, err)
		}
		return nil
	}

	node.id, err = node.createPage(node.root.id, pageContents, !isParentPage)
	if err != nil {
		return fmt.Errorf("create page error for folder path [%s]: %w", nodeAbsolutePath, err)
	}

	return nil
}

func (node *Node) createPage(parentId int, pageContents *markdown.FileContents, isRootPage bool) (int, error) {
	parentIdString := strconv.Itoa(parentId)
	body := string(pageContents.Body)
	title, ok := pageContents.MetaData["title"].(string)
	if !ok {
		return 0, fmt.Errorf("title is empty")
	}
	createPageRequest := confluence.NewCreatePageRequest(strconv.Itoa(flags.ConfluenceSpaceID))
	createPageRequest.Status = &defaultPageStatus
	createPageRequest.Title = &title
	if !isRootPage {
		createPageRequest.ParentId = &parentIdString
	}
	createPageRequest.Body = &confluence.CreatePageRequestBody{
		PageBodyWrite: &confluence.PageBodyWrite{
			Representation: &defaultPageBodyType,
			Value:          &body,
		},
	}
	resp, _, err := confluenceAPIClient.PageApi.
		CreatePage(confluenceAuthContext).
		CreatePageRequest(*createPageRequest).
		Execute() //, node.masterID, newPageContents, isParentPage)
	if err != nil {
		return 0, fmt.Errorf("failed to create page in confluence %w", err)
	}
	return int(*resp.Id.Int64), nil
}
