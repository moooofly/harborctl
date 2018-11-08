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

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// targetPingCmd represents the ping command
var targetPingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping validates target.",
	Long:  `This endpoint let user to ping a target to validate whether it is reachable and whether the credential is valid.`,
	Run: func(cmd *cobra.Command, args []string) {
		pingTarget()
	},
}

var targetPing struct {
	ID       int64  `json:"id"`
	Endpoint string `json:"endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
	Insecure bool   `json:"insecure"`
}

func init() {
	targetCmd.AddCommand(targetPingCmd)

	targetPingCmd.Flags().Int64VarP(&targetPing.ID,
		"id",
		"i", 0,
		"(REQUIRED) The replication's target ID.")
	targetPingCmd.MarkFlagRequired("id")

	targetPingCmd.Flags().StringVarP(&targetPing.Endpoint,
		"endpoint",
		"e", "",
		"(REQUIRED) The target address URL string. (NOTE: must be of format '<http|https>:xx.xx.xx.xx')")
	targetPingCmd.MarkFlagRequired("endpoint")

	targetPingCmd.Flags().StringVarP(&targetPing.Username,
		"username",
		"u", "",
		"(REQUIRED) The username used to login target server.")
	targetPingCmd.MarkFlagRequired("username")

	targetPingCmd.Flags().StringVarP(&targetPing.Password,
		"password",
		"p", "",
		"(REQUIRED) The password used to login target server.")
	targetPingCmd.MarkFlagRequired("password")

	targetPingCmd.Flags().BoolVarP(&targetPing.Insecure,
		"insecure",
		"", false,
		"(REQUIRED) Whether or not the certificate will be verified when Harbor tries to access the server.")
	targetPingCmd.MarkFlagRequired("insecure")
}

func pingTarget() {
	targetURL := targetsURL + "/ping"

	p, err := json.Marshal(&targetPing)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}
