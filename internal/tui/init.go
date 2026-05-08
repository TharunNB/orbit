package tui

import (
	"fmt"
	"orbit/internal/models"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	models   []models.Model
	selected int
}

func InitialModel() model {
	return model{
		models: models.AvailableModels,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.models)-1 {
				m.cursor++
			}

		case "enter":
			m.selected = m.cursor
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Select a model:\n\n"

	for i, choice := range m.models {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf(
			"%s %s (%s)\n",
			cursor,
			choice.Name,
			choice.Size,
		)
	}

	s += "\nPress q to quit.\n"

	return s
}

func Start() models.Model {
	p := tea.NewProgram(InitialModel())

	finalModel, err := p.Run()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	m := finalModel.(model)

	return m.models[m.selected]
}
