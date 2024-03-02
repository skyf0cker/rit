package ui

import tea "github.com/charmbracelet/bubbletea"

type Pane interface {
	Update(tea.Msg) (Pane, tea.Cmd)
	View() string
	Activate() Pane
	Deactivate()

	Size() (width, height int)
	SetSize(width, height int)
}

type PaneHeader struct{}

func (p *PaneHeader) Update(_ tea.Msg) (Pane, tea.Cmd) {
	return nil, nil
}

func (p *PaneHeader) View() string {
	return ""
}

func (p *PaneHeader) Activate() Pane {
	return nil
}

func (p *PaneHeader) Deactivate() {

}

func (p *PaneHeader) Size() (width int, height int) {
	return 0, 0
}

func (p *PaneHeader) SetSize(width int, height int) {
}

type PaneList struct{}

func (p *PaneList) Update(_ tea.Msg) (Pane, tea.Cmd) {
	return nil, nil
}

func (p *PaneList) View() string {
	return ""
}

func (p *PaneList) Activate() Pane {
	return nil
}

func (p *PaneList) Deactivate() {

}

func (p *PaneList) Size() (width int, height int) {
	return 0, 0
}

func (p *PaneList) SetSize(width int, height int) {
}

type PaneFooter struct{}

func (p *PaneFooter) Update(_ tea.Msg) (Pane, tea.Cmd) {
	return nil, nil
}

func (p *PaneFooter) View() string {
	return ""
}

func (p *PaneFooter) Activate() Pane {
	return nil
}

func (p *PaneFooter) Deactivate() {

}

func (p *PaneFooter) Size() (width int, height int) {
	return 0, 0
}

func (p *PaneFooter) SetSize(width int, height int) {
}

type PaneView struct {
}

func (p *PaneView) Update(_ tea.Msg) (Pane, tea.Cmd) {
	return nil, nil
}

func (p *PaneView) View() string {
	return ""
}

func (p *PaneView) Activate() Pane {
	return nil
}

func (p *PaneView) Deactivate() {

}

func (p *PaneView) Size() (width int, height int) {
	return 0, 0
}

func (p *PaneView) SetSize(width int, height int) {

}
