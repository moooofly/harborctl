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

// retagCmd represents the retag command
var retagCmd = &cobra.Command{
	Use:   "retag",
	Short: "Retag a image.",
	Long:  `This endpoint tags an existing image with another tag in this repo, source images can be in different repos or projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		retagImage()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/17)
var imageRetag struct {
	Tag      string `json:"tag"`
	srcImage string `json:"src_image"`
	override bool   `json:"override"`
}

func init() {
	tagCmd.AddCommand(retagCmd)

	retagCmd.Flags().StringVarP(&imageRetag.Tag,
		"tag",
		"t", "",
		"(REQUIRED) New tag to be created. (NOTE: unknown format)")
	retagCmd.MarkFlagRequired("tag")

	retagCmd.Flags().StringVarP(&imageRetag.srcImage,
		"src_image",
		"s", "",
		"(REQUIRED) Source image to be retagged, e.g. 'stage/app:v1.0'")
	retagCmd.MarkFlagRequired("src_image")

	retagCmd.Flags().BoolVarP(&imageRetag.override,
		"override",
		"o", false,
		"If the target tag already exists, whether to override it.")
}

func retagImage() {
	targetURL := repositoryURL + "/" + repoTagGet.repoName + "/tags"

	p, err := json.Marshal(&imageRetag)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}
