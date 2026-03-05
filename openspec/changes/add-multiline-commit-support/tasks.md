## 1. Model & UI Structure

- [ ] 1.1 Add `Body` (`textarea.Model`) and `FocusedInput` (int) to `Model` in `internal/ui/model.go`.
- [ ] 1.2 Initialize `Body` in `NewModel()` with placeholder "Optional details (Body)...".

## 2. Navigation & Input Handling

- [ ] 2.1 Update `Update()` for `StateEnteringDescription` to handle `Tab` key for switching focus between summary and body fields.
- [ ] 2.2 Ensure only the focused input receives keyboard events.
- [ ] 2.3 Update `Update()` to transition to `StateConfirming` when `Enter` is pressed in the summary field or `Ctrl+Enter` in the body field.

## 3. View Implementation

- [ ] 3.1 Update `View()` for `StateEnteringDescription` to render both summary and body fields with labels.
- [ ] 3.2 Apply `FocusedInputStyle` and `InputStyle` dynamically based on `FocusedInput`.

## 4. Message Generation

- [ ] 4.1 Update `generateMessage()` in `internal/ui/model.go` to include the optional body (with a blank line separator) if provided.

## 5. Verification

- [ ] 5.1 Manually verify switching focus with Tab works correctly.
- [ ] 5.2 Verify formatted output for both subject-only and subject+body commit messages.
