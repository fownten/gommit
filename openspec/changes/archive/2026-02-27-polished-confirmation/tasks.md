## 1. UI Styling

- [x] 1.1 Update `internal/ui/styles.go` with `ButtonStyle` and `ActiveButtonStyle`.

## 2. Model Logic

- [x] 2.1 Update `internal/ui/model.go` to include a `ConfirmFocused` boolean or index.
- [x] 2.2 Update the `Update` method to handle arrow keys and Tab for toggling confirmation choice.
- [x] 2.3 Update the `View` method to render the stylized buttons instead of the text prompt.

## 3. Verification

- [x] 3.1 Run `go run main.go` and verify keyboard navigation in the confirmation state.
- [x] 3.2 Verify that "Commit" still executes the commit and "Cancel" still aborts.
