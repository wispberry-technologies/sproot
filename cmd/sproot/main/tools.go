package main

import (
	"fmt"
	"os/exec"
)

type tool struct {
	name        string
	description string
	install     func() error
	setup       func() error
	run         func(args []string) error
}

type Tool interface {
	Install() error
	Setup() error
	Run(args []string) error
}

func (t tool) Install() error {
	return t.install()
}
func (t tool) Setup() error {
	return t.setup()
}
func (t tool) Run(args []string) error {
	return t.run(args)
}

// ---------------
// Tool
// ---------------
var Crush Tool = tool{
	name:        "crush",
	description: "Your new coding bestie, now available in your favourite terminal. Your tools, your code, and your workflows, wired into your LLM of choice.",
	install: func() error {
		// check if the crush is installed using "crush -v"
		_, err := exec.Command("crushasd", "-v").Output()
		if err != nil {
			return fmt.Errorf("crush is not installed: %v", err)
		}
		return nil
	},
	setup: func() error {
		// Setup logic for Crush
		return nil
	},
	run: func(args []string) error {
		// Running logic for Crush
		return nil
	},
}
