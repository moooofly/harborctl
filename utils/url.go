package utils

import "github.com/spf13/viper"

// URLGen generates target URL.
func URLGen(uri string) string {
	return viper.GetString("scheme") + "://" + viper.GetString("address") + uri
}
