package ui

import tea "github.com/charmbracelet/bubbletea"

type Window interface {
	Update(tea.Msg) (Window, tea.Cmd)
	View() string
}

type WindowList struct {
	header *PaneHeader
	list   *PaneList
	footer *PaneFooter

	active Pane
}

func (w *WindowList) Update(msg tea.Msg) (Window, tea.Cmd) {
	return nil, nil
}

func (w *WindowList) View() string {
	return "hi"
}

type WindowView struct {
	header *PaneHeader
	list   *PaneList
	footer *PaneFooter

	active Pane
}

func (w *WindowView) Update(_ tea.Msg) (Window, tea.Cmd) {
	return nil, nil
}

func (w *WindowView) View() string {
	return ""
}
