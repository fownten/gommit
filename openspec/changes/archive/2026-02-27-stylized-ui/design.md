## Context

We are moving from a standard `fmt.Scan` based UI to a stylized, interactive UI using the Charmbracelet ecosystem (Bubble Tea and Lipgloss). The goal is to provide a modern experience that feels native to the CLI and doesn't disrupt the terminal history.

## Goals / Non-Goals

**Goals:**
- Implement a state-driven UI using Bubble Tea.
- Apply consistent styling and layouts using Lipgloss.
- Maintain a non-alt-screen (inline) execution mode.
- Centralize styling constants for easy modification.

**Non-Goals:**
- Implementing a full text editor (just simple single-line inputs).
- Supporting complex layout resizing (keep it simple for now).

## Decisions

### Decision 1: Bubble Tea Model Structure
We will use a single `Model` that manages the state of the commit creation process (Selection → Input → Confirmation).

### Decision 2: Lipgloss Styling Package
We'll create a dedicated `internal/ui/styles.go` to hold common Lipgloss styles (colors, borders, paddings). This ensures consistency across the app.

### Decision 3: Inline Rendering
We will use the `tea.WithAltScreen()` option set to `false` (default) to ensure the UI renders within the existing terminal scrollback.

## Risks / Trade-offs

- **[Risk] Terminal Compatibility** → Older terminals might not handle Lipgloss colors/borders correctly. **Mitigation**: Use standard ASCII characters where possible or detect color support.
- **[Risk] State Complexity** → Managing multiple inputs in a single Bubble Tea model can get messy. **Mitigation**: Use a state machine or sub-models if complexity increases.
