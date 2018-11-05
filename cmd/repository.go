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
	"strconv"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

var repositoryURL string

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "'/repositories' API.",
	Long:  `The subcommand of '/repositoriesj' hierachy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		repositoryURL = utils.URLGen("/api/repositories")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl repository --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(repositoryCmd)

	initRepoGet()
	initRepoDelete()
	initRepoUpdate()
}

// repoGetCmd represents the get command
var repoGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get repositories info accompany with a relevant project by project_id.",
	Long: `This endpoint let user get repositories info accompanying with a relevant project by project_id.

NOTE:

- You can filter results by repo name ("q=") and label ID ("label_id=").
- You can sort results in either ascending or descending order ('-' stands for descending order) by
    - repo name ('name' or '-name')
    - creation_time ('creation_time' or '-creation_time')
    - update_time ('update_time' or '-update_time')
`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepositoryInfo()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/13)
var repoGet struct {
	projectID int64
	q         string
	sort      string
	labelID   int64
	page      int64
	pageSize  int64
}

func initRepoGet() {
	repositoryCmd.AddCommand(repoGetCmd)

	repoGetCmd.Flags().Int64VarP(&repoGet.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Relevant project ID.")
	repoGetCmd.MarkFlagRequired("project_id")

	repoGetCmd.Flags().StringVarP(&repoGet.q,
		"repo_name",
		"q", "",
		"Repo name for filtering results.")

	repoGetCmd.Flags().StringVarP(&repoGet.sort,
		"sort",
		"o", "",
		"Sort method, valid values include: 'name’, '-name’, 'creation_time’, '-creation_time’, 'update_time’, '-update_time’. Here '-' stands for descending order.")

	repoGetCmd.Flags().Int64VarP(&repoGet.labelID,
		"label_id",
		"l", 0,
		"The ID of label used to filter the result.")

	repoGetCmd.Flags().Int64VarP(&repoGet.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")

	repoGetCmd.Flags().Int64VarP(&repoGet.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func getRepositoryInfo() {
	targetURL := repositoryURL + "?project_id=" + strconv.FormatInt(repoGet.projectID, 10) +
		"&q=" + repoGet.q +
		"&sort=" + repoGet.sort +
		"&label_id=" + strconv.FormatInt(repoGet.labelID, 10) +
		"&page=" + strconv.FormatInt(repoGet.page, 10) +
		"&page_size=" + strconv.FormatInt(repoGet.pageSize, 10)
	utils.Get(targetURL)
}

// repoDeleteCmd represents the delete command
var repoDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repository by repo_name.",
	Long:  `This endpoint let user delete a repository by repo_name.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteRepository()
	},
}

var repoDelete struct {
	repoName string
}

func initRepoDelete() {
	repositoryCmd.AddCommand(repoDeleteCmd)

	repoDeleteCmd.Flags().StringVarP(&repoDelete.repoName,
		"repo_name",
		"n", "",
		"(REQUIRED) The name of repository which will be deleted.")
	repoDeleteCmd.MarkFlagRequired("repo_name")
}

func deleteRepository() {
	targetURL := repositoryURL + "/" + repoDelete.repoName
	utils.Delete(targetURL)
}

// repoUpdateCmd represents the update command
var repoUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update description of the repository.",
	Long:  `This endpoint is used to update description of the repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateRepoDescription()
	},
}

var repoUpdate struct {
	repoName string

	Description string `json:"description"`
}

func initRepoUpdate() {
	repositoryCmd.AddCommand(repoUpdateCmd)

	repoUpdateCmd.Flags().StringVarP(&repoUpdate.repoName,
		"repo_name",
		"n", "",
		"(REQUIRED) The name of repository which will be updated.")
	repoUpdateCmd.MarkFlagRequired("repo_name")

	repoUpdateCmd.Flags().StringVarP(&repoUpdate.Description,
		"description",
		"d", "",
		"(REQUIRED) The description of the repository.")
	repoUpdateCmd.MarkFlagRequired("description")
}

func updateRepoDescription() {
	targetURL := repositoryURL + "/" + repoUpdate.repoName

	p, err := json.Marshal(&repoUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
