package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list *WindowList
	view *WindowView

	active Window
}

func (m *Model) Init() tea.Cmd {
	return tea.Sequence(
		Activate("list"),
	)
}

func (m *Model) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m *Model) View() string {
	return "Hello"
}
