package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	PrimaryColor   = lipgloss.Color("#7D56F4")
	SecondaryColor = lipgloss.Color("#04B575")
	ErrorColor     = lipgloss.Color("#FF4C4C")
	DimColor       = lipgloss.Color("#777777")

	// Styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			MarginBottom(1)

	ListStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1).
			MarginBottom(1)

	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	SelectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(SecondaryColor).
				Bold(true)

	PillStyle = lipgloss.NewStyle().
			Padding(0, 1).
			MarginRight(1).
			Foreground(DimColor)

	SelectedPillStyle = lipgloss.NewStyle().
				Padding(0, 1).
				MarginRight(1).
				Background(PrimaryColor).
				Foreground(lipgloss.Color("#FFFFFF")).
				Bold(true)

	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(DimColor).
			Padding(0, 1)

	FocusedInputStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(PrimaryColor).
				Padding(0, 1)

	ButtonStyle = lipgloss.NewStyle().
			Foreground(DimColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(DimColor).
			Padding(0, 3).
			MarginTop(1).
			MarginRight(2)

	ActiveButtonStyle = lipgloss.NewStyle().
				Foreground(PrimaryColor).
				Border(lipgloss.ThickBorder()).
				BorderForeground(PrimaryColor).
				Padding(0, 3).
				MarginTop(1).
				MarginRight(2).
				Bold(true)
)
