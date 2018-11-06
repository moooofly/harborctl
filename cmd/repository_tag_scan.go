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

// tagScanCmd represents the scan command
var tagScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the image. (NOTE: need Clair deployment)",
	Long: `Trigger jobservice to call Clair API to scan the image identified by the repo_name and tag.

NOTE: Only project admins have permission to scan images under the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanRepoTag()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/18)
var repoTagScan struct {
	repoName string
	tag      string
}

func init() {
	tagCmd.AddCommand(tagScanCmd)

	tagScanCmd.Flags().StringVarP(&repoTagScan.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	tagScanCmd.MarkFlagRequired("repo_name")

	tagScanCmd.Flags().StringVarP(&repoTagScan.tag,
		"tag",
		"t", "",
		"(REQUIRED) The name of tag.")
	tagScanCmd.MarkFlagRequired("tag")
}

func scanRepoTag() {
	targetURL := repositoryURL + "/" + repoTagScan.repoName +
		"/tags/" + repoTagScan.tag + "/scan"
	// NOTE: as per what swagger UI dose, there is no content body here
	utils.Post(targetURL, "")
}
