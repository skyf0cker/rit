package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	_homeMode = &Mode{Name: "home", Description: "reddit homepage"}
	_subMode  = &Mode{Name: "subs", Description: "subscribed subreddits"}

	_supportModes = map[string]*Mode{
		"home": _homeMode,
		"subs": _subMode,
	}
)

type Mode struct {
	Name        string
	Description string
}

func createModeItems(modes []*Mode) []list.Item {
	var items []list.Item
	for _, mode := range modes {
		items = append(items, &item{
			title:       mode.Name,
			description: mode.Description,
		})
	}

	return items
}

type ModeModel struct {
	supportModes list.Model

	selectedMode *Mode
}

func NewMode(modes []*Mode) *ModeModel {
	modeList := list.New(createModeItems(modes), list.NewDefaultDelegate(), 0, 0)
	modeList.Title = "select mode"
	return &ModeModel{
		supportModes: modeList,
	}
}

func (m ModeModel) Init() tea.Cmd {
	return nil
}

func (m *ModeModel) Update(msg tea.Msg) (*ModeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		} else if msg.Type == tea.KeyEnter {
			item := m.supportModes.SelectedItem()
			m.selectedMode = _supportModes[item.FilterValue()]
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.supportModes.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.supportModes, cmd = m.supportModes.Update(msg)
	return m, cmd
}

func (m ModeModel) View() string {
	return docStyle.Render(m.supportModes.View())
}

func (m ModeModel) GetSelected() *Mode {
	return m.selectedMode
}
