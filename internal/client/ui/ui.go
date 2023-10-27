package ui

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Client interface {
	GetUserDataList() error
	GetUserData() error
	SaveData() error
}

type MainModel struct {
	choices []string
	cursor  int
	client  Client
}

func InitialUIModel(client Client) MainModel {
	choices := []string{
		"Get all saved data",
		"Get some saved data",
		"Save some data",
	}
	return MainModel{
		choices: choices,
		client:  client,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: // Is it a key press?
		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c", "q": // These keys should exit the program.
			return m, tea.Quit
		case "up": // The "up" key move the cursor up
			if m.cursor > 0 {
				m.cursor--
			}
		case "down": // The "down" key move the cursor down
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			choice := m.choices[m.cursor]
			err := processCommand(m.client, choice)
			if err != nil {
				log.Fatal(err)
			}
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m MainModel) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %v. %s\n", cursor, i+1, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func processCommand(c Client, msg string) error {
	switch msg {
	case "Get all saved data":
		err := c.GetUserDataList()
		if err != nil {
			return err
		}
	case "Get some saved data":
		err := c.GetUserData()
		if err != nil {
			return err
		}
	case "Save some data":
		err := c.SaveData()
		if err != nil {
			return err
		}
	}

	return nil
}
