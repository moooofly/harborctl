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
	"fmt"
	"strconv"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project by project_id.",
	Long:  `This endpoint is aimed to delete a project by project_id.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectDelete()
	},
}

var prjDelete struct {
	projectID int32
}

func init() {
	projectCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().Int32VarP(&prjDelete.projectID,
		"project_id",
		"i", 0,
		"(REQUIRED) Project ID of project which will be deleted.")
	deleteCmd.MarkFlagRequired("project_id")
}

func projectDelete() {
	targetURL := utils.URLGen("/api/projects") + "/" + strconv.FormatInt(int64(prjDelete.projectID), 10)
	fmt.Println("==> DELETE", targetURL)

	c, err := utils.CookieLoad()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.Request.Delete(targetURL).
		Set("Cookie", "harbor-lang=zh-cn; beegosessionID="+c.BeegosessionID).
		End(utils.PrintStatus)
}
