// Copyright © 2018 moooofly <centos.sf@gmail.com>
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
	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

var searchURL string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for projects and repositories.",
	Long: `The Search endpoint returns information about the projects and repositories offered at public status
or related to the current logged in user. The response includes the project and repository list in a proper display order.

NOTE: This endpoint can be used without cookie.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		searchURL = utils.URLGen("/api/search")
	},
	Run: func(cmd *cobra.Command, args []string) {
		searchAll()
	},
}

var search struct {
	query string
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&search.query, "query", "q", "", "Search parameter for project and repository name. If not set, search all.")
}

// searchAll returns information about the projects and repositories offered at public status or related to the current logged in user.
func searchAll() {
	targetURL := searchURL + "?q=" + search.query
	utils.Get(targetURL)

	// NOTE:
	// As experiment shows, "/api/search" can be used without cookie setting,
	// which can be verified by "offered at public status or related to the current logged in user" from doc.
	// The cookie setting above is redundant.
}
