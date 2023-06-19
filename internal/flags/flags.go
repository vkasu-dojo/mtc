// Package common is for storing common constants/vars used in app
package flags

var (
	// ConfluenceBaseURL is the base URL for the confluence page you want the API to connect to
	ConfluenceBaseURL string

	// ConfluenceUsername is to collect external arg for confluence username
	ConfluenceUsername string

	// ConfluenceAPIKey is to collect external arg for api key
	ConfluenceAPIKey string

	// ConfluenceSpaceName is to collect external arg for confluence space
	ConfluenceSpaceName string

	// ConfluenceSpaceID is to collect external arg for confluence space
	ConfluenceSpaceID int

	// SourceDocsPath is to collect external arg for project path
	SourceDocsPath string

	// RootPageID is to collect external arg for the correct parent page ID
	RootPageID int

	// OnlyDocs is a flag to decide whether it is only the /docs folder to copy across
	OnlyDocs bool
)
