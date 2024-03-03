package ui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/skyf0cker/rit/reddit"
)

type Pane interface {
	Update(tea.Msg) (Pane, tea.Cmd)
	View() string
	Activate() Pane
	Deactivate()

	Size() (width, height int)
	SetSize(width, height int)
}

type PaneList struct {
	model list.Model
	style lipgloss.Style
}

func NewPaneList() *PaneList {
	color := lipgloss.Color("#ff6600")
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(color).BorderLeftForeground(color)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle.Copy().Faint(true)

	model := list.New([]list.Item{}, delegate, 0, 0)
	model.SetShowHelp(false)
	model.SetShowStatusBar(false)
	model.SetShowTitle(false)
	model.SetShowPagination(false)
	return &PaneList{
		model: model,
		style: lipgloss.NewStyle().Margin(1, 2),
	}
}

func (p *PaneList) Update(msg tea.Msg) (Pane, tea.Cmd) {
	rd := reddit.GetReddit()
	switch msg := msg.(type) {
	case ListMsg[string]:
		switch strings.ToLower(msg.Value) {
		case "home":
			return p, func() tea.Msg {
				posts, err := rd.GetHomePage()
				if err != nil {
					return err
				}

				var items []*PostItem
				for _, post := range posts {
					items = append(items, &PostItem{post: post})
				}

				return ListMsg[[]*PostItem]{
					Value: items,
				}
			}
		default:
			return p, nil
		}
	case ListMsg[[]*PostItem]:
		items := msg.Value
		sort.Slice(items, func(i, j int) bool {
			return items[i].post.CreatedUtc > items[j].post.CreatedUtc
		})

		var modelItems []list.Item
		for _, item := range items {
			modelItems = append(modelItems, item)
		}

		return p, p.model.SetItems(modelItems)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			post := p.model.SelectedItem().(*PostItem)
			return p, tea.Sequence(
				Activate("view"),
				View(post),
			)
		case "k", "up":
			if p.model.Index() == 0 {
				return p, Activate("header")
			}
		case "tab":
			return p, Activate("toggle")
		}
	}

	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd

}

func (p *PaneList) View() string {
	return p.style.Render(p.model.View())
}

func (p *PaneList) Activate() Pane {
	return p
}

func (p *PaneList) Deactivate() {

}

func (p *PaneList) Size() (width, height int) {
	h, v := p.style.GetFrameSize()
	return p.model.Width() + h, p.model.Height() + v
}

func (p *PaneList) SetSize(width, height int) {
	h, v := p.style.GetFrameSize()
	p.model.SetSize(width-h, height-v)
}

type PaneView struct {
	*PostItem
	style lipgloss.Style
	model viewport.Model

	content strings.Builder

	styleTitle        lipgloss.Style
	styleDescription  lipgloss.Style
	styleCommentTitle lipgloss.Style
	styleOP           lipgloss.Style
}

func NewPaneView() *PaneView {
	return &PaneView{
		style: lipgloss.NewStyle().Margin(1, 2),
		model: viewport.New(0, 0),
		styleTitle: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}),
		styleDescription: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#a49fa5", Dark: "#777777"}),
		styleCommentTitle: lipgloss.NewStyle().Foreground(lipgloss.Color("#ff6600")),
		styleOP:           lipgloss.NewStyle().Foreground(lipgloss.Color("#0099ff")).SetString("OP"),
	}
}

func (p *PaneView) Update(msg tea.Msg) (Pane, tea.Cmd) {
	switch msg := msg.(type) {
	case ViewMsg[*PostItem]:
		p.PostItem = msg.Value
		p.Render()
		// get comments here
		return p, nil
	// case ViewMsg[*Comment]:
	// 	p.Render()
	// 	return p, tea.Batch(comments(msg.Value.Item)...)
	case tea.KeyMsg:
		switch msg.String() {
		case "k", "up":
			if p.model.AtTop() {
				return p, Activate("header")
			}
		case "g", "home":
			p.model.GotoTop()
		case "G", "end":
			p.model.GotoBottom()
		case "tab":
			return p, Activate("toggle")
		}
	case tea.WindowSizeMsg:
		p.Render()
	}

	var cmd tea.Cmd
	p.model, cmd = p.model.Update(msg)
	return p, cmd
}

func (p *PaneView) View() string {
	return p.style.Render(p.model.View())
}

func (p *PaneView) Render() {
	p.content.Reset()
	if s := p.PostItem; s != nil {
		title := s.post.Title
		fmt.Fprintln(&p.content, p.styleTitle.Render(title))

		// description := strings.TrimSpace(s.post.Selftext)
		// fmt.Fprintln(&p.content, p.styleDescription.Render(description))

		if s.post.SelftextHTML != "" {
			fmt.Fprintln(&p.content, p.styleDescription.Copy().MarginTop(1).Width(p.style.GetWidth()).Render(HTMLText(s.post.SelftextHTML)))
		}

		// styleComment := lipgloss.NewStyle().
		// 	Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		// 	Border(lipgloss.NormalBorder(), false).
		// 	BorderLeft(true).
		// 	PaddingLeft(1).
		// 	MarginTop(1)
		//
		// h, _ := styleComment.GetFrameSize()
		//
		// var view func(lipgloss.Style, []*Comment) string
		// view = func(style lipgloss.Style, comments []*Comment) string {
		// 	var lines []string
		// 	for _, comment := range comments {
		// 		if comment.By != "" {
		// 			var sb strings.Builder
		//
		// 			by := p.styleCommentTitle.Render(comment.By)
		// 			if comment.By == s.By {
		// 				by = fmt.Sprintf("%s %s", by, p.styleOP.String())
		// 			}
		//
		// 			fmt.Fprintln(&sb, by, p.styleCommentTitle.Copy().Faint(true).Render(humanize(time.Unix(comment.Time, 0))))
		//
		// 			sb.WriteString(HTMLText(comment.Text))
		//
		// 			if len(comment.Comments) > 0 {
		// 				fmt.Fprintln(&sb)
		// 				fmt.Fprint(&sb, view(style.Copy().Width(style.GetWidth()-h), comment.Comments))
		// 			}
		//
		// 			lines = append(lines, style.Render(sb.String()))
		// 		}
		// 	}
		//
		// 	return strings.Join(lines, "\n")
		// }
		//
		// fmt.Fprintln(&p.content, view(styleComment.Copy().Width(p.style.GetWidth()-h), s.Comments))
	}

	p.model.SetContent(p.content.String())
}

func (p *PaneView) Size() (width, height int) {
	h, v := p.style.GetFrameSize()
	return p.style.GetWidth() + h, p.style.GetHeight() + v
}

func (p *PaneView) SetSize(width, height int) {
	h, v := p.style.GetFrameSize()
	p.style = p.style.Width(width - h).Height(height - v)
	p.model.Width, p.model.Height = width-h, height-v
}

func (p *PaneView) Activate() Pane {
	p.model.GotoTop()
	return p
}

func (p *PaneView) Deactivate() {
}

type PaneHeader struct {
	index         int
	width, height int
	active        bool

	style lipgloss.Style
	items []lipgloss.Style
	funcs []func() tea.Cmd
}

type PaneHeaderItem struct {
	Name string
	Func func() tea.Cmd
}

func NewPaneHeader(items ...PaneHeaderItem) *PaneHeader {
	pane := PaneHeader{
		style: lipgloss.NewStyle().Margin(1, 2),
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"})

	for _, item := range items {
		pane.items = append(pane.items, style.Copy().SetString(item.Name))
		pane.funcs = append(pane.funcs, item.Func)
	}

	return &pane
}

func (p *PaneHeader) Update(msg tea.Msg) (Pane, tea.Cmd) {
	switch msg := msg.(type) {
	case HeaderMsg:
		p.index = int(msg)
		return p, p.funcs[int(msg)]()
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return p, Header(p.index)
		case "h", "left":
			p.index = mod(p.index-1, len(p.items))
		case "l", "right":
			p.index = mod(p.index+1, len(p.items))
		case "j", "down", "tab":
			return p, Activate("toggle")
		}
	}

	return p, nil
}

func (p *PaneHeader) View() string {
	var views []string
	for i := range p.items {
		state := p.items[i]
		if i == p.index {
			state = state.Copy().Underline(true)
			if p.active {
				state = state.Copy().Foreground(lipgloss.Color("#ff6600"))
			}
		}

		if i < len(p.items)-1 {
			state = state.Copy().MarginRight(2)
		}

		views = append(views, state.String())
	}

	var sb strings.Builder

	left := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff6600")).
		Bold(true).
		Render("rit")
	right := lipgloss.JoinHorizontal(lipgloss.Top, views...)

	sb.WriteString(left)
	if pad := p.width - lipgloss.Width(left) - lipgloss.Width(right); pad > 0 {
		sb.WriteString(strings.Repeat(" ", pad))
	}

	sb.WriteString(right)
	return p.style.Render(sb.String())
}

func (p *PaneHeader) Size() (width, height int) {
	_, v := p.style.GetFrameSize()
	return 0, v + 1
}

func (p *PaneHeader) SetSize(width, height int) {
	h, v := p.style.GetFrameSize()
	p.width, p.height = width-h, height-v
}

func (p *PaneHeader) Activate() Pane {
	p.active = true
	return p
}

func (p *PaneHeader) Deactivate() {
	p.active = false
}

type PaneFooter struct {
	width, height int
	style         lipgloss.Style
	left, right   func() string
}

func NewPaneFooter(left, right func() string) *PaneFooter {
	return &PaneFooter{
		left:  left,
		right: right,
		style: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#a49fa5", Dark: "#777777"}).
			Faint(true).
			Margin(1, 2, 0),
	}
}

func (p *PaneFooter) Update(tea.Msg) (Pane, tea.Cmd) {
	return p, nil
}

func (p *PaneFooter) View() string {
	var sb strings.Builder

	left := p.left()
	right := p.right()

	sb.WriteString(left)
	if fill := p.width - lipgloss.Width(left) - lipgloss.Width(right); fill > 0 {
		sb.WriteString(strings.Repeat(" ", fill))
	}

	sb.WriteString(right)
	return p.style.Render(sb.String())
}

func (p *PaneFooter) Size() (width, height int) {
	_, v := p.style.GetFrameSize()
	return 0, v + 1
}

func (p *PaneFooter) SetSize(width, height int) {
	h, v := p.style.GetFrameSize()
	p.width, p.height = width-h, height-v
}

func (p *PaneFooter) Activate() Pane {
	return p
}

func (p *PaneFooter) Deactivate() {
}

func mod(a, b int) int {
	return (a%b + b) % b
}
