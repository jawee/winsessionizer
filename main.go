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
    choice folder
}

func initialModel() model {
    dirs := []string{"/home/figge/go/src/github.com/jawee","/home/figge/projects"}
    folders := make([]folder,0)
    for _, dir := range dirs {
        files, _ := os.ReadDir(dir)
        for _, f := range files {
            if f.IsDir() {
                path := fmt.Sprintf("%s/%s", dir, f.Name())
                folders = append(folders, folder{path})
            }
        }
    }
    return model{
        choices: folders,
        filteredChoices: folders,
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
            m.choice = m.filteredChoices[m.cursor]
            return m, tea.Quit
        case "backspace":
            m.searchBox = m.searchBox[:len(m.searchBox)-1]
        default: 
            m.searchBox += msg.String()
            m.filter()
        }
    }

    return m, nil
}

func (m model) filter() {
    filterStr := m.searchBox

    newArray := make([]folder, 0)
    for _, s := range m.choices {
        if strings.Contains(s.path, filterStr) {
            newArray = append(newArray, s)
        }
    }

    m.filteredChoices = newArray
}

func (m model) View() string {
    if m.choice.path != "" {
        return fmt.Sprintf("%s\n", m.choice.path)
    }
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
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    m, err := p.Run() 
    if err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }

    if m, ok := m.(model); ok && m.choice.path != "" {
        // fmt.Printf("cd to %s\n", m.choice.path)
        fmt.Printf(m.choice.path)
        // os.Chdir(m.choice.path)
    }
}
