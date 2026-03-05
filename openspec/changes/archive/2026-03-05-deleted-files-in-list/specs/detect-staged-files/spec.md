## ADDED Requirements

### Requirement: Detect All Staged Change Types
The system SHALL identify all files staged for commit, specifically including files that have been deleted in the index.

#### Scenario: Staged deletion
- **WHEN** a file is staged for deletion
- **THEN** the system SHALL include the file in the staged files list with status "Deleted"
