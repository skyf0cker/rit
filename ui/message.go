package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ActivateMsg string

func Activate(name string) tea.Cmd {
	return func() tea.Msg {
		return ActivateMsg(name)
	}
}

type HeaderMsg int

func Header(n int) tea.Cmd {
	return func() tea.Msg {
		return HeaderMsg(n)
	}
}

type ListType interface {
	string | []*PostItem
}

type ListMsg[T ListType] struct {
	Value T
}

func List[T ListType](t T) tea.Cmd {
	return func() tea.Msg {
		return ListMsg[T]{
			Value: t,
		}
	}
}

type ViewType interface {
	*PostItem
}

type ViewMsg[T ViewType] struct {
	Value T
}

func View[T ViewType](t T) tea.Cmd {
	return func() tea.Msg {
		return ViewMsg[T]{
			Value: t,
		}
	}
}
