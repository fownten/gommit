## MODIFIED Requirements

### Requirement: Format Message with Files
The system SHALL produce a commit message string that includes the type, summary (subject), optional body (details), and a list of staged files with their status.

#### Scenario: Message with optional body
- **WHEN** I have "main.go" (Added) staged
- **AND** I provide type "feat", summary "add core", and body "This change introduces the core logic for the application.\nIt includes the main entry point and basic structures."
- **THEN** the resulting message includes:
  ```
  feat: add core

  This change introduces the core logic for the application.
  It includes the main entry point and basic structures.

  Staged files:
  - [Added] main.go
  ```
