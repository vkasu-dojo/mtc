package node

// delete - methods regarding deleting pages in confluence wiki

import (
	"github.com/dojo-engineering/markdown-to-confluence/internal/confluence"
	"github.com/sirupsen/logrus"
)

// findPagesToDelete method grabs results of page to begin deleting
func (node *Node) findPagesToDelete(pageId int64) {
	if confluenceAPIClient != nil {
		children, _, err := confluenceAPIClient.ChildrenApi.GetChildPages(confluenceAuthContext, pageId).Execute()
		if err != nil {
			logrus.Debugf("error finding page: %s", err)
		}

		if children != nil {
			node.deletePages(children)
		}
	}
}

// deletePages method is to find a page to delete
// and any children pages that might need to be deleted
func (node *Node) deletePages(children *confluence.MultiEntityResultChildPage) {
	for index := range children.Results {
		var noDelete bool

		for index2 := range node.titles {
			if *children.Results[index].Title == node.titles[index2] {
				noDelete = true
				break
			}
		}

		if !noDelete {
			pageId := *children.Results[index].Id.Int64
			node.findPagesToDelete(pageId)

			go node.deletePage(pageId)
		}
	}
}

// deletePage method converts id to integer to pass to the API method DeletePage
// this method can be run concurrently with no wait needed:
// the pages are deleted by ID and don't need parent page reference
func (node *Node) deletePage(pageId int64) {
	_, err := confluenceAPIClient.PageApi.DeletePage(confluenceAuthContext, pageId).Execute()
	if err != nil {
		logrus.Debugf("error deleting page: %s", err)
	}
}
