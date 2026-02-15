package main

import (
	"dskicon/ui"
	"dskicon/util"
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {

	app, err := util.LoadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = ui.Run(app); err != promptui.ErrInterrupt && err != promptui.ErrEOF && err != nil {
		fmt.Println(err)
	}
}
