package input

import (
	"bufio"
	"errors"
	"fmt"
	"lyes/task/utils"
	"lyes/task/validators"
	"os"
	"os/exec"
	"strings"
)

type CommitInfo struct {
	CommitType  string
	CommitTitle string
	CommitDesc  string
}

func GetUserInput() (CommitInfo, error) {
	var commit_type, commit_title, commit_desc string

	scanner := bufio.NewScanner(os.Stdin)

	commit_type = getCommitType(*scanner)
	commit_title = getCommitTitle(*scanner)
	commit_desc, err := getCommitDescription()

	if err != nil {
		return CommitInfo{}, err
	}
	return CommitInfo{
		CommitType:  commit_type,
		CommitTitle: commit_title,
		CommitDesc:  commit_desc,
	}, nil
}

func getTextFromEditor(OPERATING_SYSTEM string) (string, error) {

	file, err := os.Create(".git/tmp_commit_msg.txt")

	if err != nil {
		return "", err
	}
	file.Close()
	var cmd *exec.Cmd
	switch OPERATING_SYSTEM {
	case "linux":
		cmd = exec.Command("nano", file.Name())
	case "windows":
		cmd = exec.Command("nano", file.Name())
	default:
		return "", errors.New("Operating System not handled")
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	content_byte, err := os.ReadFile(file.Name())
	if err != nil {
		return "", err
	}

	content := string(content_byte)

	os.Remove(file.Name())

	return content, err
}

func getCommitType(scanner bufio.Scanner) string {

	for {
		fmt.Print("Type of the commit (fix, feat, docs,...) :")
		if scanner.Scan() {
			if err := validators.ValidateType(scanner.Text()); err == nil {
				return scanner.Text()
			} else {
				fmt.Println(err)
			}
		}
	}
}

func getCommitTitle(scanner bufio.Scanner) string {

	for {
		fmt.Print("Title of the commit :")
		if scanner.Scan() {
			if err := validators.ValidateTitle(scanner.Text()); err != nil {
				return scanner.Text()
			} else {
				fmt.Println(err)
			}
		}
	}
}

func getCommitDescription() (string, error) {
	var is_desc string
	var commit_desc = ""
	for {
		fmt.Print("Do you want to specify description ? (y/n) ")
		fmt.Scan(&is_desc)
		if strings.TrimSpace(is_desc) == "y" || strings.TrimSpace(is_desc) == "n" {
			break
		}
	}

	switch strings.ToLower(strings.TrimSpace(is_desc)) {
	case "y":
		var err error
		commit_desc, err = getTextFromEditor(utils.GetOS())
		if err != nil {
			return "", err
		}
	case "n":
		commit_desc = ""
	}

	return commit_desc, nil
}
