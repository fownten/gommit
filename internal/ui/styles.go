package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	PinkPastel = lipgloss.Color("#F48FB1")
	BluePastel = lipgloss.Color("#90CAF9")
	ErrorColor = lipgloss.Color("#FF4C4C")
	GrayPastel = lipgloss.Color("#B0BEC5")
	DarkGray   = lipgloss.Color("#12424a")

	// Styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(PinkPastel).
			Bold(true).
			MarginBottom(1)

	ListStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PinkPastel).
			Padding(1).
			MarginBottom(1)

	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	SelectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(BluePastel).
				Bold(true)

	PillStyle = lipgloss.NewStyle().
			Padding(0, 1).
			MarginRight(1).
			Foreground(PinkPastel)

	SelectedPillStyle = lipgloss.NewStyle().
				Padding(0, 1).
				MarginRight(1).
				Background(PinkPastel).
				Foreground(DarkGray).
				Bold(true)

	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(GrayPastel).
			Padding(0, 1)

	FocusedInputStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(PinkPastel).
				Padding(0, 1)

	ButtonStyle = lipgloss.NewStyle().
			Foreground(GrayPastel).
			Border(lipgloss.NormalBorder()).
			BorderForeground(GrayPastel).
			Padding(0, 3).
			MarginTop(1).
			MarginRight(2)

	ActiveButtonStyle = lipgloss.NewStyle().
				Foreground(PinkPastel).
				Border(lipgloss.ThickBorder()).
				BorderForeground(PinkPastel).
				Padding(0, 3).
				MarginTop(1).
				MarginRight(2).
				Bold(true)
)
