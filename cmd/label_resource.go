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
	"strconv"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// labelResourceCmd represents the resource command
var labelResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Get the resources that the label is referenced by.",
	Long:  `This endpoint let user get the resources that the label is referenced by. Only the replication policies are returned for now.`,
	Run: func(cmd *cobra.Command, args []string) {
		getLabelResource()
	},
}

var labelResourceGet struct {
	ID int64
}

func init() {
	labelCmd.AddCommand(labelResourceCmd)

	labelResourceCmd.Flags().Int64VarP(&labelResourceGet.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the label.")
	labelResourceCmd.MarkFlagRequired("id")
}

func getLabelResource() {
	targetURL := labelURL + "/" + strconv.FormatInt(labelResourceGet.ID, 10) + "/resources"
	utils.Get(targetURL)
}
