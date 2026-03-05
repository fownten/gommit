## MODIFIED Requirements

### Requirement: Description Text Input
The system SHALL provide a stylized text input field for entering the concise commit summary (subject).

#### Scenario: Entering summary
- **WHEN** the user types characters into the summary field
- **THEN** the input field SHALL display the text in real-time

## ADDED Requirements

### Requirement: Body Text Area
The system SHALL provide a stylized multiline text area for entering optional commit details (body).

#### Scenario: Entering body
- **WHEN** the user types into the body field
- **THEN** the text area SHALL display the multiline content correctly

### Requirement: Switching Input Fields
The system SHALL allow users to switch focus between the summary and body fields using a keyboard shortcut (e.g., Tab).

#### Scenario: Switching focus
- **WHEN** the user presses "tab" while focus is on the summary field
- **THEN** focus SHALL move to the body field
- **AND WHEN** the user presses "tab" again
- **THEN** focus SHALL return to the summary field
