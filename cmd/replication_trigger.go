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

// triggerCmd represents the trigger command
var triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Trigger the replication by the specified policy ID.",
	Long:  `This endpoint is used to trigger a replication by the specified policy ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		trigger()
	},
}

var replicationTrigger struct {
	PolicyID int64 `json:"policy_id"`
}

func init() {
	replicationCmd.AddCommand(triggerCmd)

	triggerCmd.Flags().Int64VarP(&replicationTrigger.PolicyID,
		"policy_id",
		"i", 0,
		"(REQUIRED) The ID of replication policy.")
	triggerCmd.MarkFlagRequired("policy_id")
}

func trigger() {
	targetURL := replicationURL

	p, err := json.Marshal(&replicationTrigger)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	utils.Post(targetURL, string(p))
}
