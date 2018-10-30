package utils

import (
	"crypto/tls"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// request is a new SuperAgent object with a setting of not verifying
// server's certificate chain and host name.
var request = gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})

func AgentGet() *gorequest.SuperAgent {
	return request
}

func Get(targetURL string) {
	fmt.Println("==> GET", targetURL)

	c, _ := CookieLoad()
	request.Get(targetURL).
		Set("Cookie", "harbor-lang=zh-cn; beegosessionID="+c).
		End(printStatus)
}

func Delete(targetURL string) {
	fmt.Println("==> DELETE", targetURL)

	c, _ := CookieLoad()
	request.Delete(targetURL).
		Set("Cookie", "harbor-lang=zh-cn; beegosessionID="+c).
		End(printStatus)
}

func Post(targetURL string, body string) {
	fmt.Println("==> POST", targetURL)

	c, _ := CookieLoad()
	request.Post(targetURL).
		Set("Cookie", "harbor-lang=zh-cn; beegosessionID="+c).
		Send(body).
		End(printStatus)
}

func Put(targetURL string, body string) {
	fmt.Println("==> PUT", targetURL)

	c, _ := CookieLoad()
	request.Put(targetURL).
		Set("Cookie", "harbor-lang=zh-cn; beegosessionID="+c).
		Send(body).
		End(printStatus)
}

// printStatus is a regular simple output callback function.
func printStatus(resp gorequest.Response, body string, errs []error) {
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
