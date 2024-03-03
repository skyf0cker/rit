package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/skyf0cker/rit/ui"
)

type BrowseCmd struct {
	Name string
}

func (b *BrowseCmd) Run() error {
	program := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}

	return nil
}
