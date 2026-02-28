## Context

Currently, the commit type selection in `gommit` uses a vertical list. While functional, it consumes more vertical space than necessary for a small number of options (feat, fix, chore). The project uses Bubble Tea for the TUI framework and Lipgloss for styling.

## Goals / Non-Goals

**Goals:**
- Transition the commit type selection to a horizontal layout.
- Implement intuitive horizontal navigation (Left/Right/H/L/Tab).
- Provide clear visual feedback for the selected type using "pill" or "tab" styles.
- Reduce the initial vertical footprint of the tool.

**Non-Goals:**
- Adding new commit types (out of scope for this UI change).
- Implementing multi-select (only one type should be selected).

## Decisions

- **UI Layout**: Use `lipgloss.JoinHorizontal` to arrange commit type options in a single row. Each option will have horizontal padding to ensure clear separation.
- **Styling**: 
    - Unselected options will be rendered as simple text with light padding.
    - The selected option will be rendered with inverted colors or a distinct background using Lipgloss's `Reverse(true)` or `Background()` methods to create a "pill" effect.
- **Navigation**:
    - Add handlers for `left`, `right`, `h`, `l`, and `tab` in the `Update` loop for `StateSelectingType`.
    - Keep `up` and `down` as aliases for consistency, or remove them to strictly enforce horizontal mental model (Decision: keep for now but prioritize horizontal).
- **Separation**: Use a small amount of margin or a separator character (like `|`) between horizontal items if needed, though padding on the "pills" is preferred.

## Risks / Trade-offs

- **[Risk]** Horizontal space constraints on very narrow terminals. → **Mitigation**: The three default types (feat, fix, chore) are short enough that they will fit in any reasonable terminal width (under 40 characters).
- **[Risk]** Users accustomed to vertical navigation might be confused. → **Mitigation**: Retain `up/down` as secondary navigation keys or ensure the visual layout clearly implies horizontal movement.
