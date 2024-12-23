package cli

import (
	"cass/src/config"
	"cass/src/git"
)

func initAction() (string, error) {
	if err := config.WriteConfigFile(); err != nil {
		return "", err
	}

	commit, err := runTUI()
	if err != nil {
		return "", err
	}

	commitMessage := git.BuildCommitMessage(commit)

	return commitMessage, nil
}

func runTUI() (git.CommitMessage, error) {
	commit := git.CommitMessage{}

	println("")

	commitType, err := GetCommitType()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.Type = commitType

	println("")

	scope, err := GetCommitScope()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.Scope = scope

	println("")

	title, err := GetCommitTitle()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.Title = title

	println("")

	body, err := GetCommitBody()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.Body = body

	println("")

	ticketRef, wordRef, err := GetCommitTicket()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.TicketRef = ticketRef
	commit.WordRef = wordRef

	println("")

	useEmoji, err := GetCommitEmoji()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.Emoji = useEmoji

	println("")

	breakingChange, breakingDescription, err := GetCommitBreakingChange()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.BreakingChange = breakingChange
	commit.BreakingDescription = breakingDescription

	println("")

	wip, err := GetCommitWIP()
	if err != nil {
		return git.CommitMessage{}, err
	}
	commit.Wip = wip

	println("")

	return commit, nil
}
