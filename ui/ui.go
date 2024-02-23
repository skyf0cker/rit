package ui

import tea "github.com/charmbracelet/bubbletea"

func Run() error {
	app := NewApp()
	program := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}

	return nil
}
