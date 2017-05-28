// this file will be used on windows-based machines
package config

import (
	"os"
	"path/filepath"
)

func GetPathOfConfig() string {
	// TODO: test and fix on windows
	// C:\Users\<your user name>\AppData\Local\doc-cli
	// also try os.Getenv("AppData")
	return os.ExpandEnv(filepath.FromSlash("$APPDATA/Local/doc-cli"))
}
