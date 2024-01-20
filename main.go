package main

import (
	"fmt"
	"time"

	"github.com/msrevive/maphash/cmd"
)

var spMsg string = `
checksum

Copyright © %d, Team MSRebirth
Website: https://msrebirth.net/
License: GPL-3.0

`

func main() {
	fmt.Printf(spMsg, time.Now().Year())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
