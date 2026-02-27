package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fownten/gommit/internal/git"
	"github.com/charmbracelet/lipgloss"
)

type State int

const (
	StateSelectingType State = iota
	StateEnteringDescription
	StateConfirming
	StateDone
)

type Model struct {
	State           State
	CommitTypes     []string
	SelectedIdx     int
	Description     textinput.Model
	StagedFiles     []git.StagedFile
	Confirmed       bool
	FinalMsg        string
	ConfirmSelected int // 0 for Commit, 1 for Cancel
}

func NewModel(staged []git.StagedFile) Model {
	ti := textinput.New()
	ti.Placeholder = "Enter commit description..."
	ti.Focus()

	return Model{
		State:           StateSelectingType,
		CommitTypes:     []string{"feat", "fix", "chore"},
		SelectedIdx:     0,
		Description:     ti,
		StagedFiles:     staged,
		ConfirmSelected: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		switch m.State {
		case StateSelectingType:
			switch msg.String() {
			case "up", "k":
				if m.SelectedIdx > 0 {
					m.SelectedIdx--
				}
			case "down", "j":
				if m.SelectedIdx < len(m.CommitTypes)-1 {
					m.SelectedIdx++
				}
			case "enter":
				m.State = StateEnteringDescription
				m.Description.Focus()
			}

		case StateEnteringDescription:
			switch msg.String() {
			case "enter":
				if strings.TrimSpace(m.Description.Value()) != "" {
					m.FinalMsg = m.generateMessage()
					m.State = StateConfirming
				}
			case "esc":
				m.State = StateSelectingType
			default:
				m.Description, cmd = m.Description.Update(msg)
			}

		case StateConfirming:
			switch msg.String() {
			case "left", "right", "h", "l", "tab":
				if m.ConfirmSelected == 0 {
					m.ConfirmSelected = 1
				} else {
					m.ConfirmSelected = 0
				}
			case "enter":
				if m.ConfirmSelected == 0 {
					m.Confirmed = true
				} else {
					m.Confirmed = false
				}
				m.State = StateDone
				return m, tea.Quit
			case "esc":
				m.State = StateEnteringDescription
			}
		}
	}

	return m, cmd
}

func (m Model) View() string {
	var s strings.Builder

	s.WriteString(TitleStyle.Render("Gommit - Interactive Commit"))
	s.WriteString("\n")

	switch m.State {
	case StateSelectingType:
		s.WriteString("Select commit type:\n")
		var list strings.Builder
		for i, t := range m.CommitTypes {
			if i == m.SelectedIdx {
				list.WriteString(SelectedItemStyle.Render(fmt.Sprintf("> %s", t)))
			} else {
				list.WriteString(ItemStyle.Render(fmt.Sprintf("  %s", t)))
			}
			list.WriteString("\n")
		}
		s.WriteString(ListStyle.Render(list.String()))

	case StateEnteringDescription:
		s.WriteString(fmt.Sprintf("Type: %s\n\n", m.CommitTypes[m.SelectedIdx]))
		s.WriteString("Enter description:\n")
		s.WriteString(FocusedInputStyle.Render(m.Description.View()))

	case StateConfirming, StateDone:
		s.WriteString("Preview:\n")
		s.WriteString(ListStyle.Render(m.FinalMsg))
		if m.State == StateConfirming {
			var commitBtn, cancelBtn string
			if m.ConfirmSelected == 0 {
				commitBtn = ActiveButtonStyle.Render("Commit")
				cancelBtn = ButtonStyle.Render("Cancel")
			} else {
				commitBtn = ButtonStyle.Render("Commit")
				cancelBtn = ActiveButtonStyle.Render("Cancel")
			}
			s.WriteString("\n" + lipgloss.JoinHorizontal(lipgloss.Top, commitBtn, cancelBtn))
		}
	}

	return s.String()
}

func (m Model) generateMessage() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s: %s\n\n", m.CommitTypes[m.SelectedIdx], m.Description.Value()))
	sb.WriteString("Staged files:\n")
	for _, f := range m.StagedFiles {
		sb.WriteString(fmt.Sprintf("- [%s] %s\n", f.Status, f.Path))
	}
	return sb.String()
}
