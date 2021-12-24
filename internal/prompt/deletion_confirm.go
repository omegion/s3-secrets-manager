package prompt

import (
	"github.com/manifoldco/promptui"
)

func (p Prompt) DeletionConfirm() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Delete Secret",
		IsConfirm: true,
	}

	return prompt.Run()
}
