// Package markdown provides a method for working with and parsing markdown documents
//
//nolint:wsl // is fine
package markdown

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dojo-engineering/markdown-to-confluence/internal/flags"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/sirupsen/logrus"
	m "gitlab.com/golang-commonmark/markdown"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// // GrabAuthors - do we want to collect authors?
// var (
// 	GrabAuthors bool
// )

// FileContents contains information from a file after being parsed from markdown.
// `Metadata` in the format of a `map[string]interface{}` this can contain title, description, slug etc.
// `Body` a `[]byte` that contains the resulting HTML after parsing the markdown and converting to HTML using Goldmark.
type FileContents struct {
	MetaData map[string]interface{}
	Body     []byte
}

// newFileContents function creates a new filecontents object
func newFileContents() *FileContents {
	f := FileContents{}
	f.MetaData = make(map[string]interface{})

	return &f
}

// ParseMarkdown function uses external parsing library to grab markdown contents
// and return a filecontents object
func ParseMarkdown(rootID int, content []byte, isIndex bool, pages map[string]string, path, nodeAbsolutePath, fileName string) (*FileContents, error) {
	r := bytes.NewReader(content)
	f := newFileContents()
	var frontMatterExists bool
	fmc, err := pageparser.ParseFrontMatterAndContent(r)
	if err != nil {
		logrus.WithError(err).Debugf("Issue with parsing frontmatter - (using # title instead)")
	} else if len(fmc.FrontMatter) != 0 {
		frontMatterExists = true
		f.MetaData = fmc.FrontMatter
	}

	// if the file name is readme.md then then the space should be named after the final folder
	if strings.EqualFold(fileName, "readme.md") {
		fileName = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	}
	var pageContent []byte
	if !frontMatterExists {
		f.MetaData["title"] = fileName
		pageContent = content
	} else {
		//Content without frontmatter
		pageContent = fmc.Content
	}

	title, ok := f.MetaData["title"]
	if !ok {
		return nil, fmt.Errorf("markdown page parsing error - page title is not assigned via toml or # section")
	}

	if title == "" {
		return nil, fmt.Errorf("markdown page parsing error - page title is empty")
	}

	md := m.New(
		m.HTML(true),
		m.Tables(true),
		m.Linkify(true),
		m.Typographer(false),
		m.XHTMLOutput(true),
	)

	preFormattedContent := md.RenderToString(pageContent)
	f.Body = stripFrontmatterReplaceURL(preFormattedContent, isIndex, pages, nodeAbsolutePath, fileName, title.(string))

	//TODO: Grab Authors from Git Commit
	// if GrabAuthors {
	// 	f.Body = append(f.Body, []byte(capGit(path))...)
	// }

	return f, nil
}

// for absolute links we don't need to do any fancy pants processing
func linkFilterLogic(item string) bool {
	if strings.Contains(item, "https://") {
		return true
	}

	if strings.Contains(item, "http://") {
		return true
	}

	if strings.Contains(item, "www") {
		return true
	}

	if strings.Contains(item, "mailto:") {
		return true
	}

	return false
}

// stripFrontmatterReplaceURL function takes in parent page ID and
// markdown file contents and removes TOML frontmatter, and replaces
// local URL with relative confluence URL
func stripFrontmatterReplaceURL(content string, isIndex bool, pages map[string]string, nodeAbsolutePath, fileName, title string) []byte {
	var pre string

	// var frontmatter bool

	lines := strings.Split(content, "\n")

	for index := range lines {
		// if strings.Contains(lines[index], "+++") {
		// 	frontmatter = flip(frontmatter)
		// 	continue
		// }

		// header lines are converted to ProperCase for confluence local linking
		htmlHeaders := []string{"<h1>", "<h2>", "<h3>", "<h4>", "<h5>", "<h6>"}

		for _, header := range htmlHeaders {
			if strings.Contains(lines[index], header) {
				lines[index] = updateHeaderToProperCase(lines[index])
			}
		}

		// correct the local url paths to be absolute paths
		if strings.Contains(lines[index], "<a href=") && !linkFilterLogic(lines[index]) {
			lines[index] = relativeURLdetector(lines[index], pages, nodeAbsolutePath, fileName, title)
		}

		// set up the local url image links
		if strings.Contains(lines[index], "<img src=") {
			lines[index] = URLConverter(pages, lines[index], isIndex, nodeAbsolutePath)
		}

		// if !frontmatter {
		pre += lines[index] + "\n"
		// }
	}

	pre = strings.TrimSpace(pre)

	return []byte(pre)
}

// updateHeaderToProperCase makes all headers be in Proper Case so local links work
func updateHeaderToProperCase(line string) string {
	splitOnLinks := strings.Split(line, `<a href="`)
	caser := cases.Title(language.English)

	if len(splitOnLinks) == 1 { // means there are no links in the header
		line = updateHeaderWithNoLinks(line, caser)
	} else { // means there are links in the header - don't want to alter the links
		line = updateHeaderWithLinks(splitOnLinks, caser)
	}

	return line
}

func updateHeaderWithNoLinks(line string, caser cases.Caser) string {
	line = caser.String(line)

	// this changes the html tags to be capitalized so reverse them back
	line = updateHTMLHeaders(line)

	// drop brackets from the title
	line = strings.ReplaceAll(line, "(", "")
	line = strings.ReplaceAll(line, ")", "")

	return line
}

func updateHeaderWithLinks(splitOnLinks []string, caser cases.Caser) string {
	line := splitOnLinks[0]
	line = caser.String(line)

	for i := 1; i < len(splitOnLinks); i++ {
		line += `<a href="` // add the '<a href="' back in

		extraParts := strings.SplitN(splitOnLinks[i], ">", 2) //nolint:gomnd // split the final part on the first >

		for ii := range extraParts {
			if ii == 0 { // add the link and the '>' back in
				line += extraParts[ii]
				line += `>`
			}

			if ii > 0 { // ProperCase the header and add it to the line
				extraParts[ii] = caser.String(extraParts[ii])
				line += extraParts[ii]
			}
		}

		// this changes the html tags to be capitalized so reverse them back
		line = updateHTMLHeaders(line)

		// drop brackets from the title
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
	}

	return line
}

// uncapitalize the 'h' in the html header tag
func updateHTMLHeaders(line string) string {
	line = strings.ReplaceAll(line, "H1>", "h1>")
	line = strings.ReplaceAll(line, "H2>", "h2>")
	line = strings.ReplaceAll(line, "H3>", "h3>")
	line = strings.ReplaceAll(line, "H4>", "h4>")
	line = strings.ReplaceAll(line, "H5>", "h5>")
	line = strings.ReplaceAll(line, "H6>", "h6>")

	return line
}

// // flip function returns the opposite of bool
// func flip(b bool) bool {
// 	return !b
// }

// takes in the absolute URL and will match the relative link to a generated confluence page
// if this fails, it will just return a template
func relativeURLdetector(line string, page map[string]string, nodeAbsolutePath, fileName, title string) string {
	const fail = `<p>[failed during relativeURLdetector]</p>`

	urlLink := strings.Split(line, `</a>`)

	originalURLslice := strings.Split(strings.ReplaceAll(urlLink[0], "<p>", ""), `>`)
	if len(originalURLslice) <= 1 {
		return fail
	}

	sliceOne := strings.Split(line, `<a href="`)
	if len(sliceOne) <= 1 {
		return fail
	}

	updatedURLs, links := generateRelativeURLs(sliceOne, nodeAbsolutePath, fileName, title)

	// replace the relative url in the item with the absolute url
	splitItem := strings.Split(line, "<a href=")

	return generateLineToReturn(updatedURLs, links, splitItem, page)
}

func generateRelativeURLs(sliceOne []string, nodeAbsolutePath, fileName, title string) ([]string, []string) {
	var (
		updatedURLs []string
		links       []string
	)

	for i := range sliceOne {
		if i == 0 {
			continue
		}

		url := strings.Split(sliceOne[i], `"`)[0] // the local/relative URL for the page

		// create the absolute url
		updatedURL, localLink := convertRelativeToAbsoluteURL(nodeAbsolutePath, url)

		// confluence is case sensitive - headers are saved using proper case (i.e. So The Title Is Always Like This)
		caser := cases.Title(language.English)
		localLink = caser.String(localLink)

		var link string

		// if there were any local links (identified by #) then create the link for confluence
		if localLink != "" {
			// as there is a local link the path must be in a .md file
			// can either be in a .md file OR a README.md (this would not have .md in the confluence page name)
			// if the updated URL does not contain a .md and the fileName does
			// then it must be a local link for the current .md file so add it
			if !strings.Contains(updatedURL, ".md") {
				if strings.Contains(fileName, ".md") {
					updatedURL += "/" + fileName
				}
			}

			fileName = strings.ReplaceAll(fileName, " ", "+")
			link = "/" + strings.ReplaceAll(title, " ", "+") + "#" + localLink
		}

		updatedURLs = append(updatedURLs, updatedURL)
		links = append(links, link)
	}

	return updatedURLs, links
}

func generateLineToReturn(updatedURL, links []string, splitItem []string, page map[string]string) string {
	stringToReturn := splitItem[0]

	for i := 1; i < len(splitItem); i++ {
		link := generateLink(page, updatedURL[i-1], links[i-1])
		extraParts := strings.SplitN(splitItem[i], ">", 2) //nolint:gomnd // only want to split the final part on the first >

		stringToReturn += link

		for ii := range extraParts {
			if ii != 0 {
				stringToReturn += extraParts[ii]
			}
		}
	}

	return stringToReturn
}

func generateLink(page map[string]string, updatedURL string, localLink string) string {
	// to format this in confluence we must follow how confluence formats its content in the web frontend
	a := `<a href="/wiki/spaces/` + flags.ConfluenceSpaceKey + `/pages/` + page[updatedURL] + localLink + `">`

	// b := `data-linked-resource-id="` + page[updatedURL] + `" `

	// c := `data-linked-resource-type="page">`

	return a //+ b + c
}

// URLConverter function for images to be loaded in to confluence page
// (they must be in same directory as markdown to work)
// this function replaces local url paths in html img links
// with a confluence path for folder page attachments on parent page
func URLConverter(page map[string]string, item string, isindex bool, nodeAbsolutePath string) string {
	sliceOne := strings.Split(item, `<img src="`)

	if len(sliceOne) > 1 {
		sliceTwo := strings.Split(sliceOne[1], `"`)

		if len(sliceTwo) > 1 {
			attachmentFileName := sliceTwo[0]

			// create the absolute url
			updatedURL, _ := convertRelativeToAbsoluteURL(nodeAbsolutePath, attachmentFileName)

			// split the path so we can rip out the file name
			splitURL := strings.Split(updatedURL, "/")
			urlWithoutFile := strings.Join(splitURL[:len(splitURL)-1], "/")

			stringToReturn := sliceOne[0]

			stringToReturn += `<span class="confluence-embedded-file-wrapped">`

			stringToReturn += `<img class="confluence-embedded-image" loading="lazy" `

			//nolint:lll /// set text
			stringToReturn += `src="` + flags.ConfluenceBaseURL + `/wiki/download/attachments/` + page[urlWithoutFile] + `/` + splitURL[len(splitURL)-1] + `" `

			//nolint:lll /// set text
			stringToReturn += `data-image-src="` + flags.ConfluenceBaseURL + `/wiki/download/attachments/` + page[urlWithoutFile] + `/` + splitURL[len(splitURL)-1] + `" `

			stringToReturn += `data-linked-resource-id="` + page[urlWithoutFile] + `" `

			stringToReturn += `data-linked-resource-type="attachment"></img></span>`

			stringToReturn += strings.Replace(sliceTwo[len(sliceTwo)-1], " />", "", 1)

			return stringToReturn
		}
	}

	return item
}

// convertRelativeToAbsoluteURL function takes in a relative url, and generates
// the correct absolute url based on the file path passed in
func convertRelativeToAbsoluteURL(nodeAbsolutePath, url string) (string, string) {
	var localLink string

	// split on #
	// length 1 means no local links
	// length 2 means local links so save everything after the #
	if len(strings.Split(url, "#")) == 2 { //nolint: gomnd // magic length 2
		localLink = strings.Split(url, "#")[1]
	}

	url = strings.ReplaceAll(url, "%20", " ")
	nodeAbsolutePath = strings.ReplaceAll(nodeAbsolutePath, "%20", " ")

	splitRelativeURL := strings.Split(url, "/")
	splitAbsoluteURL := strings.Split(nodeAbsolutePath, "/")

	var firstFolder int

	for i := range splitRelativeURL {
		switch splitRelativeURL[i] {
		case ".":
			firstFolder = i + 1
		case "..": // need to remove the final folder for each '..'
			splitAbsoluteURL = splitAbsoluteURL[:len(splitAbsoluteURL)-1]

			firstFolder = i + 1
		}
	}

	// remove any trailing # links on the relative url
	splitRelativeURL[len(splitRelativeURL)-1] = strings.Split(splitRelativeURL[len(splitRelativeURL)-1], "#")[0]

	// if the last element is blank remove it so it doesn't add a trailing slash
	if splitRelativeURL[len(splitRelativeURL)-1] == "" {
		splitRelativeURL = splitRelativeURL[:len(splitRelativeURL)-1]
	}

	// append to the end of the absolute url
	return strings.Join(append(splitAbsoluteURL, splitRelativeURL[firstFolder:]...), "/"), localLink
}

//TODO: Grab Authors from Git Commit
// type author struct {
// 	name    string
// 	howmany int
// }

// type authors []author // not using a map so the order of authors can be maintained

// func (a *authors) append(item string) {
// 	au := *a
// 	var exists bool
// 	for index := range au {
// 		if au[index].name == item {
// 			au[index].howmany++
// 			exists = true
// 			break
// 		}
// 	}

// 	if !exists {
// 		au = append(au, author{
// 			name:    item,
// 			howmany: 1,
// 		})
// 	}

// 	*a = au
// }

// func (a *authors) sort() {
// 	au := *a

// 	sort.Slice(au, func(i, j int) bool {
// 		return au[i].howmany > au[j].howmany
// 	})

// 	*a = au
// }

// // use git to capture authors by username & email & commits
// //
// //nolint:gosec // is fine
// func capGit(path string) string {
// 	here, _ := os.Getwd()
// 	log.Println("collecting authorship for ", path)
// 	git := exec.Command("git", "log", "--all", `--format='%an | %ae'`, "--", here)

// 	out, err := git.CombinedOutput()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	a := authors{}
// 	lines := strings.Split(string(out), "\n")
// 	for _, line := range lines {
// 		l := strings.ReplaceAll(line, `'`, "")
// 		a.append(l)
// 	}

// 	a.sort()

// 	// to let the output be displayed in confluence - wrapping it in code block
// 	output := `<pre><code>
// [authors | email addresses | how many commits]
// `

// 	index := 0
// 	for _, author := range a {
// 		if author.name == "" {
// 			continue
// 		}

// 		no := strconv.Itoa(author.howmany)

// 		if index == 0 {
// 			output += author.name + " - total commits: " + no
// 		} else {
// 			output += `
// ` + author.name + " - total commits: " + no
// 		}

// 		index++
// 	}

// 	output += `
// </code></pre>`

// 	return output
// }
