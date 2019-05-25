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
	"context"
	"errors"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"strings"
)

// Repository is the structure of repository.
type Repository struct {
	Owner string
	Name  string
}

// New create the repository structure from repository path.
func New(repositoryPath string) (*Repository, error) {
	splits := strings.FieldsFunc(string(repositoryPath), func(r rune) bool {
		return r == '/'
	})

	if len(splits) != 2 {
		msg := fmt.Sprintf("%v is wrong repository path", repositoryPath)
		return nil, errors.New(msg)
	}

	return &Repository{
		Owner: splits[0],
		Name:  splits[1],
	}, nil
}

// ExecuteQuery execute GraphQL query to Github.
func (r Repository) ExecuteQuery(token string) (*Query, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	variables := map[string]interface{}{
		"owner": githubv4.String(r.Owner),
		"name":  githubv4.String(r.Name),
	}

	var q Query
	err := client.Query(context.Background(), &q, variables)
	if err != nil {
		return nil, err
	}
	return &q, nil
}