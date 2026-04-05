package main

import (
	"os"

	"github.com/laveeshb/youtube-cli/pkg"
)

func main() {
	if err := pkg.Execute("yt"); err != nil {
		os.Exit(1)
	}
}
