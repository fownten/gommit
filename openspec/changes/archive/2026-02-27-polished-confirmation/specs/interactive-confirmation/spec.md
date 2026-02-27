## ADDED Requirements

### Requirement: Interactive Decision Selection
The system SHALL provide an interactive interface for confirming the commit, allowing the user to choose between "Commit" and "Cancel".

#### Scenario: Toggling between options
- **WHEN** the user presses left/right arrow keys or Tab
- **THEN** the highlighted choice SHALL switch between "Commit" and "Cancel"

### Requirement: Visual Highlighting
The system SHALL visually distinguish the currently selected choice using Lipgloss styles.

#### Scenario: Selection state
- **WHEN** "Commit" is selected
- **THEN** it SHALL be rendered with a distinct background or border color to indicate focus
