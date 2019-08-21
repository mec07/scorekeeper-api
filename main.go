package main

import (
	"github.com/mec07/rununtil"
	"github.com/mec07/scorekeeper-api/webserver"
)

func main() {
	rununtil.KillSignal(webserver.NewRunner())
}
