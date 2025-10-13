package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	libgloss "github.com/charmbracelet/lipgloss"

	convert "github.com/muige/impsi/conversions"
)

var gap = "\n\n"

var conversions = []conversionPair{
	{"Length", "ft", "m", convert.FtToM, convert.MToFt},
}

type model struct {
	// input   string
	textInput textinput.Model
	spinner   spinner.Model
	err       error
}

type conversionPair struct {
	// name to show
	name string
	// string: from which unit
	from string
	// string: to which unit
	to string
	// fucntion to convert from -> to
	fwdfunc func(float64) float64
	// function to convert to -> from
	backfunc func(float64) float64
}

// Function checks that input contains only allowed characters (Runes in bubbletea)
// The allowed characters are: integers form 0 to 9, '-' if as the first char and one decimal dot.
func isAllowedRunes(runes []rune, current string) bool {
	dotSeen := false
	for _, c := range current {
		if c == '.' {
			dotSeen = true
			break
		}
	}

	for i, r := range runes {
		switch {
		case unicode.IsDigit(r):
			continue
		case r == '.' && !dotSeen:
			dotSeen = true
		case r == '-' && len(current) == 0 && i == 0:
			continue
		default:
			return false
		}
	}
	return true
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = libgloss.NewStyle().Foreground(libgloss.Color("205"))

	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 20

	return model{
		spinner:   s,
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, textinput.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Handle quit keys
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyRunes:
			if !isAllowedRunes(msg.Runes, m.textInput.Value()) {
				return m, nil
			}
		}
	}
	var cmd tea.Cmd
	// Update spinner
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	// Update text input
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// Header
	s := "Converting SI <-> Imperial\n\n"
	s += fmt.Sprintf("%s Insert a number: %s\n%s",
		m.spinner.View(),
		m.textInput.View(),
		gap)
	val, err := strconv.ParseFloat(m.textInput.Value(), 64)
	if err != nil {
		s += "Cannot convert input value to a floating point number"
	} else {
		for _, conv := range conversions {
			s += fmt.Sprintf("%s\n", conv.name)
			s += fmt.Sprintf("%.2f %s = %.2f %s\n", val, conv.from, conv.fwdfunc(val), conv.to)
			s += fmt.Sprintf("%.2f %s = %.2f %s\n", val, conv.to, conv.backfunc(val), conv.from)
		}
	}
	s += "\nPress ctrl-c or Esc to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There was an error: %v", err)
		os.Exit(1)
	}
}
