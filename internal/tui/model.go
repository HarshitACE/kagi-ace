package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	count int
}

func NewModel() Model {
	return Model{count: 0}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.count++
		case "down", "j":
			m.count--
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	return tea.NewView(fmt.Sprintf(
		"Counter: %d\n\nPress ↑ or k to add 1.\nPress ↓ or j to subtract 1.\nPress q to quit.\n",
		m.count,
	))
}
