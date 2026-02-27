## Why
The current confirmation prompt is a simple text-based `(y/n)`. We want to replace it with a stylized, interactive Bubble Tea component that aligns with the rest of the UI.

## What Changes
- Replace the standard input confirmation with a visual selection component.
- Add keyboard navigation (arrows or tab) to toggle between "Commit" and "Cancel".
- Style the buttons using Lipgloss for a professional look.

## Capabilities

### New Capabilities
- `interactive-confirmation`: A Bubble Tea component for the final commit decision.

### Modified Capabilities

## Impact
- `internal/ui/model.go`: Update the `StateConfirming` view and logic.
- `internal/ui/styles.go`: Add new button and selection styles.
