package main

import (
	"fmt"
)

func (m model) View() string {
	if m.chatMode {
		s := fmt.Sprintf("Chat Mode (Model: %s)\n\n", m.currentModel)

		for _, msg := range m.chatHistory {
			s += fmt.Sprintf("%s\n", msg)
		}

		s += fmt.Sprintf("\n> %s", m.chatInput)
		s += "\n\nPress ESC to go back.\n"

		return s
	}

	s := "Tollama\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
