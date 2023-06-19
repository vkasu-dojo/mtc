package node

// helper & factory functions

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

// newNode function creates a new node object
func newNode() *Node {
	node := Node{}

	node.mu = &sync.RWMutex{}

	return &node
}

// withinDirectory function checks to see if the file (base) is within the folder (path)
func (node *Node) withinDirectory(base, path string) bool {
	return strings.Count(path, "/")-strings.Count(base, "/") == 1
}

// isVendorOrGit function takes in name of folder and
// checks if it is a vendor or github folder
func isVendorOrGit(name string) bool {
	if strings.Contains(name, "vendor") || strings.Contains(name, ".github") || strings.Contains(name, ".git") {
		return true
	}

	return false
}

// isFolder function checks whether a file is a folder or not
func isFolder(name string) bool {
	file, err := os.Open(filepath.Clean(name))
	if err != nil {
		logrus.WithError(err).Debugf("Failed to open folder %s", name)
		return false
	}

	defer func() {
		err := file.Close()
		if err != nil {
			logrus.WithError(err).Debugf("Failed to close file %s", name)
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		logrus.WithError(err).Debugf("Failed to get file stats %s", name)
		return false
	}

	return fileInfo.IsDir()
}
