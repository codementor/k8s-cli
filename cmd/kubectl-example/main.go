package main

import (
	"os"

	"github.com/codementor/k8s-cli/pkg/example/cmd"
)

func main() {
	if err := cmd.NewExampleCmd().Execute(); err != nil {
		os.Exit(-1)
	}
}
