// Package node is to enable reading through a repo and create a tree of content on confluence
package node

//notodo: no need
import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dojo-engineering/markdown-to-confluence/internal/confluence"
	"github.com/dojo-engineering/markdown-to-confluence/internal/flags"
	"github.com/dojo-engineering/markdown-to-confluence/internal/semaphore"
	"github.com/sirupsen/logrus"
)

const (
	numberOfRoutines = 10 // limit number of goroutines (to balance load on confluence API)
	indexName        = "readme.md"
)

var (
	mapSem                = make(chan struct{}, 1)                   // for controlling access to Tree map
	wg                    = semaphore.NewSemaphore(numberOfRoutines) // for controlling number of goroutines
	confluenceAPIClient   *confluence.APIClient                      // api client will be stored here
	confluenceAuthContext context.Context
	t                     *Tree
	defaultPageBodyType   string = "storage"
	defaultPageStatus     string = "current"
)

// Tree - capture what has been generated
type Tree struct {
	branches map[string]string
}

// Node struct enables creation of a page tree
type Node struct {
	treeLink     *Tree
	parentPageID int    // the confluence page ID for the parent page the mtc tool should create the files in
	id           int    // when page is created, page ID will be stored here.
	alive        bool   // for tracking if the folder has any valid content within it asides more folders
	path         string // file / folderpath will be stored here
	hasIndex     bool
	root         *Node           // the parent page node will be linked here
	branches     []*Node         // any children page nodes will be stored here (for deleting)
	titles       []string        // titles of pages created by node (for deleting)
	images       map[string]bool // image files uploaded to prevent uploading them multiple times
	mu           *sync.RWMutex   // for locking/unlocking when multiple goroutines are working on same node
	indexPage    bool
	indexName    string
}

func RootNode(client *confluence.APIClient) *Node {
	confluenceAPIClient = client
	confluenceAuthContext = context.WithValue(context.Background(), confluence.ContextBasicAuth, confluence.BasicAuth{
		UserName: flags.ConfluenceUsername,
		Password: flags.ConfluenceAPIKey,
	})
	return &Node{}
}

// Start method begins the generation of a tree of the repo for confluence
// first it validates whether the project path is a folder
// if yes then it sets the rootDir as the project path folder name
// then begins the recursive method generateMaster
// and returns bool - if true then it means pages have been created/updated/checked on confluence
// and there is markdown content in the folder
func (node *Node) Start(rootPageId int, sourceDocsPath string, onlyDocs bool) bool {
	spaces, _, err := confluenceAPIClient.SpaceApi.GetSpaces(confluenceAuthContext).Keys([]string{flags.ConfluenceSpaceKey}).Execute()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to fetch space id from space key")
	}
	for _, s := range spaces.Results {
		s := s
		if s.Key == &flags.ConfluenceSpaceKey {
			flags.ConfluenceSpaceID = *s.Id.Int64
			break
		}
	}
	if flags.ConfluenceSpaceID == 0 {
		logrus.Fatal("Failed to set confluence space id")
	}

	logrus.Infof("Docs sync started from %s directory to confluence url %s and space %s", flags.SourceDocsPath, flags.ConfluenceBaseURL, flags.ConfluenceSpaceKey)

	if t == nil {
		logrus.Debug("Instantiating Tree")

		t = func() *Tree { // t - Tree - will contain tree of pages created and their subsequent confluence URL
			return &Tree{
				branches: make(map[string]string),
			}
		}()
	}

	node.treeLink = t

	node.mu = &sync.RWMutex{}

	node.images = map[string]bool{}

	/*
		FOR RELATIVE FILE LINKS IN CONFLUENCE...

		Start needs to be called twice because confluence page ID's are captured _only_ when they are generated
		successfully. So, if there are local links inside pages, they cannot be determined until after the page
		is created - bizarre i know - so the logic has to be run twice so that we can first:

		- generate a tree of pages with their confluence ID's
		- re-generate a tree of pages with the confluence ID's known
	*/

	if isFolder(sourceDocsPath) {
		node.parentPageID = rootPageId

		node.path = sourceDocsPath

		dirEntries, err := os.ReadDir(sourceDocsPath)
		if err != nil {
			logrus.WithError(err).Debugf("Failed to traverse source directory")
		}
		var hasIndex bool = false
		for _, e := range dirEntries {
			fileInfo, err := e.Info()
			if err != nil {
				logrus.WithError(err).Debugf("Failed to get info for directory entry %s", e.Name())
				continue
			}
			if !fileInfo.IsDir() && strings.ToLower(fileInfo.Name()) == indexName {
				hasIndex = true
				node.indexName = indexName
				break
			}
		}

		err = node.generateFolderPage(hasIndex) // create the main page first
		if err != nil {
			return false
		}

		node.generateMaster() // contains concurrency

		logrus.Debug("WAITING FOR GOROUTINES [1] - generating first run to capture page ID's")

		wg.Wait()

		node.generateMaster() // contains concurrency

		logrus.Debug("WAITING FOR GOROUTINES [2] - now trying to match relative links to page ID's")

		wg.Wait()

		logrus.Debug("Here are the pages I got today:")

		node.Tree()

		logrus.Debug("FINISHED GOROUTINES - NOW CHECKING FOR DELETE")

		return true
	}

	return false
}

// Tree - print out what has been generated
func (node *Node) Tree() {
	for path, id := range t.branches {
		logrus.Debugf("%s | %s/wiki/spaces/%s/pages/%s", path, flags.ConfluenceBaseURL, flags.ConfluenceSpaceKey, id)
	}
}

// iterate method is to scan through the files or folders in a folder.
// and takes in two bools (justChecking, foldersOnly)
// if justChecking is true then it will only check whether
// there is a valid file in folder and return true if there is
// if foldersOnly is true then it will only iterate through folders
func (node *Node) iterate(justChecking, foldersOnly bool) bool {
	var thereIsAValidFile bool

	err := filepath.Walk(node.path, func(fpath string, info os.FileInfo, err error) error {
		if node.withinDirectory(node.path, fpath) {
			if strings.ToLower(filepath.Base(fpath)) == indexName {
				node.hasIndex = true
				node.alive = true
				if node.root != nil {
					node.root.indexName = filepath.Base(fpath)
				} else {
					node.indexName = filepath.Base(fpath)
				}
			}

			validFile := node.fileInDirectoryCheck(fpath, justChecking, foldersOnly)
			if validFile {
				thereIsAValidFile = true
			}
		}

		return nil
	})
	if err != nil {
		logrus.Debug("iterate: ", err)
	}

	return thereIsAValidFile
}

// Delete method starts loop through node.branches
// and calls this method on each subnode of the node
// if node.id != 0 (i.e not the root node) then
// it calls method findPagesToDelete
func (node *Node) Delete() {
	if node.id != 0 {
		node.findPagesToDelete(int64(node.id))
	}

	for index := range node.branches {
		node.branches[index].Delete()
	}
}
