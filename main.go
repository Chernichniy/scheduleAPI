package main

import (
	"os"

	"github.com/Chernichniy/scheduleAPI/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
