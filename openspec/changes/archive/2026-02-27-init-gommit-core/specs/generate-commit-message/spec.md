## ADDED Requirements

### Requirement: Format Message

The system SHALL produce a commit message string in the format: `type: description`.

#### Scenario: Valid inputs provided

- **WHEN** I provide type "feat" and description "add user login"
- **THEN** the resulting message is "feat: add user login"
