package validators

import (
	"errors"
	"strings"
)

func ValidateType(commit_type string) error {
	if len(strings.Fields(commit_type)) == 1 {
		return nil
	}
	return errors.New("The type can only contain one and only one character")
}

func ValidateTitle(commit_title string) error {
	if len(commit_title) > 50 {
		return errors.New("The title can contain a maximum of 50 characters")
	}
	return nil
}
