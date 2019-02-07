package utils

import "errors"

const Logo = `
  __ __   ____  ____   ____    ___   ____      __ ______  _
 |  |  | /    ||    \ |    \  /   \ |    \    /  ]      || |
 |  |  ||  o  ||  D  )|  o  )|     ||  D  )  /  /|      || |
 |  _  ||     ||    / |     ||  O  ||    /  /  / |_|  |_|| |___
 |  |  ||  _  ||    \ |  O  ||     ||    \ /   \_  |  |  |     |
 |  |  ||  |  ||  .  \|     ||     ||  .  \\     | |  |  |     |
 |__|__||__|__||__|\_||_____| \___/ |__|\_| \____| |__|  |_____|
`

const Mark = `+----------------------+------------------------------------------+`

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
