## Suggested debugging path

1. Set a breakpoint inside `TitleCase`, on this line:
```go
words[i] = strings.ToUpper(w[:1]) + w[1:]
```
2. Start debugging (VS Code: **Run and Debug** → `F5`).
3. When it stops:
    - Inspect `w`
    - Inspect `w[:1]`
    - Inspect `len(w)`
4. Step through the failing cases and compare what you expect vs what actually happens.

## Hint
In Go, strings are bytes. Slicing like `w[:1]` takes the first **byte**, not the first **character** (rune).