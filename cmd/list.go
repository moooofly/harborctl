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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects by filter.",
	Long:  `This endpoint returns all projects created by Harbor which can be filtered by (project) name, owner and public property.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectList()
	},
}

var prjList struct {
	name string

	// NOTE:
	// As per swagger file defination, the type of 'public' is boolean.
	// I change the type of 'public' from boolean to string in order for three-valued logic
	// 1. If true/TRUE/1, only public projects returned
	// 2. If false/FALSE/0, only private projects returned
	// 3. If not set (namely empty string), both private and public projects returned

	// NOTE:
	// The content returned by "GET /api/projects" API depends on login status
	// 1. If in login state, it can obtain both public projects and private projects which created by this login user.
	// 2. If not in login state, it can obtain only public projects.
	public   string
	owner    string
	page     int32
	pageSize int32
}

func init() {
	projectCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&prjList.name,
		"name",
		"n", "",
		"The name of project.")
	listCmd.Flags().StringVarP(&prjList.public,
		"public",
		"k", "",
		"The project is public or private.")
	listCmd.Flags().StringVarP(&prjList.owner,
		"owner",
		"o", "",
		"The name of project owner.")
	listCmd.Flags().Int32VarP(&prjList.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")
	listCmd.Flags().Int32VarP(&prjList.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func projectList() {
	targetURL := utils.URLGen("/api/projects") + "?name=" + prjList.name +
		"&public=" + prjList.public +
		"&owner=" + prjList.owner +
		"&page=" + strconv.FormatInt(int64(prjList.page), 10) +
		"&page_size=" + strconv.FormatInt(int64(prjList.pageSize), 10)
	fmt.Println("==> GET", targetURL)

	c, err := utils.CookieLoad()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.Request.Get(targetURL).
		Set("Cookie", "harbor-lang=zh-cn; beegosessionID="+c.BeegosessionID).
		// NOTE:
		// If need, can obtain the total count of projects from Rsp Header by X-Total-Count
		End(utils.PrintStatus)
}
