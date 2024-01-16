package main

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"log"
	"o/dbops"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	textInput  textinput.Model
	err        error
	tableInput table.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Ronald"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	columns := []table.Column{{Title: "id", Width: 10}, {Title: "name", Width: 20}}
	rows := []table.Row{}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	return model{
		textInput:  ti,
		tableInput: t,
		err:        nil,
	}
}

func (m model) Init() tea.Cmd {
	dbops.InitDB()
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	// this actually needs to be a fuzzy search or at least first few chars

	if m.textInput.Value() != "" {
		user, err := dbops.GetUsersByNameStartsWith(m.textInput.Value())
		if err != nil {
			// Handle the error appropriately, set an error message in the model, and return.
			return model{
				textInput: m.textInput,
				err:       fmt.Errorf("error fetching user: %w", err),
			}, cmd
		}

		var rows []table.Row
		for _, u := range user {
			rows = append(rows, table.Row{strconv.Itoa(u.ID), u.Name})
		}
		m.tableInput.SetRows(rows)

	}

	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Which contact do you want to edit?\n\n%s\n\n%s\n\n%s",
		m.textInput.View(),
		m.tableInput.View(),
		"(esc to quit)",
	) + "\n"
}
