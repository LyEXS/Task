package main

import (
	"fmt"
	"log"
	"lyes/task/input"
)

func main() {

	commit_infos, err := input.GetUserInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commit_infos)
}
