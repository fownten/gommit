## Context

Currently, `gommit` only allows a single-line commit message. This design outlines how to add an optional, multiline body field to provide more context.

## Goals / Non-Goals

**Goals:**
- Provide a dedicated input field for an optional commit body.
- Support switching between summary and body fields.
- Include the body in the final generated commit message.

**Non-Goals:**
- Multiple types or breaking changes (out of scope for this specific change).
- Rich text or markdown preview for the body (plain text only).

## Decisions

- **Input Component**: Use `github.com/charmbracelet/bubbles/textarea` for the body and continue using `textinput` for the summary.
- **Model State**: Add a new `Body textinput.Model` (or `textarea.Model`) to the `Model` struct.
- **Focus Control**:
    - Focus summary by default.
    - `Tab` key toggles focus between summary and body fields in `StateEnteringDescription`.
- **UI Layout**:
    - Display summary input at the top.
    - Display body input below the summary, with a clear label (e.g., "Optional Details:").
    - Use `FocusedInputStyle` to highlight which field is currently active.
- **Message Generation**:
    - `generateMessage()` will check if `Body` is non-empty.
    - If non-empty, append to subject with a blank line separator.

## Risks / Trade-offs

- **[Risk]** Large bodies might overflow the terminal. → **Mitigation**: `textarea` handles basic scrolling and wrapping.
- **[Risk]** Keyboard conflicts with `tab`. → **Mitigation**: `Tab` is standard for switching focus in many TUIs and GUIs.
