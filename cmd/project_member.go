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

// memberCmd represents the member command
var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "'/projects/{project_id}/members' API.",
	Long:  `The subcommand of '/projects/{project_id}/members' hierarchy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl project member --help\" for more information about this command.")
	},
}

func init() {
	projectCmd.AddCommand(memberCmd)

	initMemberCreate()
	initMemberGet()
	initMemberUpdate()
	initMemberDelete()
	initMemberList()
}

// memberCreateCmd represents the create command
var memberCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a member of a project. (add a member to a project)",
	Long: `Create project member relationship.

- The member can be one of the user_member and group_member
    - The user_member need to specify user_id or username.
	If the user already exist in harbor DB, specify the user_id, If does not exist in harbor DB, it will SearchAndOnBoard the user.
    - The group_member need to specify id or ldap_group_dn.
	If the group already exist in harbor DB. specify the user group's id, If does not exist, it will SearchAndOnBoard the group.

NOTE: currently, support user_member only.`,
	Run: func(cmd *cobra.Command, args []string) {
		createProjectMember()
	},
}

var prjMemberCreate struct {
	projectID int64

	// project_member
	roleID int64

	// user_member
	userID   int64
	username string
}

func initMemberCreate() {
	memberCmd.AddCommand(memberCreateCmd)

	memberCreateCmd.Flags().Int64VarP(&prjMemberCreate.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) The ID of project.")
	memberCreateCmd.MarkFlagRequired("project_id")

	memberCreateCmd.Flags().Int64VarP(&prjMemberCreate.roleID,
		"role_id",
		"r", 0,
		"(REQUIRED) Role ID. 1 for projectAdmin, 2 for developer, 3 for guest.")
	memberCreateCmd.MarkFlagRequired("role_id")

	memberCreateCmd.Flags().Int64VarP(&prjMemberCreate.userID,
		"user_id",
		"i", 0,
		`The ID of the user (wrong id cause nothing happend).
Either user_id or username must be set, and user_id with high priority.`)

	memberCreateCmd.Flags().StringVarP(&prjMemberCreate.username,
		"username",
		"n", "",
		`The name of the user (wrong name generates error).
Either user_id or username must be set, and user_id with high priority.`)
}

// NOTE: support only user_member right now.
// related issue: https://github.com/moooofly/harbor-go-client/issues/25
func createProjectMember() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMemberCreate.projectID, 10) + "/members"

	// projectMember defines the member settings of a project.
	type projectMember struct {
		RoleID int64 `json:"role_id,omitempty"`
		User   struct {
			UserID   int64  `json:"user_id,omitempty,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"member_user,omitempty"`
		Group struct {
			ID          int64  `json:"id,omitempty"`
			GroupName   string `json:"group_name,omitempty"`
			GroupType   int64  `json:"group_type,omitempty"`
			LdapGroupDn string `json:"ldap_group_dn,omitempty"`
		} `json:"member_group,omitempty"`
	}

	var prjMember projectMember
	prjMember.RoleID = prjMemberCreate.roleID
	prjMember.User.UserID = prjMemberCreate.userID
	prjMember.User.Username = prjMemberCreate.username

	p, err := json.Marshal(&prjMember)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	utils.Post(targetURL, string(p))
}

// memberGetCmd represents the get command
var memberGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a member of a project.",
	Long:  `This endpoint is aimed to get a member info of a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		getProjectMember()
	},
}

var prjMemberGet struct {
	projectID int64
	mID       int64
}

func initMemberGet() {
	memberCmd.AddCommand(memberGetCmd)

	memberGetCmd.Flags().Int64VarP(&prjMemberGet.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) The ID of project.")
	memberGetCmd.MarkFlagRequired("project_id")

	memberGetCmd.Flags().Int64VarP(&prjMemberGet.mID,
		"mid",
		"m", 0,
		"(REQUIRED) Member ID.")
	memberGetCmd.MarkFlagRequired("mid")
}

func getProjectMember() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMemberGet.projectID, 10) +
		"/members/" + strconv.FormatInt(prjMemberGet.mID, 10)

	utils.Get(targetURL)
}

// memberUpdateCmd represents the update command
var memberUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a member of a project.",
	Long:  `Update a member of a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateProjectMember()
	},
}

var prjMemberUpdate struct {
	projectID int64
	mID       int64

	RoleID int64 `json:"role_id"`
}

func initMemberUpdate() {
	memberCmd.AddCommand(memberUpdateCmd)

	memberUpdateCmd.Flags().Int64VarP(&prjMemberUpdate.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) The ID of project.")
	memberUpdateCmd.MarkFlagRequired("project_id")

	memberUpdateCmd.Flags().Int64VarP(&prjMemberUpdate.mID,
		"mid",
		"m", 0,
		"(REQUIRED) Member ID.")
	memberUpdateCmd.MarkFlagRequired("mid")

	memberUpdateCmd.Flags().Int64VarP(&prjMemberUpdate.RoleID,
		"role_id",
		"r", 0,
		"Role ID. 1 for projectAdmin, 2 for developer, 3 for guest.")
}

func updateProjectMember() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMemberUpdate.projectID, 10) +
		"/members/" + strconv.FormatInt(prjMemberUpdate.mID, 10)

	p, err := json.Marshal(&prjMemberUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	utils.Put(targetURL, string(p))
}

// memberDeleteCmd represents the delete command
var memberDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a member of a project.",
	Long:  `Delete a member of a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteProjectMember()
	},
}

var prjMemberDelete struct {
	projectID int64
	mID       int64
}

func initMemberDelete() {
	memberCmd.AddCommand(memberDeleteCmd)

	memberDeleteCmd.Flags().Int64VarP(&prjMemberDelete.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) The ID of project.")
	memberDeleteCmd.MarkFlagRequired("project_id")

	memberDeleteCmd.Flags().Int64VarP(&prjMemberDelete.mID,
		"mid",
		"m", 0,
		"(REQUIRED) The ID of member.")
	memberDeleteCmd.MarkFlagRequired("mid")
}

func deleteProjectMember() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMemberDelete.projectID, 10) +
		"/members/" + strconv.FormatInt(prjMemberDelete.mID, 10)
	utils.Delete(targetURL)
}

// memberlistCmd represents the list command
var memberlistCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all members information of a project.",
	Long:  `Get all members information of a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		listProjectMember()
	},
}

var prjMemberList struct {
	projectID  int64
	entityname string
}

func initMemberList() {
	memberCmd.AddCommand(memberlistCmd)

	memberlistCmd.Flags().Int64VarP(&prjMemberList.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) The ID of project.")
	memberlistCmd.MarkFlagRequired("project_id")

	memberlistCmd.Flags().StringVarP(&prjMemberList.entityname,
		"username",
		"n", "",
		"The entity (user) name to search.")
}

func listProjectMember() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMemberList.projectID, 10) +
		"/members?entityname=" + prjMemberList.entityname

	utils.Get(targetURL)
}
