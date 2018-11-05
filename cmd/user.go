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

var userURL string

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "'/users' API.",
	Long:  `The subcommand of '/users' hierachy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		userURL = utils.URLGen("/api/users")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl user --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	initUserList()
	initUserGet()
	initUserDelete()
	initUserCreate()
	initUserUpdate()
}

// userListCmd represents the list command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all registered users of Harbor (not include admin user).",
	Long: `This endpoint is for user to search registered users, support for filtering results with username. Notice, by now this operation is only for administrator.

NOTE: The results returned will not include admin user profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		listUser()
	},
}

var userList struct {
	username string
	email    string
	page     int64
	pageSize int64
}

func initUserList() {
	userCmd.AddCommand(userListCmd)

	userListCmd.Flags().StringVarP(&userList.username,
		"username",
		"u", "",
		"Username for filtering results.")
	userListCmd.Flags().StringVarP(&userList.email,
		"email",
		"e", "",
		"Email for filtering results.")
	userListCmd.Flags().Int64VarP(&userList.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")
	userListCmd.Flags().Int64VarP(&userList.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func listUser() {
	targetURL := userURL + "?username=" + userList.username +
		"&email=" + userList.email +
		"&page=" + strconv.FormatInt(userList.page, 10) +
		"&page_size=" + strconv.FormatInt(userList.pageSize, 10)

	utils.Get(targetURL)
}

// userGetCmd represents the get command
var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user's profile.",
	Long:  `Get a user's profile by user_id.`,
	Run: func(cmd *cobra.Command, args []string) {
		getUser()
	},
}

var userGet struct {
	userID int64
}

func initUserGet() {
	userCmd.AddCommand(userGetCmd)

	userGetCmd.Flags().Int64VarP(&userGet.userID,
		"user_id",
		"i", 0,
		"(REQUIRED) Registered user ID.")
	userGetCmd.MarkFlagRequired("user_id")
}

func getUser() {
	targetURL := userURL + "/" + strconv.FormatInt(userGet.userID, 10)
	utils.Get(targetURL)
}

// userDeleteCmd represents the delete command
var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Mark a registered user as be removed.",
	Long:  `This endpoint let administrator of Harbor mark a registered user as be removed. It actually won't be deleted from DB.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteUser()
	},
}

var userDelete struct {
	userID int64
}

func initUserDelete() {
	userCmd.AddCommand(userDeleteCmd)

	userDeleteCmd.Flags().Int64VarP(&userDelete.userID,
		"user_id",
		"i", 0,
		"(REQUIRED) Registered user ID.")
	userDeleteCmd.MarkFlagRequired("user_id")
}

func deleteUser() {
	targetURL := userURL + "/" + strconv.FormatInt(userDelete.userID, 10)
	utils.Delete(targetURL)
}

// userCreateCmd represents the create command
var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new user account.",
	Long:  `This endpoint is to create a user if the user does not already exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		createUser()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/7)
var userCreate struct {
	// the parameters that are effectively used by this API
	Username     string `json:"username"`
	Password     string `json:"password"`
	HasAdminRole bool   `json:"has_admin_role"`
	Email        string `json:"email"`
	Realname     string `json:"realname"`
	Comment      string `json:"comment"`

	// the parameters that take no effect
	UserID    int64  `json:"user_id,omitempty"`
	Deleted   bool   `json:"deleted,omitempty"`
	RoleName  string `json:"role_name,omitempty"`
	RoleID    int64  `json:"role_id,omitempty"`
	ResetUUID string `json:"reset_uuid,omitempty"`
	Salt      string `json:"salt,omitempty"`

	// the parameters that should not appear in this API
	CreationTime string `json:"creation_time,omitempty"`
	UpdateTime   string `json:"update_time,omitempty"`
}

func initUserCreate() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().StringVarP(&userCreate.Username,
		"username",
		"u", "",
		"(REQUIRED) The name of the user you create.")
	userCreateCmd.MarkFlagRequired("username")

	userCreateCmd.Flags().StringVarP(&userCreate.Password,
		"password",
		"p", "",
		"(REQUIRED) The password of the user you create.")
	userCreateCmd.MarkFlagRequired("password")

	userCreateCmd.Flags().BoolVarP(&userCreate.HasAdminRole,
		"has_admin_role",
		"a", false,
		"(REQUIRED) Mark a user whether is admin or not. (NOTE: should be used as --has_admin_role=[true|false] or -a=[true|false] format)")
	userCreateCmd.MarkFlagRequired("has_admin_role")

	userCreateCmd.Flags().StringVarP(&userCreate.Email,
		"email",
		"e", "",
		"(REQUIRED) The email of the user you create.")
	userCreateCmd.MarkFlagRequired("email")

	userCreateCmd.Flags().StringVarP(&userCreate.Realname,
		"realname",
		"r", "",
		"(REQUIRED) The realname of the user you create.")
	userCreateCmd.MarkFlagRequired("realname")

	userCreateCmd.Flags().StringVarP(&userCreate.Comment,
		"comment",
		"c", "",
		"(REQUIRED) The comment with the user you create.")
	userCreateCmd.MarkFlagRequired("comment")

	userCreateCmd.Flags().Int64VarP(&userCreate.UserID,
		"user_id",
		"", 0,
		"NOTE: This register ID is assigned by Harbor.")

	userCreateCmd.Flags().BoolVarP(&userCreate.Deleted,
		"deleted",
		"", false,
		"NOTE: Need to clarify this parameter.")

	userCreateCmd.Flags().StringVarP(&userCreate.RoleName,
		"role_name",
		"", "",
		"The role name of user you create.")

	userCreateCmd.Flags().Int64VarP(&userCreate.RoleID,
		"role_id",
		"", 0,
		"The role ID of user you create.")

	userCreateCmd.Flags().StringVarP(&userCreate.ResetUUID,
		"reset_uuid",
		"", "",
		"NOTE: Need to clarify this parameter.")

	userCreateCmd.Flags().StringVarP(&userCreate.Salt,
		"salt",
		"", "",
		"NOTE: Need to clarify this parameter.")
}

func createUser() {
	targetURL := userURL

	p, err := json.Marshal(&userCreate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// userUpdateCmd represents the update command
var userUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a registered user to change his profile.",
	Long:  `This endpoint let a registered user change his profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateUser()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/10)
var userUpdate struct {
	userID int64

	Email    string `json:"email"`
	Realname string `json:"realname"`
	Comment  string `json:"comment,omitempty"`
}

func initUserUpdate() {
	userCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().Int64VarP(&userUpdate.userID,
		"user_id",
		"i", 0,
		"(REQUIRED) Registered user ID.")
	userUpdateCmd.MarkFlagRequired("user_id")

	userUpdateCmd.Flags().StringVarP(&userUpdate.Email,
		"email",
		"e", "",
		"(REQUIRED) The email to update. NOTE: must be with a valid email format.")
	userUpdateCmd.MarkFlagRequired("email")

	userUpdateCmd.Flags().StringVarP(&userUpdate.Realname,
		"realname",
		"r", "",
		"(REQUIRED) The realname to update. NOTE: must contain one character at least.")
	userUpdateCmd.MarkFlagRequired("realname")

	userUpdateCmd.Flags().StringVarP(&userUpdate.Comment,
		"comment",
		"c", "",
		"The comment to update.")
}

func updateUser() {
	targetURL := userURL + "/" + strconv.FormatInt(userUpdate.userID, 10)

	p, err := json.Marshal(&userUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
