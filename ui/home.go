package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type HomeModel struct {
	posts list.Model
}

func (hm *HomeModel) Init() tea.Cmd {
	return nil
}

func (hm *HomeModel) Update(msg tea.Msg) (*HomeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return hm, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		hm.posts.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	hm.posts, cmd = hm.posts.Update(msg)
	return hm, cmd
}

func (h *HomeModel) View() string {
	return "homepage"
}
