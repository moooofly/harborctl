// Copyright © 2019 moooofly <centos.sf@gmail.com>
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

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

//  represents the prov command
var provCmd = &cobra.Command{
	Use:   "prov",
	Short: "'/chartrepo/{repo}/prov' API.",
	Long:  `The subcommand of '/chartrepo/{repo}/prov' hierachy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl chartrepo prov --help\" for more information about this command.")
	},
}

func init() {
	chartrepoCmd.AddCommand(provCmd)

	initProvUpload()
}

// uploadCmd represents the prov command
var provUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a provance file to the specified project.",
	Long:  `Upload a provance file to the specified project. The provance file should be targeted for an existing chart file.`,
	Run: func(cmd *cobra.Command, args []string) {
		uploadProv()
	},
}

var provUpload struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName string
	provFile    string
}

func initProvUpload() {
	provCmd.AddCommand(provUploadCmd)

	provUploadCmd.Flags().StringVarP(&provUpload.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	provUploadCmd.MarkFlagRequired("project_name")

	provUploadCmd.Flags().StringVarP(&provUpload.provFile,
		"prov_file",
		"p", "",
		"(REQUIRED) The provenance file name")
	provUploadCmd.MarkFlagRequired("prov_file")
}

func uploadProv() {
	targetURL := chartrepoURL + "/" + provUpload.projectName + "/prov"
	utils.Multipart(targetURL, provUpload.provFile)
}
