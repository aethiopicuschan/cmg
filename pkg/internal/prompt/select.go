package prompt

import (
	"fmt"
	"os"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	label   string
	options []string
	cursor  int
	choice  int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// Exit without selection (optional)
		case "ctrl+c", "q":
			os.Exit(0)

		// Move cursor
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
			return m, nil

		// Confirm selection
		case "enter":
			m.choice = m.cursor
			return m, tea.Quit

		// Ignore all other keys
		default:
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("%s:\n", m.label)
	for i, opt := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = "▶"
		}
		s += fmt.Sprintf(" %s %s\n", cursor, opt)
	}
	s += "\nMove with ↑ ↓ and confirm with Enter"
	return s
}

func PromptSelect(label string, options []string, defaultIndex int) int {
	m := model{
		label:   label,
		options: options,
		cursor:  defaultIndex,
		choice:  defaultIndex,
	}

	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return finalModel.(model).choice
}

// PromptSelectOrInputIndex lets the user choose from predefined options
// or enter a custom value manually.
//
// Return values:
//   - If a predefined option is selected: (index, "")
//   - If "Other (custom)" is selected:    (-1, customValue)
func PromptSelectOrInputIndex(
	label string,
	inputLabel string,
	options []string,
	defaultValue string,
) (int, string) {
	const customLabel = "Other (custom)"

	// If no options are provided, fallback to free input
	if len(options) == 0 {
		return -1, PromptString(label, defaultValue)
	}

	// Build selection list with custom option
	selectOptions := slices.Clone(options)
	selectOptions = append(selectOptions, customLabel)

	// Determine default index
	defaultIndex := indexOf(selectOptions, defaultValue)

	choiceIndex := PromptSelect(
		label,
		selectOptions,
		defaultIndex,
	)

	// Custom input selected
	if choiceIndex == len(selectOptions)-1 {
		for {
			custom := PromptString(
				inputLabel,
				"",
			)
			if custom != "" {
				return -1, custom
			}
		}
	}

	// Predefined option selected
	return choiceIndex, ""
}

func indexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return 0
}
