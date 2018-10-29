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
	"net/url"

	"github.com/moooofly/harborctl/utils"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Harbor.",
	Long:  `Log in to Harbor with username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		loginHarbor()
	},
}

var li struct {
	username string
	password string
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&li.username, "username", "u", "", "(REQUIRED) Current login username.")
	loginCmd.MarkFlagRequired("username")
	loginCmd.Flags().StringVarP(&li.password, "password", "p", "", "Current user login password.")

	// TODO(mooofly): add `address` flag to overide config from .harborctl.yaml
}

func loginHarbor() {
	if li.password == "" {
		passwd, err := utils.ReadPasswordFromTerm()
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if passwd == "" {
			fmt.Println("error: Password Required.")
			return
		}

		li.password = passwd
	} else {
		fmt.Println("WARNING! Using --password via the CLI is insecure.")
	}

	targetURL := utils.URLGen("/login")
	fmt.Println("==> POST", targetURL)

	//fmt.Printf("==> username: %s   password: %s   escape: %s\n", li.username, li.password, url.QueryEscape(li.password))

	utils.Request.Post(targetURL).
		Set("Content-Type", "application/x-www-form-urlencoded;param=value").
		// NOTE:
		// After some experiments, conclude that the value of Cookie has two forms:
		// 1. Cookie:rem-username=admin; harbor-lang=zh-cn; beegosessionID=720210**3d76d6
		// 2. Cookie:rem-username=admin; harbor-lang=zh-cn;
		//
		// The fist one reuses the value of beegosessionID in Set-Cookie from response headers.
		// The second one is equivalent to a fresh login.
		//
		// Taking the second form just for long-live coding.
		Set("Cookie", "harbor-lang=zh-cn").
		Send("principal=" + li.username + "&password=" + url.QueryEscape(li.password)).
		End(loginProc)
}

// LoginProc is the callback function for login.
func loginProc(resp gorequest.Response, body string, errs []error) {
	for _, e := range errs {
		if e != nil {
			fmt.Println(e)
			return
		}
	}

	cookies := (*http.Response)(resp).Cookies()
	fmt.Println("<== Cookies:", cookies)

	sid, err := utils.CookieFilter(cookies, "beegosessionID")
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO(moooofly): can do something usefull based on response code
	fmt.Println("<== Rsp Status:", resp.Status)
	fmt.Println("<== Rsp Body:", body)

	err = utils.CookieSave(sid)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
