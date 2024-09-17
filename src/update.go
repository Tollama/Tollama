package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.chatMode {
			switch msg.Type {
			case tea.KeyEnter:
				m.chatHistory = append(m.chatHistory, "You: "+m.chatInput)
				reply, err := m.sendMessageToAPI(m.chatInput)
				if err != nil {
					m.chatHistory = append(m.chatHistory, "Error: "+err.Error())
				} else {
					m.chatHistory = append(m.chatHistory, "Model: "+reply)
				}
				m.chatInput = ""
			case tea.KeyEsc:
				m.chatMode = false
			case tea.KeyBackspace:
				if len(m.chatInput) > 0 {
					m.chatInput = m.chatInput[:len(m.chatInput)-1]
				}
			default:
				m.chatInput += msg.String()
			}
		} else {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			case "enter":
				m.selected = m.choices[m.cursor]

				switch m.selected {
				case "Chat":
					m.chatMode = true
				case "List Models":
					return m, m.listModels
				case "Quit":
					return m, tea.Quit
				case "Back":
					m.choices = []string{"Chat", "List Models", "Quit"}
					m.cursor = 0
				default:
					// Handle model selection and download
					for _, model := range m.models {
						if model.Name == m.selected {
							m.downloadModel(model.Name)
							m.chatMode = true
							break
						}
					}
				}
			}
		}

	case modelsMsg:
		m.models = msg.models
		m.choices = make([]string, len(msg.models)+1)
		for i, model := range msg.models {
			m.choices[i] = model.Name
		}
		m.choices[len(msg.models)] = "Back"
		m.cursor = 0

	case errMsg:
		m.choices = []string{"Error: " + msg.err.Error(), "Back"}
		m.cursor = 0
	}

	return m, nil
}
