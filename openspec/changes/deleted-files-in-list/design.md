## Context
We use `git status --porcelain` to identify staged files. The first column of the output indicates the status in the index.

## Goals / Non-Goals
**Goals:**
- Ensure deleted files (status `D`) are captured and displayed.
- Support renamed files (status `R`) if encountered.

## Decisions

### Decision 1: Audit Parsing Logic
We will verify and potentially expand the switch statement in `ParsePorcelainStatus` to robustly handle `D` and `R`.

## Risks / Trade-offs
- **[Risk] Path parsing** → Renamed files in porcelain mode show `R old_path -> new_path`. We need to handle the arrow correctly if we want to show the full rename. **Mitigation**: For now, we will focus on basic statuses; if `R` is encountered, we'll extract the new path.
