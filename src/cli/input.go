package cli

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/input"
	"github.com/cqroot/prompt/write"
	"github.com/fatih/color"
)

func GetCommitType() (string, error) {
	commitTypes := []choose.Choice{
		{Text: "feat", Note: "A new feature"},
		{Text: "fix", Note: "A bug fix"},
		{Text: "docs", Note: "Documentation changes"},
		{Text: "style", Note: "Code style changes (formatting, no logic)"},
		{Text: "refactor", Note: "Code refactoring (no fixes or new features)"},
		{Text: "perf", Note: "Performance improvements"},
		{Text: "test", Note: "Adding or updating tests"},
		{Text: "chore", Note: "Other changes (e.g., build or tool updates)"},
		{Text: "build", Note: "Changes that affect the build system or dependencies"},
		{Text: "ci", Note: "Continuous integration related changes"},
		{Text: "revert", Note: "Reverts a previous commit"},
	}

	commitType, err := prompt.New().
		Ask("What type of commit would you like to make?").
		AdvancedChoose(commitTypes, choose.WithTheme(choose.ThemeArrow))

	if err != nil {
		return "", err
	}

	return commitType, nil
}

func GetCommitScope() (string, error) {
	// dev only
	var savedScopes = []string{
		"frontend",
		"backend",
	}

	savedScopesStr := ""
	for i, savedScope := range savedScopes {
		if i == 0 {
			savedScopesStr += fmt.Sprintf("%s, ", savedScope)
		} else if i == len(savedScopes)-1 {
			savedScopesStr += fmt.Sprintf(" %s", savedScope)
		} else {
			savedScopesStr += fmt.Sprintf(" %s,", savedScope)
		}
	}

	wantsScope, err := prompt.New().Ask("Do you want to put a scope on the commit?").AdvancedChoose([]choose.Choice{
		{Text: "Yes, from my saves.", Note: savedScopesStr},
		{Text: "Yes, I'm going to create a new one."},
		{Text: "No."},
	})

	if err != nil {
		return "", err
	}

	if wantsScope == "Yes, from my saves." {
		// TODO: Obtain saved scopes
	} else if wantsScope == "Yes, I'm going to create a new one." {
		newScope, err := prompt.New().Ask("Write the new scope:").Input("")
		if err != nil {
			return "", err
		}

		return newScope, nil
	}

	return "", nil
}

func GetCommitTitle() (string, error) {
	commitTitle, err := prompt.New().Ask("Write a short title for the commit").Input("")

	if err != nil {
		return "", err
	}

	return commitTitle, nil
}

func GetCommitBody() (string, error) {
	commitBody, err := prompt.New().Ask("Do you want to add a detailed description of the change?").Write("", write.WithHelp(true))

	if err != nil {
		return "", err
	}

	return commitBody, nil
}

func GetCommitTicket() (string, string, error) {
	commitRef, err := prompt.New().Ask("Is this commit related to an issue or ticket? (example: #123)").Input("", input.WithValidateFunc(func(input string) error {
		validReference := regexp.MustCompile(`^#\d+$`)
		if input != "" && !validReference.MatchString(input) {
			return errors.New("invalid issue/ticket reference format (use #123)")
		}

		if input == "" {
			return nil
		}

		return nil
	}))
	if err != nil {
		return "", "", err
	}

	var wordRef = ""

	if commitRef != "" {
		wordRef, err = prompt.New().Ask("What type of relationship does this commit have with the issue/ticket?").
			AdvancedChoose([]choose.Choice{
				{Text: "closes", Note: "Automatically closes the issue when merged."},
				{Text: "fixes", Note: "Indicates the commit fixes the issue."},
				{Text: "resolves", Note: "Marks the issue as resolved."},
				{Text: "related to", Note: "Links the commit to the issue without closing it."},
				{Text: "partially fixes", Note: "Shows partial progress towards resolving the issue."},
			})

		if err != nil {
			return "", "", err
		}
	}

	return commitRef, wordRef, nil
}

func GetCommitEmoji() (bool, error) {
	commitEmoji, err := prompt.New().Ask("Do you want to add the emoji?").Choose(
		[]string{"Yes", "No"},
		choose.WithTheme(choose.ThemeLine),
		choose.WithKeyMap(choose.HorizontalKeyMap),
	)

	/*
		blue := color.New(color.FgBlue).SprintFunc()
		fmt.Printf("\n\n%s%s", blue("It would look like this:"), "")
	*/

	if commitEmoji == "No" || err != nil {
		return false, nil
	}

	return true, nil
}

func GetCommitBreakingChange() (bool, string, error) {
	commitBreaking, err := prompt.New().Ask("Does this commit introduce a breaking change?").Choose(
		[]string{"Yes", "No"},
		choose.WithTheme(choose.ThemeLine),
		choose.WithKeyMap(choose.HorizontalKeyMap),
	)

	if err != nil {
		return false, "", err
	}

	if commitBreaking == "No" {
		return false, "", nil
	}

	details, err := prompt.New().Ask("Describe the breaking change and any migration steps").Write("")

	if err != nil {
		return false, "", err
	}

	return true, details, nil
}

func GetCommitWIP() (bool, error) {
	commitWIP, err := prompt.New().Ask("Is this commit a work in progress (WIP)?").Choose(
		[]string{"Yes", "No"},
		choose.WithTheme(choose.ThemeLine),
		choose.WithKeyMap(choose.HorizontalKeyMap),
	)

	if commitWIP == "No" || err != nil {
		return false, nil
	}

	return true, nil
}

func ConfirmCommitMessage(commitMessage string) (bool, error) {
	lines := strings.Split(commitMessage, "\n")
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	line := color.New(color.FgCyan).Sprint(strings.Repeat("-", maxLength))

	fmt.Printf("\n%s\n\n%s\n\n%s\n\n", line, commitMessage, line)

	confirm, err := prompt.New().Ask("Are you sure you want to commit with this message?").Choose(
		[]string{"Yes", "No"},
		choose.WithTheme(choose.ThemeLine),
		choose.WithKeyMap(choose.HorizontalKeyMap),
	)
	if err != nil {
		return false, err
	}

	return confirm == "Yes", nil
}

func ConfirmStageChanges() (bool, error) {
	confirm, err := prompt.New().Ask("Do you want to stage the changes?").Choose(
		[]string{"Yes", "No"},
		choose.WithTheme(choose.ThemeLine),
		choose.WithKeyMap(choose.HorizontalKeyMap),
	)

	if err != nil {
		return false, err
	}

	return confirm == "Yes", nil
}
