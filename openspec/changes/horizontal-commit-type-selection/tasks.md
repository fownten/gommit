## 1. UI Styling

- [ ] 1.1 Add `PillStyle` and `SelectedPillStyle` to `internal/ui/styles.go` to support horizontal layout.

## 2. Model & Navigation

- [ ] 2.1 Update the `Update` method in `internal/ui/model.go` to include horizontal navigation keys (Left, Right, H, L, Tab).
- [ ] 2.2 Ensure the navigation logic correctly increments or decrements `SelectedIdx` based on horizontal input.

## 3. View Implementation

- [ ] 3.1 Refactor the `StateSelectingType` rendering in `internal/ui/model.go`'s `View` method.
- [ ] 3.2 Implement `lipgloss.JoinHorizontal` to arrange commit types in a single row with proper styling and spacing.

## 4. Verification

- [ ] 4.1 Manually verify that Left/Right/H/L/Tab navigation works as expected.
- [ ] 4.2 Verify that the selected commit type is clearly highlighted in the horizontal row.
- [ ] 4.3 Confirm the overall UI is more compact vertically.
