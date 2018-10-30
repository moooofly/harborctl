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
	"fmt"
	"net/http"

	"github.com/moooofly/harborctl/utils"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

var logoutURL string

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from Harbor.",
	Long: `Log out current user from Harbor.

NOTE: multiple logout cause nothing happened.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logoutURL = utils.URLGen("/log_out")
	},
	Run: func(cmd *cobra.Command, args []string) {
		logoutHarbor()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

func logoutHarbor() {
	targetURL := logoutURL
	fmt.Println("==> GET", targetURL)

	utils.AgentGet().Get(targetURL).End(logoutProc)
}

// logoutProc is the callback function for logout.
func logoutProc(resp gorequest.Response, body string, errs []error) {
	for _, e := range errs {
		if e != nil {
			fmt.Println(e)
			return
		}
	}

	fmt.Println("<== Rsp Cookies:", (*http.Response)(resp).Cookies())
	fmt.Println("<== Rsp Status:", resp.Status)
	fmt.Println("<== Rsp Body:", body)

	utils.CookieClean()
}
