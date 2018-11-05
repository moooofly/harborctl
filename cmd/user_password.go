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

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Change the password of a user that already exists.",
	Long: `This endpoint is for user to update password.

NOTE:
- Users with the admin role can change any user's password.
- Guest users can change only their own password.`,
	Run: func(cmd *cobra.Command, args []string) {
		updatePassword()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/11)
var updateUserPassword struct {
	userID int64

	// The attribute ‘old_password’ is optional when the API is called by the system administrator.
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func init() {
	userCmd.AddCommand(passwordCmd)

	passwordCmd.Flags().Int64VarP(&updateUserPassword.userID,
		"user_id",
		"i", 0,
		"(REQUIRED) Registered user ID.")
	passwordCmd.MarkFlagRequired("user_id")

	passwordCmd.Flags().StringVarP(&updateUserPassword.OldPassword,
		"old_password",
		"o", "",
		"(REQUIRED) The user’s existing password. The attribute ‘old_password’ is optional when the API is called by the system administrator.")
	passwordCmd.MarkFlagRequired("old_password")

	passwordCmd.Flags().StringVarP(&updateUserPassword.NewPassword,
		"new_password",
		"n", "",
		"(REQUIRED) New password for marking as to be updated.")
	passwordCmd.MarkFlagRequired("new_password")
}

func updatePassword() {
	targetURL := userURL + "/" + strconv.FormatInt(updateUserPassword.userID, 10) + "/password"

	p, err := json.Marshal(&updateUserPassword)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
