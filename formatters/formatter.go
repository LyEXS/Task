package formatters

import (
	"lyes/task/input"
	"strings"
)

const COMMIT_GUIDE = `# At the end: Include Co-authored-by for all contributors.
# Include at least one empty line before it. Format:
# Co-authored-by: name <user@users.noreply.github.com>
#
# How to Write a Git Commit Message:
# https://chris.beams.io/posts/git-commit/
#
# 1. Separate subject from body with a blank line
# 2. Limit the subject line to 50 characters
# 3. Capitalize the subject line
# 4. Do not end the subject line with a period
# 5. Use the imperative mood in the subject line
# 6. Wrap the body at 72 characters
# 7. Use the body to explain what and why vs. how
`

func FormatCommitInfos(infos input.CommitInfo) string {
	var format_builder strings.Builder

	format_builder.WriteString(infos.CommitType)
	format_builder.WriteString(": ")
	format_builder.WriteString(infos.CommitTitle)
	format_builder.WriteString("\n\n")
	format_builder.WriteString(infos.CommitDesc)
	format_builder.WriteString("\n\n\n")
	format_builder.WriteString(COMMIT_GUIDE)

	return format_builder.String()
}
