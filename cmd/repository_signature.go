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

// signatureCmd represents the signature command
var signatureCmd = &cobra.Command{
	Use:   "signature",
	Short: "Get signature information of a repository from notary instance.",
	Long: `This endpoint aims to retrieve signature information of a repository, the data is from the nested notary instance of Harbor.

If the repository does not have any signature information in notary, this API will return an empty list with response code 200, instead of 404`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepoSignature()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/15)
var repoSignature struct {
	repoName string
}

func init() {
	repositoryCmd.AddCommand(signatureCmd)

	signatureCmd.Flags().StringVarP(&repoSignature.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository for which to get signature.")
	signatureCmd.MarkFlagRequired("repo_name")
}

func getRepoSignature() {
	targetURL := repositoryURL + "/" + repoSignature.repoName + "/signatures"
	utils.Get(targetURL)
}
