// Copyright © 2019 megaman-go-go
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package repository

import (
	"fmt"
	"time"
)

// Query is the structure of GraphQL query.
type Query struct {
	Repository struct {
		PullRequests PullRequestConnection `graphql:"pullRequests(last:20 states:OPEN orderBy:{direction:ASC field: CREATED_AT})"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

// PullRequestConnection is the structure of pull request connection.
type PullRequestConnection struct {
	Edges      Edges
	TotalCount int
}

// Edges is the structure of pull request edge.
type Edges []pullRequestEdge

type pullRequestEdge struct {
	Node pullRequest
}

type pullRequest struct {
	Title     string
	Author    author
	URL       string
	CreatedAt time.Time
}

type author struct {
	Login string
}

// PrintContents print pull request contents.
func (prc PullRequestConnection) PrintContents() {
	fmt.Println("Total count:", prc.TotalCount)
	fmt.Print("\n")

	for _, edge := range prc.Edges {
		fmt.Println("Title:", edge.Node.Title)
		fmt.Println("Author:", edge.Node.Author.Login)
		fmt.Println("Created at:", edge.Node.CreatedAt)
		fmt.Println("URL:", edge.Node.URL)
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

// FilterByAuthors filter pull request connection by author.
func (prc PullRequestConnection) FilterByAuthors(authors []string) *PullRequestConnection {
	if len(authors) == 0 {
		return &prc
	}

	var result Edges
	for _, author := range authors {
		result = append(result, *prc.Edges.filterByAuthor(author)...)
	}

	return &PullRequestConnection{
		Edges:      result,
		TotalCount: len(result),
	}
}

func (e Edges) filterByAuthor(author string) *Edges {
	var result Edges
	for _, edge := range e {
		if edge.Node.Author.Login == author {
			result = append(result, edge)
		}
	}

	return &result
}
