[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/moooofly/harborctl) [![Build Status](https://travis-ci.org/moooofly/harborctl.svg?branch=master)](https://travis-ci.org/moooofly/harborctl) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/moooofly/harborctl/blob/master/LICENSE)

# harborctl

```
  __ __   ____  ____   ____    ___   ____      __ ______  _
 |  |  | /    ||    \ |    \  /   \ |    \    /  ]      || |
 |  |  ||  o  ||  D  )|  o  )|     ||  D  )  /  /|      || |
 |  _  ||     ||    / |     ||  O  ||    /  /  / |_|  |_|| |___
 |  |  ||  _  ||    \ |  O  ||     ||    \ /   \_  |  |  |     |
 |  |  ||  |  ||  .  \|     ||     ||  .  \\     | |  |  |     |
 |__|__||__|__||__|\_||_____| \___/ |__|\_| \____| |__|  |_____|
```

A CLI tool for the Docker Registry Harbor.

This project offer a command-line interface to the Harbor API, you can use it to manager your Harbor service as from Harbor UI.

## NOTE

- This project named [`harborctl`](https://github.com/moooofly/harborctl) is still under developement, which is based on harbor v1.6.0-66709daa and swagger api version 1.6.0.
- Another project named `harbor-go-client` is based on harbor v1.5.0-d59c257e and swagger api version 1.4.0. If you want to use this CLI tool with Harbro v1.6.0+, you may encounter incompatible issues. See "[The issue with API version](https://github.com/moooofly/harbor-go-client/issues/27)" for more details.

## Features

Current Harbor API support status, see [issue/5](https://github.com/moooofly/harborctl/issues/5):

## Installation


Assuming you already have a recent version of Go installed, pull down the code with go get:

```
go get -u github.com/moooofly/harborctl
```

You can also obtain the pre-compiled binary files from [here](https://github.com/moooofly/harborctl/releases), which is recommanded.

## Documentation

See [docs](https://github.com/moooofly/harborctl/tree/master/docs)

## Credits

- [parnurzeal/gorequest](https://github.com/parnurzeal/gorequest) - Simplified HTTP client (inspired by famous SuperAgent lib in Node.js)
- [spf13/cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions.
- [spf13/viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [go-yaml/yaml](https://github.com/go-yaml/yaml) - YAML support for the Go language.

## License

harborctl is licensed under the MIT License. See [LICENSE](https://github.com/moooofly/harborctl/blob/master/LICENSE) for the full license text.

This project uses open source components which have additional licensing terms. The licensing terms for these open source components can be found at the following locations:

- gorequest: [license](https://github.com/parnurzeal/gorequest/blob/master/LICENSE)
- cobra: [license](https://github.com/spf13/cobra/blob/master/LICENSE.txt)
- viper: [license](https://github.com/spf13/viper/blob/master/LICENSE)
- yaml: [license](https://github.com/go-yaml/yaml/blob/v2/LICENSE)
