## 1. Model & UI Structure

- [x] 1.1 Add `Body` (`textarea.Model`) and `FocusedInput` (int) to `Model` in `internal/ui/model.go`.
- [x] 1.2 Initialize `Body` in `NewModel()` with placeholder "Optional details (Body)...".

## 2. Navigation & Input Handling

- [x] 2.1 Update `Update()` for `StateEnteringDescription` to handle `Tab` key for switching focus between summary and body fields.
- [x] 2.2 Ensure only the focused input receives keyboard events.
- [x] 2.3 Update `Update()` to transition to `StateConfirming` when `Enter` is pressed in the summary field or `Ctrl+Enter` in the body field.

## 3. View Implementation

- [x] 3.1 Update `View()` for `StateEnteringDescription` to render both summary and body fields with labels.
- [x] 3.2 Apply `FocusedInputStyle` and `InputStyle` dynamically based on `FocusedInput`.

## 4. Message Generation

- [x] 4.1 Update `generateMessage()` in `internal/ui/model.go` to include the optional body (with a blank line separator) if provided.

## 5. Verification

- [x] 5.1 Manually verify switching focus with Tab works correctly.
- [x] 5.2 Verify formatted output for both subject-only and subject+body commit messages.
