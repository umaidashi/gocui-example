package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
	ID        int
	Name      string
	IsDone    bool
	CreatedAt time.Time
}

type model struct {
	cursor int
	tasks  []*Task
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if m.cursor < len(m.tasks) {
				m.cursor++
			}
		case "k":
			if m.cursor > 1 {
				m.cursor--
			}
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "--YOUR TASKS--\n\n"

	for i, v := range m.tasks {
		cursor := " "
		if i == m.cursor-1 {
			cursor = ">"
		}

		timeLayout := "2006-01-02 15:04"
		s += fmt.Sprintf("%s #%d %s (%s)\n", cursor, v.ID, v.Name, v.CreatedAt.Format(timeLayout))
	}

	s += "\nPress 'q' to quit\n"

	return s
}

func main() {
	m := model{
		cursor: 1,
		tasks: []*Task{
			{
				ID:        1,
				Name:      "First task!",
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Name:      "Write an article about bubbletea",
				CreatedAt: time.Now(),
			},
		},
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("app-name: %s", err.Error())
		os.Exit(1)
	}
}
