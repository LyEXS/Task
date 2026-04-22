package main

import (
	"log"
	"lyes/task/config"
	"lyes/task/git"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	git.HandleGit(config)
}
