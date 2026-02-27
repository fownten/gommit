# generate-commit-message Specification

## Purpose
TBD - created by archiving change include-files-in-message. Update Purpose after archive.
## Requirements
### Requirement: Format Message with Files

The system SHALL produce a commit message string that includes the type, description, and a list of staged files with their status.

#### Scenario: Multiple files staged
- **WHEN** I have "main.go" (Added) and "README.md" (Modified) staged
- **AND** I provide type "feat" and description "add core"
- **THEN** the resulting message includes:
  ```
  feat: add core

  Staged files:
  - [Added] main.go
  - [Modified] README.md
  ```

