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
	"github.com/charmbracelet/lipgloss/table"

	convert "github.com/muige/impsi/conversions"
)

var gap = "\n\n"

type model struct {
	// input   string
	textInput textinput.Model
	spinner   spinner.Model
	err       error
}

// Function checks that input contains only allowed runes
// The allowed runes are: integers form 0 to 9, '-' if as the first char and one decimal dot.
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
	// Styles TODO: More styles with adaptive colors
	headerStyle := libgloss.NewStyle().Bold(true).Foreground(libgloss.AdaptiveColor{Light: "0", Dark: "12"})
	footerStyle := libgloss.NewStyle().Bold(true).Foreground(libgloss.AdaptiveColor{Light: "0", Dark: "12"})
	cellStyle := libgloss.NewStyle().Padding(0, 2)
	// borderStyle := libgloss.NewStyle().Border(libgloss.NormalBorder())
	// columnHeaderStyle := libgloss.NewStyle().Bold(true).Underline(true).Foreground(libgloss.AdaptiveColor{Light: "0", Dark: "12"})

	s := headerStyle.Render("Converting SI <-> Imperial\n")
	s += fmt.Sprintf("\n%s Insert a number: %s\n%s",
		m.spinner.View(),
		m.textInput.View(),
		gap)
	var val float64 = 0
	val, _ = strconv.ParseFloat(m.textInput.Value(), 64)
	// Create the tables
	impToSIRows := [][]string{
		{"Distance", "mi -> km", fmt.Sprintf("%.2f", convert.MilesToKm(val))},
		{"", "ft -> m", fmt.Sprintf("%.2f", convert.FtToM(val))},
		{"", "yd -> m", fmt.Sprintf("%.2f", convert.YdToM(val))},
		{"", "in -> cm", fmt.Sprintf("%.2f", convert.InToCm(val))},
		{"Weight", "lbs -> kg", fmt.Sprintf("%.2f", convert.LbsToKg(val))},
		{"", "oz -> g", fmt.Sprintf("%.2f", convert.OzToG(val))},
		{"Volume", "Gal -> l", fmt.Sprintf("%.2f", convert.GalToL(val))},
		{"", "fl oz -> ml", fmt.Sprintf("%.2f", convert.FlOzToMl(val))},
		{"Temperature", "F -> C", fmt.Sprintf("%.2f", convert.FToC(val))},
	}
	siToImpRows := [][]string{
		{"Distance", "km -> mi", fmt.Sprintf("%.2f", convert.KmToMiles(val))},
		{"", "m -> ft", fmt.Sprintf("%.2f", convert.MToFt(val))},
		{"", "m -> yd", fmt.Sprintf("%.2f", convert.MToYd(val))},
		{"", "cm -> in", fmt.Sprintf("%.2f", convert.CmToIn(val))},
		{"Weight", "kg -> lbs", fmt.Sprintf("%.2f", convert.KgToLbs(val))},
		{"", "g -> oz", fmt.Sprintf("%.2f", convert.GToOz(val))},
		{"Volume", "l -> Gal", fmt.Sprintf("%.2f", convert.LToGal(val))},
		{"", "ml -> fl oz", fmt.Sprintf("%.2f", convert.MlToFlOz(val))},
		{"Temperature", "C -> F", fmt.Sprintf("%.2f", convert.CToF(val))},
	}
	// TODO: Make the tables prettier and use adaptive colors
	impToSIRowsTable := table.New().
		Border(libgloss.NormalBorder()).
		StyleFunc(func(row, col int) libgloss.Style { return cellStyle }).
		Rows(impToSIRows...)
	siToImpRowsTable := table.New().
		Border(libgloss.NormalBorder()).
		StyleFunc(func(row, col int) libgloss.Style { return cellStyle }).
		Rows(siToImpRows...)
	// Use libgloss to join the tables side by side
	tables := libgloss.JoinHorizontal(libgloss.Top, impToSIRowsTable.Render(), "    ", siToImpRowsTable.Render())
	s += tables

	// Render the tables and add them to s
	s += footerStyle.Render("\nPress ctrl-c or Esc to quit.\n")
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There was an error: %v", err)
		os.Exit(1)
	}
}
