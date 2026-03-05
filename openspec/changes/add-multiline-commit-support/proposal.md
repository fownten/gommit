## Why

A single-line commit description is often insufficient for providing necessary context about a change. This follows the industry-standard git practice of having a concise subject line followed by an optional, detailed body.

## What Changes

- **Optional Body Field**: A new, optional multiline text area for providing detailed commit notes.
- **Enhanced Message Generation**: If a body is provided, it will be included in the final commit message string after a blank line following the subject line.
- **Transitioning UI**: Users can switch between the summary (subject) input and the body (details) input using a keyboard shortcut (e.g., Tab).

## Capabilities

### New Capabilities
- None

### Modified Capabilities
- `generate-commit-message`: Now includes an optional body in the generated message.
- `stylized-interactive-ui`: Now includes a multiline text area for entering optional commit details.

## Impact

- `internal/ui/model.go`: The `Model` struct will need a new field for the multiline input.
- `internal/ui/model.go`: Navigation logic needs to handle switching focus between fields.
- `internal/ui/model.go`: `generateMessage()` method must incorporate the optional body.
- `internal/ui/styles.go`: New styles for the multiline text area may be needed.
