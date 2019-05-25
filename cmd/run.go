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

package cmd

import (
	"fmt"
	"github.com/go-go-megaman/prchecker/repository"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Display list of pull requests.",
	Long:  `Display list of pull requests from specific repositories.`,
	Run:   run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(_ *cobra.Command, _ []string) {
	for _, repositoryPath := range config.Repositories {
		r, err := repository.New(repositoryPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		q, err := r.ExecuteQuery(config.Token)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Repository:", repositoryPath)

		q.Repository.
			PullRequests.
			FilterByAuthors(config.Authors).
			PrintContents()
	}
}
