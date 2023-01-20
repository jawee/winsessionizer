package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type folder struct {
    path string
}

type model struct {
    choices []folder
    filteredChoices []folder
    searchBox string
    cursor int
}

func initialModel() *model {
    folders := []folder{{"/asdf/asdf/carrots"}, {"/asdf/asdf/celery"}, {"/asdf/asdf/cucumber"}}
    return &model{
        choices: folders,
        filteredChoices: folders,
    }
}

func (m *model) Init() tea.Cmd {
    return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit

        case "up":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        case "enter":
            //we're done. cd to selected one
        default: 
            m.searchBox += msg.String()
            m.filter()
        }
    }

    return m, nil
}

func (m *model) filter() {
    filterStr := m.searchBox

    newArray := make([]folder, 0)
    for _, s := range m.filteredChoices {
        if strings.Contains(s.path, filterStr) {
            newArray = append(newArray, s)
        }
    }

    m.filteredChoices = newArray
}

func (m *model) View() string {
    s := ""
    for i, choice := range m.filteredChoices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        s += fmt.Sprintf("%s %s\n", cursor, choice.path)
    }

    s += fmt.Sprintf("Filter> %s", m.searchBox)

    return s
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
