## ADDED Requirements

### Requirement: List Staged Files

The system SHALL identify files that are currently staged in the git index and categorize them as Added, Modified, or Deleted.

#### Scenario: Staged changes exist

- **WHEN** I have files staged for commit
- **THEN** the system returns a list of objects containing the file path and its status (Added, Modified, or Deleted)
