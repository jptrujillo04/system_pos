package main

import (
	"os"
	"pos_system/cmd/api/app"
)

func main() {
	application := app.New()
	if application != nil {
		os.Exit(1)
	}
	return
}
