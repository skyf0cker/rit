package ui

import tea "github.com/charmbracelet/bubbletea"

type ActivateMsg string

func Activate(name string) tea.Cmd {
	return func() tea.Msg {
		return ActivateMsg(name)
	}
}
