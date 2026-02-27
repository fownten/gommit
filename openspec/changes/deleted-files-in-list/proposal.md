## Why
The current logic for detecting staged files might be missing certain git statuses, specifically deleted files, or may not be clearly presenting them in the final list. Users want to ensure all types of staged changes (Added, Modified, Deleted) are explicitly tracked and displayed.

## What Changes
- Audit and refine the `git status --porcelain` parsing logic to ensure `D` (Deleted) and potentially `R` (Renamed) statuses are correctly captured.
- Ensure the UI correctly displays these deleted files in the "Staged files" list.

## Capabilities

### New Capabilities

### Modified Capabilities
- `detect-staged-files`: Update parsing logic to robustly include deletions and renames.

## Impact
- `internal/git/status.go`: Update `ParsePorcelainStatus`.
- `internal/git/status_test.go`: Add test cases for deleted and renamed files.
