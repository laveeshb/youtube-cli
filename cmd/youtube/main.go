package main

import (
	"os"

	"github.com/laveeshb/youtube-cli/pkg"
)

func main() {
	if err := pkg.Execute("youtube"); err != nil {
		os.Exit(1)
	}
}
