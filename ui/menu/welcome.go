package menu

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Welcome() (string, error) {

	ShowWelcome()

	fmt.Println("\nEnter the desktop directory path")

	prompt := promptui.Prompt{
		Label:       "",
		Validate:    Valid,
		HideEntered: true,
	}

	return prompt.Run()
}