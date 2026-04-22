package tests

import (
	"lyes/task/formatters"
	"lyes/task/input"
	"regexp"
	"testing"
)

func TestFormatting(t *testing.T) {
	input := input.CommitInfo{
		CommitType:  "feat",
		CommitTitle: "added a X feature",
		CommitDesc:  "this is a test description",
	}

	got := formatters.FormatCommitInfos(input)

	expected := regexp.MustCompile("(?s)^" + regexp.QuoteMeta(input.CommitType+": "+input.CommitTitle) + ".*" + regexp.QuoteMeta(input.CommitDesc))
	if !expected.MatchString(got) {
		t.Errorf(`TestFormatting(data) =%q,want match for %#q, nil`, got, expected)
	}
}
