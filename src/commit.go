package main

import "fmt"

type CommitMessage struct {
	emoji bool
	_type string
	scope string
	title string
	wip   bool

	body      string
	ticketRef string
	wordRef   string

	breakingChange      bool
	breakingDescription string
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

	if commit.emoji {
		commitMessage += fmt.Sprintf("%s ", emojiTypes[commit._type])
	}

	commitMessage += commit._type

	if commit.scope != "" {
		commitMessage += fmt.Sprintf("(%s): ", commit.scope)
	} else {
		commitMessage += ": "
	}

	commitMessage += commit.title

	if commit.wip {
		commitMessage += " [WIP]"
	}

	if commit.body != "" {
		commitMessage += fmt.Sprintf("\n\n%s", commit.body)
	}

	if commit.ticketRef != "" {
		commitMessage += fmt.Sprintf("\n\nRelated to: %s", commit.ticketRef)
		if commit.wordRef != "" {
			commitMessage += fmt.Sprintf(" (%s)", commit.wordRef)
		}
	}

	if commit.breakingChange {
		commitMessage += fmt.Sprintf("\n\nBREAKING CHANGE: %s", commit.breakingDescription)
	}

	return commitMessage
}
