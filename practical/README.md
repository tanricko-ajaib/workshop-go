# Debugging Exercise: TitleCase

## Goal
Use the VS Code debugger (Delve) to find and fix a bug in `TitleCase`.

This function is intended to **capitalize the first character of each word**.

## What you should see
Run the program:

```bash
go run ./...
```

Two test cases should fail:
- `"ångström unit"` should become `"Ångström Unit"`
- `"élise dupont"` should become `"Élise Dupont"`

## Rules

- Do **not** "guess-fix" by rewriting everything.
- Use the debugger to prove the root cause.
- Make the **smallest correct change**.
- Keep all ASCII cases working.

## Success criteria

All test cases print `got == want`.