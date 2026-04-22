package input

import (
	"bufio"
	"errors"
	"fmt"
	"lyes/task/config"
	"lyes/task/utils"
	"lyes/task/validators"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
)

type CommitInfo struct {
	CommitType  string
	CommitTitle string
	CommitDesc  string
}

func GetUserInput(cfg *config.Config) (CommitInfo, error) {
	var commit_type, commit_title, commit_desc string

	scanner := bufio.NewScanner(os.Stdin)

	commit_type = getCommitType(cfg)
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

	file, err := os.Create(".git/tmp_commit_msg.md")

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

func getCommitType(cfg *config.Config) string {
	var commitType string

	var options []huh.Option[string]
	for _, ct := range cfg.CommitTypes {
		labelWithEmoji := fmt.Sprintf("%s %s", ct.Label, ct.Emoji)

		options = append(options, huh.NewOption(labelWithEmoji, ct.Label+" "+ct.Emoji))
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Options(options...).Value(&commitType),
		),
	)
	err := form.Run()
	if err != nil {
		return ""
	}
	return commitType

}

func getCommitTitle(scanner bufio.Scanner) string {

	for {
		fmt.Print("Title of the commit :")
		if scanner.Scan() {
			if err := validators.ValidateTitle(scanner.Text()); err == nil {
				return scanner.Text()
			} else {
				fmt.Println(err)
			}
		}
	}
}

func getCommitDescription() (string, error) {
	var wantsDesc bool
	var CommitDesc string = ""

	confirm := huh.NewConfirm().Title("Do you want to specify a description ?").Affirmative("Yes").Negative("No").Value(&wantsDesc)
	err := confirm.Run()
	if err != nil {
		return "", err
	}
	if wantsDesc {
		desc, err := getTextFromEditor(utils.GetOS())
		if err != nil {
			return "", err
		}
		CommitDesc = desc
	}
	return CommitDesc, nil
}
