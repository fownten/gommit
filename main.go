package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fownten/gommit/internal/git"
	"github.com/fownten/gommit/internal/ui"
)

func main() {
	staged, err := git.GetStagedFiles()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(staged) == 0 {
		fmt.Println("No staged files found. Please stage some files first.")
		return
	}

	m := ui.NewModel(staged)
	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	m = finalModel.(ui.Model)

	if m.Confirmed {
		fmt.Println("\nCommitting...")
		if err := git.ExecuteCommit(m.FinalMsg); err != nil {
			fmt.Printf("\nError: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\n✓ Commit successful!")
	} else {
		fmt.Println("\nCommit cancelled.")
	}
}
