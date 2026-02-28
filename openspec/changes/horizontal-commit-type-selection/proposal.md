## Why

The current vertical list selection for commit types (feat, fix, chore) takes up significant vertical space in the terminal, even with only a few options. A horizontal layout is more compact and visually aligned with modern CLI tools, providing a cleaner "pill-style" or "tab-style" selection experience.

## What Changes

- **Horizontal Layout**: The commit type selection will be rendered horizontally rather than in a vertical list.
- **Navigation Update**: Navigation between commit types will switch to horizontal keys (Left/Right/H/L/Tab) to match the new layout.
- **Enhanced Styling**: Selected types will use more distinct styling (e.g., inverted colors or background highlights) to make the selection clear in a horizontal row.
- **Compact UI**: The overall vertical footprint of the initial state will be reduced.

## Capabilities

### New Capabilities
- None

### Modified Capabilities
- `stylized-interactive-ui`: Change `Interactive Type Selection` to use horizontal layout and navigation instead of vertical list selection.

## Impact

- `internal/ui/model.go`: The `Update` method will need to handle horizontal navigation keys for `StateSelectingType`.
- `internal/ui/model.go`: The `View` method will need to use `lipgloss.JoinHorizontal` or similar to render the type options in a single row.
- `internal/ui/styles.go`: May need new styles for horizontal selection (e.g., padding/margins between items).
