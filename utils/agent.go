package utils

import (
	"crypto/tls"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Request is a new SuperAgent object with a setting of not verifying
// server's certificate chain and host name.
var Request = gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})

// PrintStatus is a regular callback function.
func PrintStatus(resp gorequest.Response, body string, errs []error) {
	fmt.Println("<== ")
	for _, e := range errs {
		if e != nil {
			fmt.Println(e)
			return
		}
	}

	fmt.Println("<== Rsp Status:", resp.Status)
	fmt.Printf("<== Rsp Body: %s\n", body)
}
