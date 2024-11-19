package git

import "fmt"

type CommitMessage struct {
	Emoji bool
	Type string
	Scope string
	Title string
	Wip   bool

	Body      string
	TicketRef string
	WordRef   string

	BreakingChange      bool
	BreakingDescription string
}

var emojiTypes = map[string]string{
	"feat":     "✨",
	"fix":      "🐛",
	"docs":     "📚",
	"style":    "🎨",
	"refactor": "🔨",
	"perf":     "⚡",
	"test":     "✅",
	"chore":    "🔧",
	"build":    "👷",
	"ci":       "📦",
	"revert":   "⏪",
}

func BuildCommitMessage(commit CommitMessage) string {
	commitMessage := ""

	if commit.Emoji {
		commitMessage += fmt.Sprintf("%s ", emojiTypes[commit.Type])
	}

	commitMessage += commit.Type

	if commit.Scope != "" {
		commitMessage += fmt.Sprintf("(%s): ", commit.Scope)
	} else {
		commitMessage += ": "
	}

	commitMessage += commit.Title

	if commit.Wip {
		commitMessage += " [WIP]"
	}

	if commit.Body != "" {
		commitMessage += fmt.Sprintf("\n\n%s", commit.Body)
	}

	if commit.TicketRef != "" {
		commitMessage += fmt.Sprintf("\n\nRelated to: %s", commit.TicketRef)
		if commit.WordRef != "" {
			commitMessage += fmt.Sprintf(" (%s)", commit.WordRef)
		}
	}

	if commit.BreakingChange {
		commitMessage += fmt.Sprintf("\n\nBREAKING CHANGE: %s", commit.BreakingDescription)
	}

	return commitMessage
}
