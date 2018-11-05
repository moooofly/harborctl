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

// sysadminCmd represents the sysadmin command
var sysadminCmd = &cobra.Command{
	Use:   "sysadmin",
	Short: "Update a registered user to change to be an administrator of Harbor.",
	Long:  `This endpoint let a registered user change to be an administrator of Harbor.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateSysadmin()
	},
}

var updateUserSysadmin struct {
	userID int64

	HasAdminRole bool `json:"has_admin_role"`
}

func init() {
	userCmd.AddCommand(sysadminCmd)

	sysadminCmd.Flags().Int64VarP(&updateUserSysadmin.userID,
		"user_id",
		"i", 0,
		"(REQUIRED) Registered user ID.")
	sysadminCmd.MarkFlagRequired("user_id")

	sysadminCmd.Flags().BoolVarP(&updateUserSysadmin.HasAdminRole,
		"has_admin_role",
		"a", false,
		"(REQUIRED) Toggle a user to admin or not. (NOTE: should be used as --has_admin_role=[true|false] or -a=[true|false] format)")
	sysadminCmd.MarkFlagRequired("has_admin_role")
}

func updateSysadmin() {
	targetURL := userURL + "/" + strconv.FormatInt(updateUserSysadmin.userID, 10) + "/sysadmin"

	p, err := json.Marshal(&updateUserSysadmin)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
