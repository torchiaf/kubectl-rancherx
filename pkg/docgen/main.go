package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/torchiaf/kubectl-rancherx/pkg/cli"
)

func main() {
	rootCmd, err := cli.NewRootCmd()
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTree(rootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
