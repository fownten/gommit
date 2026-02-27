## ADDED Requirements

### Requirement: Consistent Color Palette
The system SHALL use a consistent color palette (e.g., specific Hex codes) for borders, prompts, and highlight states using Lipgloss.

#### Scenario: Focused input
- **WHEN** an input field has focus
- **THEN** its border color SHALL change to indicate active state

### Requirement: Bordered Containers
The system SHALL wrap UI components in Lipgloss borders to provide visual separation.

#### Scenario: Main UI view
- **WHEN** the main view is rendered
- **THEN** it SHALL be enclosed in a rounded border with padding
