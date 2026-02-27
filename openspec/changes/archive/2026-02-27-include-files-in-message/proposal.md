## Why
Users want to see which files are being committed within the commit message itself for better traceability.

## What Changes
- Update the commit message generation logic to append a list of staged files and their statuses.

## Capabilities
- `generate-commit-message`: (Modified) Now includes file list summary.

## Impact
- `main.go`: Logic update to format the final message.
