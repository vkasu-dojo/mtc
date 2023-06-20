package markdown

import (
	"testing"

	"github.com/dojo-engineering/markdown-to-confluence/internal/flags"
	"github.com/stretchr/testify/assert"
)

func TestRelativeURLdetector(t *testing.T) {
	testInputs := []struct {
		name           string
		arg1           string
		arg2           map[string]string
		arg3           string
		filename       string
		title          string
		expectedoutput string
	}{
		{
			name: "local most likely",
			arg1: `<p><a href="../hello/there">a_page</a></p>`,
			arg2: map[string]string{
				"/absolute/path/wrong/node": "1",
				"/absolute/path/node":       "2",
				"/absolute/path":            "3",
				"/absolute/hello/there":     "4",
			},
			arg3:     "/absolute/path",
			filename: "file",
			title:    "file",
			//nolint:lll /// test data
			expectedoutput: `<p><a href="/wiki/spaces//pages/4">a_page</a></p>`,
		},
		{
			name: "distant markdown link most likely",
			arg1: `<p><a href="./node">a_page</a></p>`,
			arg2: map[string]string{
				"/absolute/path/wrong/node": "1",
				"/absolute/path/node":       "2",
				"/absolute/path":            "3",
			},
			arg3:     "/absolute/path",
			filename: "file",
			title:    "file",
			//nolint:lll /// test data
			expectedoutput: `<p><a href="/wiki/spaces//pages/2">a_page</a></p>`,
		},
		{
			name: "relative link in same file",
			arg1: `<p><a href="./#test">a_page</a></p>`,
			arg2: map[string]string{
				"/absolute/path/wrong/node": "1",
				"/absolute/path/node":       "2",
				"/absolute/path":            "3",
				"/absolute/path/file.md":    "4",
			},
			arg3:     "/absolute/path",
			filename: "file.md",
			title:    "file",
			//nolint:lll /// test data
			expectedoutput: `<p><a href="/wiki/spaces//pages/4/file#Test">a_page</a></p>`,
		},
	}

	for _, test := range testInputs {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedoutput, relativeURLdetector(test.arg1, test.arg2, test.arg3, test.filename, test.title))
		})
	}
}

func TestParseMarkDown(t *testing.T) {
	link := flags.ConfluenceAPIBaseURL + `/wiki/download/attachments//node.png`
	testInputs := []struct {
		Name     string
		input    []byte
		expected *FileContents
	}{
		{
			Name: "title & no URL",
			input: []byte(`# Markdown to Confluence Action

This Action will trawl through a repository.`),
			expected: &FileContents{
				MetaData: map[string]interface{}{
					"title": "filename",
				},
				Body: []byte(`<h1>Markdown To Confluence Action</h1>
<p>This Action will trawl through a repository.</p>`),
			},
		},
		{
			Name: "title & URL",
			input: []byte(`# Markdown to Confluence Action

![Diagram of action methodology](node.png)`),
			expected: &FileContents{
				MetaData: map[string]interface{}{
					"title": "filename",
				},
				//nolint:lll /// test data
				Body: []byte(`<h1>Markdown To Confluence Action</h1>
<p><span class="confluence-embedded-file-wrapped"><img class="confluence-embedded-image" loading="lazy" src="` + link + `" data-image-src="` + link + `" data-linked-resource-id="" data-linked-resource-type="attachment"></img></span></p>`),
			},
		},
	}

	for _, test := range testInputs {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			result, _ := ParseMarkdown(0, test.input, false, map[string]string{}, ".", "/abs/path", "filename")
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestParseMarkdown_HappyPath(t *testing.T) {
	testContent := []byte(`
---
categories: ["Development", "Github Actions"]
date: "2021-03-10"
description: "A guide on how to use the markdown to confluence action"
slug: "markdown-to-confluence-guide"
title: "Markdown to Confluence Action Guide"
---

# Test Content
test description`)

	expectOutput := &FileContents{
		MetaData: map[string]interface{}{
			"categories":  []interface{}{"Development", "Github Actions"},
			"date":        "2021-03-10",
			"description": "A guide on how to use the markdown to confluence action",
			"slug":        "markdown-to-confluence-guide",
			"title":       "Markdown to Confluence Action Guide",
		},
		Body: []byte(`<h1>Test Content</h1>
<p>test description</p>`),
	}

	out, err := ParseMarkdown(0, testContent, false, map[string]string{}, ".", "/abs/path", "filename")
	assert.Nil(t, err)
	assert.Equal(t, out, expectOutput)
}
