package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Beegocookie is for beegosessionID storage
type Beegocookie struct {
	BeegosessionID string `yaml:"beegosessionID"`
}

// CookieFilter filters specific cookie string.
func CookieFilter(cookies []*http.Cookie, filter string) (string, error) {

	for _, cookie := range cookies {
		parts := strings.Split(strings.TrimSpace(cookie.String()), ";")

		if len(parts) == 1 && parts[0] == "" {
			return "", errMalCookies
		}

		for _, part := range parts {
			part = strings.TrimSpace(part)
			j := strings.Index(part, "=")
			if j < 0 {
				if part == filter {
					return "", errMalCookies
				}
				fmt.Println("name=", part)
				continue
			}
			name, value := part[:j], part[j+1:]
			if 0 == strings.Compare(name, filter) {
				return value, nil
			}
		}
	}
	return "", errCookiesNotAvailable
}

// CookieSave saves beegosessionID into .cookie.yaml .
//
// This function is called only in stage of login, and will reset the content of
// .cookie.yaml no matter whether it exists or not.
func CookieSave(beegosessionID string) error {

	var cookie Beegocookie
	cookie.BeegosessionID = beegosessionID

	c, err := yaml.Marshal(&cookie)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(secretfile, []byte(c), 0644); err != nil {
		return err
	}

	return nil
}

// CookieClean removes .cookie.yaml file entirely.
//
// This function is called only in stage of logout.
func CookieClean() {
	os.Remove(secretfile)
}

// CookieLoad loads beegosessionID from .cookie.yaml.
func CookieLoad() (*Beegocookie, error) {
	var cookie Beegocookie

	dataBytes, err := ioutil.ReadFile(secretfile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(dataBytes), &cookie)
	if err != nil {
		return nil, err
	}

	return &cookie, nil
}
