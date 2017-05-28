// this file will be used on unix-based machines
package config

import "os"

func GetPathOfConfig() string {
	return os.ExpandEnv("$HOME/.config/doc-api/config.json")
}
