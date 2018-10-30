package utils

import "errors"

// These variables are populated via the Go linker.
var (
	UTCBuildTime  = "unknown"
	ClientVersion = "unknown"
	GoVersion     = "unknown"
	GitBranch     = "unknown"
	GitTag        = "unknown"
	GitHash       = "unknown"
)

var errMalCookies = errors.New("get malformed cookies")
var errCookiesNotAvailable = errors.New("target cookies are not available")

var configfile = "conf/config.yaml"
var secretfile = "conf/.cookie.yaml"
