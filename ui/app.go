package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list *WindowList
	view *WindowView

	active Window
}

func NewModel() *Model {
	model := Model{
		list: NewWindowList(),
		view: NewWindowView(),
	}

	model.active = model.list
	return &model
}

func (m *Model) Init() tea.Cmd {
	return tea.Sequence(
		Activate("list"),
		List("home"),
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, nil
		case "q", "ctrl+d":
			return m, tea.Quit
		}
	case ActivateMsg:
		switch msg {
		case "list":
			m.active = m.list
		case "view":
			m.active = m.view
		}
	case tea.WindowSizeMsg:
		var cmds []tea.Cmd
		for _, window := range []Window{m.list, m.view} {
			_, cmd := window.Update(msg)
			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)
	}

	var cmd tea.Cmd
	m.active, cmd = m.active.Update(msg)
	return m, cmd

}

func (m *Model) View() string {
	return m.active.View()
}
