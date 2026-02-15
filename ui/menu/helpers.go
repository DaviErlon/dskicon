package menu

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const (
	dskicon = "\x1b[32m" + `
▓█████▄   ██████  ██ ▄█▀ ██▓ ▄████▄   ▒█████   ███▄    █ 
▒██▀ ██▌▒██    ▒  ██▄█▒ ▓██▒▒██▀ ▀█  ▒██▒  ██▒ ██ ▀█   █ 
░██   █▌░ ▓██▄   ▓███▄░ ▒██▒▒▓█    ▄ ▒██░  ██▒▓██  ▀█ ██▒
░▓█▄   ▌  ▒   ██▒▓██ █▄ ░██░▒▓▓▄ ▄██▒▒██   ██░▓██▒  ▐▌██▒
░▒████▓ ▒██████▒▒▒██▒ █▄░██░▒ ▓███▀ ░░ ████▓▒░▒██░   ▓██░
 ▒▒▓  ▒ ▒ ▒▓▒ ▒ ░▒ ▒▒ ▓▒░▓  ░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒ 
 ░ ▒  ▒ ░ ░▒  ░ ░░ ░▒ ▒░ ▒ ░  ░  ▒     ░ ▒ ▒░ ░ ░    ░ ▒░
   ░          ░     ░    ▒  ░        ░   ░       ░     ░ 
` + "\x1b[0m"
	welcome = "\x1b[32m" + `
 █░   ░█░▓█████  ██▓     ▄████▄   ▒█████   ███▄ ▄███▓▓█████ 
▓█▒ █ ░█░▓█   ▀ ▓██▒    ▒██▀ ▀█  ▒██▒  ██▒▓██▒▀█▀ ██▒▓█   ▀ 
▒█▓ █ ▒█ ▒███   ▒██░    ▒▓█    ▄ ▒██░  ██▒▓██    ▓██░▒███   
░██░█░▓█ ▒▓█  ▄ ▒██░    ▒▓█▄ ▄██▒▒██   ██░▒██    ▒██ ▒▓█  ▄ 
░░██▒██▓ ░▒████▒░██████▒▒ ▓███▀ ░░ ████▓▒░▒██▒   ░██▒░▒████▒
░ ▓░▒ ▒  ░░ ▒░ ░░ ▒░▓  ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ░  ░░░ ▒░ ░
  ░ ░ ░   ░ ░  ░░ ░ ▒  ░  ░  ░     ░ ▒ ▒░ ░  ░      ░ ░ ░  ░
    ░       ░              ░          ░      ░   	░
` + "\x1b[0m"
)

const (
	BASE = iota
	ADD
	ICONS
	CONFIG
	MDF_DESKTOP
	SEARCHDIR
	MDF_SEARCHDIR
	EXIT
)

type State = int

type Options struct {
	Name    string
	Details string
}

func ShowBanner() {
	ClearScreen()
	fmt.Print(dskicon)
}

func ShowWelcome() {
	ClearScreen()
	fmt.Print(welcome)
}

func ClearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Print("\033[3J")
}

func Valid(input string) error {
	if input == "" {
		return errors.New("Directory cannot be empty")
	}
	return nil
}

func Valid2(input string) error {
	if input == "" {
		return errors.New("Name cannot be empty")
	}
	return nil
}

func RemoveNonPrintable(s string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsPrint(r) {
			return -1
		}
		return r
	}, s)
}


