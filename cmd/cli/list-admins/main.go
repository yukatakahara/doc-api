package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/oren/doc-api"
)

func main() {
	Admin, err := admin.New()
	if err != nil {
		panic(err)
	}

	results := Admin.All()
	data, err := json.Marshal(results)

	if err != nil {
		fmt.Errorf("encode response: %v", err)
	}

	os.Stdout.Write(data)
}
