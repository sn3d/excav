package termui

import "github.com/AlecAivazis/survey/v2"

func ConfirmApply() bool {
	return Confirm("Apply patch to these repositories?", true)
}

// Confirm ask you question and return true if the answer is
// 'yes' or false if answer is 'no'. You need to set also default
// value.
func Confirm(question string, deflt bool) bool {
	result := false
	prompt := &survey.Confirm{
		Message: question,
		Default: deflt,
	}
	survey.AskOne(prompt, &result)
	return result
}
