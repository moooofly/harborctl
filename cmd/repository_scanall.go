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
	"encoding/json"
	"fmt"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// scanallCmd represents the scanall command
var scanallCmd = &cobra.Command{
	Use:   "scanall",
	Short: "Scan all images of the registry. (NOTE: need Clair deployed)",
	Long: `The server will launch different jobs to scan each image on the registry, so this is equivalent to calling the API to
scan the image one by one in background, so there's no way to track the overall status of the "scan all" action.

NOTE:
- Only system admin has permission to call this API.
- This function depends on Clair. Without Clair will trigger "503 Service Unavailable" error.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanAll()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/16)
var scan struct {
	ProjectID int64 `json:"project_id"`
}

func init() {
	repositoryCmd.AddCommand(scanallCmd)

	scanallCmd.Flags().Int64VarP(&scan.ProjectID,
		"project_id",
		"j", 0,
		"(REQUIRED) When this parameter is set, only the images under the project identified by project_id will be scanned.")
	scanallCmd.MarkFlagRequired("project_id")
}

func scanAll() {
	targetURL := repositoryURL + "/scanAll"

	p, err := json.Marshal(&scan)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}
