// this file will be used on unix-based machines
package main

import "os"

func getPathOfConfig() string {
	return os.ExpandEnv("$HOME/.config/doc-api/config.json")
}
