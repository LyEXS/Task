package git

import (
	"log"
	"lyes/task/config"
	"lyes/task/formatters"
	"lyes/task/input"
	"lyes/task/utils"
	"os"
)

func HandleGit(cnf *config.Config) {
	infos, err := input.GetUserInput(cnf)
	if err != nil {
		log.Fatal(err)
	}

	formatted_content := formatters.FormatCommitInfos(infos)

	git_commit_file_path := utils.GetFilePath()
	err = os.WriteFile(git_commit_file_path, []byte(formatted_content), 0644)
	if err != nil {
		log.Fatalf("Erreur fatale lors de l'injection du message : %v", err)
	}

}
