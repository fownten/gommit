## ADDED Requirements

### Requirement: Execute Commit
The system SHALL execute the `git commit` command with the provided message when the user confirms.

#### Scenario: Successful commit
- **WHEN** the user confirms the commit
- **THEN** the system SHALL invoke `git commit -m "<message>"`
- **AND** the system SHALL display a success message to the user

### Requirement: Error Handling
The system SHALL capture and display any errors returned by the git binary (e.g., pre-commit hook failure).

#### Scenario: Commit fails
- **WHEN** the git command returns a non-zero exit code
- **THEN** the system SHALL display the error message to the user
