package main

import (
	"embed"

	"github.com/marcelofabianov/wisp-cli/cmd/wisp/cmd"
)

//go:embed all:content
var contentFS embed.FS

func main() {
	cmd.SetContentFS(contentFS)
	cmd.Execute()
}
