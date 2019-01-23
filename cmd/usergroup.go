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
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

var usergroupURL string

// usergroupCmd represents the usergroup command
var usergroupCmd = &cobra.Command{
	Use:   "usergroup",
	Short: "'/usergroups' API.",
	Long:  `The subcommand of '/usergroups' hierachy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		usergroupURL = utils.URLGen("/api/usergroups")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl usergroup --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(usergroupCmd)

	initUsergroupList()
	initUsergroupGet()
	initUsergroupDelete()
	initUsergroupCreate()
	initUsergroupUpdate()
}

// listCmd represents the list command
var usergroupListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all user groups information.",
	Long:  `This endpoint is for user to get all user groups information.`,
	Run: func(cmd *cobra.Command, args []string) {
		listUsergroup()
	},
}

func initUsergroupList() {
	usergroupCmd.AddCommand(usergroupListCmd)
}

func listUsergroup() {
	targetURL := usergroupURL
	utils.Get(targetURL)
}

// getCmd represents the get command
var usergroupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user group information by group_id.",
	Long:  `This endpoint is for user to get a user group information by group_id.`,
	Run: func(cmd *cobra.Command, args []string) {
		getUsergroup()
	},
}

var usergroupGet struct {
	groupID int64
}

func initUsergroupGet() {
	usergroupCmd.AddCommand(usergroupGetCmd)

	usergroupGetCmd.Flags().Int64VarP(&usergroupGet.groupID,
		"group_id",
		"i", 0,
		"The Group ID.")
	usergroupGetCmd.MarkFlagRequired("group_id")
}

func getUsergroup() {
	targetURL := usergroupURL + "/" + strconv.FormatInt(usergroupGet.groupID, 10)
	utils.Get(targetURL)
}

// deleteCmd represents the delete command
var usergroupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user group by group_id.",
	Long:  `This endpoint is for user to delete a user group by group_id.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteUsergroup()
	},
}

var usergroupDelete struct {
	groupID int64
}

func initUsergroupDelete() {
	usergroupCmd.AddCommand(usergroupDeleteCmd)

	usergroupDeleteCmd.Flags().Int64VarP(&usergroupDelete.groupID,
		"group_id",
		"i", 0,
		"(REQUIRED) Group ID.")
	usergroupDeleteCmd.MarkFlagRequired("group_id")
}

func deleteUsergroup() {
	targetURL := usergroupURL + "/" + strconv.FormatInt(usergroupDelete.groupID, 10)
	utils.Delete(targetURL)
}

// createCmd represents the create command
var usergroupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user group.",
	Long:  `This endpoint is for user to create a user group.`,
	Run: func(cmd *cobra.Command, args []string) {
		createUsergroup()
	},
}

var usergroupCreate struct {
	ID          int64  `json:"id"`
	GroupName   string `json:"group_name"`
	GroupType   int64  `json:"group_type"`
	LdapGroupDN string `json:"ldap_group_dn"`
}

func initUsergroupCreate() {
	usergroupCmd.AddCommand(usergroupCreateCmd)

	usergroupCreateCmd.Flags().Int64VarP(&usergroupCreate.ID,
		"id",
		"i", 0,
		"The ID of the user group")

	usergroupCreateCmd.Flags().StringVarP(&usergroupCreate.GroupName,
		"group_name",
		"n", "",
		"(REQUIRED) The name of the user group.")
	usergroupCreateCmd.MarkFlagRequired("group_name")

	usergroupCreateCmd.Flags().Int64VarP(&usergroupCreate.GroupType,
		"group_type",
		"t", 0,
		"The group type, 1 for LDAP group.")

	usergroupCreateCmd.Flags().StringVarP(&usergroupCreate.LdapGroupDN,
		"ldap_group_dn",
		"d", "",
		"The DN of the LDAP group if group type is 1 (LDAP group).")
}

func createUsergroup() {
	targetURL := usergroupURL

	p, err := json.Marshal(&usergroupCreate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// updateCmd represents the update command
var usergroupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a group information by group_id.",
	Long:  `This endpoint is for user to update a group information by group_id.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateUsergroup()
	},
}

var usergroupUpdate struct {
	ID          int64  `json:"id"`
	GroupName   string `json:"group_name"`
	GroupType   int64  `json:"group_type"`
	LdapGroupDN string `json:"ldap_group_dn"`
}

func initUsergroupUpdate() {
	usergroupCmd.AddCommand(usergroupUpdateCmd)

	usergroupUpdateCmd.Flags().Int64VarP(&usergroupUpdate.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the user group.")
	usergroupUpdateCmd.MarkFlagRequired("id")

	usergroupUpdateCmd.Flags().StringVarP(&usergroupUpdate.GroupName,
		"group_name",
		"n", "",
		"(REQUIRED) The name of the user group.")
	usergroupUpdateCmd.MarkFlagRequired("group_name")

	usergroupUpdateCmd.Flags().Int64VarP(&usergroupUpdate.GroupType,
		"group_type",
		"t", 0,
		"The group type, 1 for LDAP group.")

	usergroupUpdateCmd.Flags().StringVarP(&usergroupUpdate.LdapGroupDN,
		"ldap_group_dn",
		"d", "",
		"The DN of the LDAP group if group type is 1 (LDAP group).")
}

func updateUsergroup() {
	targetURL := usergroupURL + "/" + strconv.FormatInt(usergroupUpdate.ID, 10)

	p, err := json.Marshal(&usergroupUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
