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
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFilterByAuthors(t *testing.T) {
	prc := PullRequestConnection{
		Edges: []pullRequestEdge{
			{
				Node: pullRequest{
					Author: author{
						Login: "Albert",
					},
				},
			},
			{
				Node: pullRequest{
					Author: author{
						Login: "Cecilia",
					},
				},
			},
			{
				Node: pullRequest{
					Author: author{
						Login: "Frieda",
					},
				},
			},
		},
		TotalCount: 3,
	}

	testCases := []struct {
		targetAuthors []string
		expected      []string
	}{
		{
			targetAuthors: []string{
				"Albert",
				"Frieda",
			},
			expected: []string{
				"Albert",
				"Frieda",
			},
		},
		{
			targetAuthors: []string{},
			expected: []string{
				"Albert",
				"Cecilia",
				"Frieda",
			},
		},
	}

	for _, testCase := range testCases {
		result := prc.FilterByAuthors(testCase.targetAuthors)
		for _, edge := range result.Edges {
			assert.Equal(t, true, edgeContainsAuthor(edge, testCase.expected))
		}
		assert.Equal(t, len(testCase.expected), result.TotalCount)
	}
}

func TestFilterByAuthor(t *testing.T) {
	target := "author"
	edges := Edges{
		{
			Node: pullRequest{
				Author: author{
					Login: target,
				},
			},
		},
		{
			Node: pullRequest{
				Author: author{
					Login: "fake",
				},
			},
		},
		{
			Node: pullRequest{
				Author: author{
					Login: target,
				},
			},
		},
	}

	result := edges.filterByAuthor(target)
	for _, edge := range *result {
		assert.Equal(t, target, edge.Node.Author.Login)
	}
}

func edgeContainsAuthor(edge pullRequestEdge, authors []string) bool {
	for _, author := range authors {
		if author == edge.Node.Author.Login {
			return true
		}
	}
	return false
}
