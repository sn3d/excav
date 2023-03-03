package termui

import (
	"github.com/AlecAivazis/survey/v2"
)

func Ask(question string) string {
	answer := ""
	prompt := &survey.Input{
		Message: question,
	}
	survey.AskOne(prompt, &answer)
	return answer
}

func Select(question string, options ...string) string {
	answer := ""
	prompt := &survey.Select{
		Message: question,
		Options: options,
	}

	survey.AskOne(prompt, &answer)
	return answer
}