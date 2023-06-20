// Package common is for storing common constants/vars used in app
package flags

var (
	// ConfluenceAPIBaseURL is the base URL for the confluence page you want the API to connect to
	ConfluenceAPIBaseURL string

	// ConfluenceBaseURL is the base URL for the confluence page
	ConfluenceBaseURL string

	// ConfluenceUsername is to collect external arg for confluence username
	ConfluenceUsername string

	// ConfluenceAPIKey is to collect external arg for api key
	ConfluenceAPIKey string

	// ConfluenceSpaceKey is to collect external arg for confluence space
	ConfluenceSpaceKey string

	// ConfluenceSpaceID is to collect external arg for confluence space
	ConfluenceSpaceID int64

	// SourceDocsPath is to collect external arg for project path
	SourceDocsPath string

	// RootPageID is to collect external arg for the correct parent page ID
	RootPageID int

	// OnlyDocs is a flag to decide whether it is only the /docs folder to copy across
	OnlyDocs bool
)
