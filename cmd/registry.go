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

var targetsURL string

// targetCmd represents the target command
var targetCmd = &cobra.Command{
	Use:   "registry",
	Short: "'/targets' API.",
	Long:  `The subcommand of '/targets' hierarchy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		targetsURL = utils.URLGen("/api/targets")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl registry --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(targetCmd)

	initTargetList()
	initTargetGet()
	initTargetDelete()
	initTargetCreate()
	initTargetUpdate()
}

// targetListCmd represents the list command
var targetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List targets filtered by target name.",
	Long:  `This endpoint let user list targets filtered by target name, if name is nil, list returns all targets.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTarget()
	},
}

var targetList struct {
	name string
}

func initTargetList() {
	targetCmd.AddCommand(targetListCmd)

	targetListCmd.Flags().StringVarP(&targetList.name,
		"name",
		"n", "",
		"The replication's target name.")
}

func listTarget() {
	targetURL := targetsURL + "?name=" + targetList.name
	utils.Get(targetURL)
}

// targetGetCmd represents the get command
var targetGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get target by target ID.",
	Long:  `This endpoint let user get a specific target by target ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTarget()
	},
}

var targetGet struct {
	ID int64
}

func initTargetGet() {
	targetCmd.AddCommand(targetGetCmd)

	targetGetCmd.Flags().Int64VarP(&targetGet.ID,
		"id",
		"i", 0,
		"(REQUIRED) The replication's target ID.")
	targetGetCmd.MarkFlagRequired("id")
}

func getTarget() {
	targetURL := targetsURL + "/" + strconv.FormatInt(targetGet.ID, 10)
	utils.Get(targetURL)
}

// targetDeleteCmd represents the delete command
var targetDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete target by target ID.",
	Long:  `This endpoint let user delete a specific target by target ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteTarget()
	},
}

var targetDelete struct {
	ID int64
}

func initTargetDelete() {
	targetCmd.AddCommand(targetDeleteCmd)

	targetDeleteCmd.Flags().Int64VarP(&targetDelete.ID,
		"id",
		"i", 0,
		"(REQUIRED) The replication's target ID.")
	targetDeleteCmd.MarkFlagRequired("id")
}

func deleteTarget() {
	targetURL := targetsURL + "/" + strconv.FormatInt(targetDelete.ID, 10)
	utils.Delete(targetURL)
}

// targetCreateCmd represents the create command
var targetCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new replication target.",
	Long:  `This endpoint let user to create a new replication target.`,
	Run: func(cmd *cobra.Command, args []string) {
		createTarget()
	},
}

var targetCreate struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
	Insecure bool   `json:"insecure"`
}

func initTargetCreate() {
	targetCmd.AddCommand(targetCreateCmd)

	targetCreateCmd.Flags().StringVarP(&targetCreate.Name,
		"name",
		"n", "",
		"(REQUIRED) The target name.")
	targetCreateCmd.MarkFlagRequired("name")

	targetCreateCmd.Flags().StringVarP(&targetCreate.Endpoint,
		"endpoint",
		"e", "",
		"(REQUIRED) The target address URL string. (NOTE: must be of format '<http|https>:xx.xx.xx.xx')")
	targetCreateCmd.MarkFlagRequired("endpoint")

	targetCreateCmd.Flags().StringVarP(&targetCreate.Username,
		"username",
		"u", "",
		"(REQUIRED) The username used to login target server.")
	targetCreateCmd.MarkFlagRequired("username")

	targetCreateCmd.Flags().StringVarP(&targetCreate.Password,
		"password",
		"p", "",
		"(REQUIRED) The password used to login target server.")
	targetCreateCmd.MarkFlagRequired("password")

	targetCreateCmd.Flags().BoolVarP(&targetCreate.Insecure,
		"insecure",
		"", false,
		"(REQUIRED) Whether or not the certificate will be verified when Harbor tries to access the server.")
	targetCreateCmd.MarkFlagRequired("insecure")
}

func createTarget() {
	targetURL := targetsURL

	p, err := json.Marshal(&targetCreate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// targetUpdateCmd represents the update command
var targetUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an already existing replication target.",
	Long:  `This endpoint let user to update a specific replication target.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateTarget()
	},
}

var targetUpdate struct {
	ID int64

	Name     string `json:"name,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Insecure bool   `json:"insecure,omitempty"`
}

func initTargetUpdate() {
	targetCmd.AddCommand(targetUpdateCmd)

	targetUpdateCmd.Flags().Int64VarP(&targetUpdate.ID,
		"id",
		"i", 0,
		"(REQUIRED) The replication's target ID.")
	targetUpdateCmd.MarkFlagRequired("id")

	targetUpdateCmd.Flags().StringVarP(&targetUpdate.Name,
		"name",
		"n", "",
		"The target name.")

	targetUpdateCmd.Flags().StringVarP(&targetUpdate.Endpoint,
		"endpoint",
		"e", "",
		"The target address URL string. (NOTE: must be of format '<http|https>:xx.xx.xx.xx')")

	targetUpdateCmd.Flags().StringVarP(&targetUpdate.Username,
		"username",
		"u", "",
		"The username used to login target server.")

	targetUpdateCmd.Flags().StringVarP(&targetUpdate.Password,
		"password",
		"p", "",
		"The password used to login target server.")

	targetUpdateCmd.Flags().BoolVarP(&targetUpdate.Insecure,
		"insecure",
		"", false,
		"Whether or not the certificate will be verified when Harbor tries to access the server.")
}

func updateTarget() {
	targetURL := targetsURL + "/" + strconv.FormatInt(targetUpdate.ID, 10)

	p, err := json.Marshal(&targetUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
