package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	// select which page should display
	mode *ModeModel

	home *HomeModel
}

func NewApp() *App {
	return &App{
		mode: NewMode([]*Mode{
			_homeMode,
			_subMode,
		}),
		home: &HomeModel{},
	}
}

func (m *App) Init() tea.Cmd {
	return nil
}

func (m *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	mode := m.mode.GetSelected()
	if mode == nil {
		var cmd tea.Cmd
		m.mode, cmd = m.mode.Update(msg)
		return m, cmd
	}

	switch mode.Name {
	case "home":
		m.home, _ = m.home.Update(msg)
	case "subs":
		panic("not implement")
	default:
		panic("invalid mode selected, this is a bug of rit")
	}

	return m, nil
}

func (m *App) View() string {
	mode := m.mode.GetSelected()
	if mode == nil {
		return m.mode.View()
	}

	switch mode.Name {
	case "home":
		return m.home.View()
	case "subs":
		panic("not implement")
	default:
		panic("invalid mode selected, this is a bug of rit")
	}
}
