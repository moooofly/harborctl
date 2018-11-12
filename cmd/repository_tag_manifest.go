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

// manifestCmd represents the manifest command
var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Get manifests of a relevant repository.",
	Long: `This endpoint aims to retrieve manifests from a relevant repository.

NOTE: This endpoint can be used without cookie.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepoTagManifest()
	},
}

var repoTagManifestGet struct {
	repoName string
	tag      string
	version  string
}

func init() {
	tagCmd.AddCommand(manifestCmd)

	manifestCmd.Flags().StringVarP(&repoTagManifestGet.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	manifestCmd.MarkFlagRequired("repo_name")

	manifestCmd.Flags().StringVarP(&repoTagManifestGet.tag,
		"tag",
		"t", "",
		"(REQUIRED) The name of tag.")
	manifestCmd.MarkFlagRequired("tag")

	manifestCmd.Flags().StringVarP(&repoTagManifestGet.version,
		"version",
		"v", "v2",
		"The version of manifest, valid value are \"v1\" and \"v2\", default is \"v2\".")
}

func getRepoTagManifest() {
	targetURL := repositoryURL + "/" + repoTagManifestGet.repoName +
		"/tags/" + repoTagManifestGet.tag +
		"/manifest?version=" + repoTagManifestGet.version
	utils.Get(targetURL)
}
