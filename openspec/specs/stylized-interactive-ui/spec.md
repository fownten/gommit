## ADDED Requirements

### Requirement: Interactive Type Selection
The system SHALL provide an interactive list (using Bubble Tea) for the user to select the commit type (feat, fix, chore).

#### Scenario: Navigating the list
- **WHEN** the user presses arrow keys
- **THEN** the highlighted selection SHALL move accordingly

### Requirement: Description Text Input
The system SHALL provide a stylized text input field for entering the commit description.

#### Scenario: Entering text
- **WHEN** the user types characters
- **THEN** the input field SHALL display the text in real-time

### Requirement: Non-Takeover Mode
The system SHALL render the UI inline within the current terminal scrollback buffer, without switching to the alternate screen.

#### Scenario: Program exit
- **WHEN** the program completes or is cancelled
- **THEN** the previous terminal history SHALL remain visible above the final output
