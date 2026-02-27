## Why
The current CLI uses basic standard input prompts, which lack visual feedback and a modern interactive feel. We want to provide a professional, engaging user experience that stays within the scrollback buffer (no full-screen takeover).

## What Changes
- Integrate **Bubble Tea** for the interactive input loop.
- Use **Lipgloss** for consistent styling (borders, colors, and typography).
- Ensure the program runs in **inline mode** (Alt-Screen disabled) for a seamless terminal experience.

## Capabilities

### New Capabilities
- `stylized-interactive-ui`: A Bubble Tea model that handles the selection of commit types and text input for descriptions.
- `lipgloss-styles`: A central styling module for the application's visual identity.

### Modified Capabilities
- `generate-commit-message`: The core logic will be wrapped in the new UI view.

## Impact
- `main.go`: Transition from procedural prompts to the Bubble Tea application loop.
- `go.mod`: New dependencies on `github.com/charmbracelet/bubbletea` and `github.com/charmbracelet/lipgloss`.
- `internal/ui/`: New package for styling and Bubble Tea components.
