## 1. Setup & Styling

- [x] 1.1 Add Bubble Tea and Lipgloss dependencies: `go get github.com/charmbracelet/bubbletea github.com/charmbracelet/lipgloss`
- [x] 1.2 Create `internal/ui/styles.go` with core Lipgloss styles (colors, borders).

## 2. Interactive UI Implementation

- [x] 2.1 Implement the Bubble Tea model and `Init`, `Update`, `View` methods in `internal/ui/model.go`.
- [x] 2.2 Add support for commit type selection (list).
- [x] 2.3 Add support for description text input (using `bubbles/textinput`).

## 3. Integration

- [x] 3.1 Refactor `main.go` to initialize and run the Bubble Tea program.
- [x] 3.2 Ensure the final commit message is rendered and confirmed via the new UI.

## 4. Verify

- [x] 4.1 Run `go run main.go` and verify interactive behavior.
- [x] 4.2 Confirm styling matches the Lipgloss spec.
