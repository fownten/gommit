## Context
The current confirmation state in the Bubble Tea model uses a simple text prompt. We want to transition this to a component-based approach with clear visual buttons.

## Goals / Non-Goals
**Goals:**
- Replace the boolean confirmation input with a focused button selection.
- Use Lipgloss for button styling (Active/Inactive states).
- Improve the visual layout of the preview screen.

## Decisions

### Decision 1: Button Layout
We will render the buttons side-by-side using `lipgloss.JoinHorizontal`. This provides a more traditional CLI "dialog" feel.

### Decision 2: State Management
The `Model` will track which button is currently focused using a boolean or integer index specifically for the `StateConfirming` state.

## Risks / Trade-offs
- **[Trade-off] Vertical Space** → Side-by-side buttons might be harder to read on very narrow terminals. **Mitigation**: Use compact padding if necessary.
