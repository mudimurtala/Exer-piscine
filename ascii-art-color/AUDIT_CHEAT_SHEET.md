# ASCII ART - 1 Page Audit Cheat Sheet

Use this as a rapid review before your audit.
Goal: explain the project from command input to terminal output, and rewrite key parts from memory.

## 1) Fast Command Patterns (Must Know)

1. Plain output:

```bash
go run . "Hello"
```

2. Color full text:

```bash
go run . --color=red "Hello"
```

3. Color only substring:

```bash
go run . --color=blue lo "Hello"
```

4. Newline inside input:

```bash
go run . "Hello\nWorld"
```

## 2) End-to-End Flow (A to Z)

1. Program starts in `main()`.
2. `parseArgs()` reads and validates CLI args.
3. If args invalid -> print usage -> return.
4. If text empty -> return.
5. Read `standard.txt` font file.
6. Split font data by `"\n"` into `[]string` lines.
7. Convert literal `\\n` in input into real newline `\n`.
8. Convert literal `\\n` in substring too (if provided).
9. `getColorCode()` resolves ANSI color code.
10. `coloredPositions()` precomputes which text indexes must be colored.
11. `asciiArtColor()` renders 8-row ASCII art for each line/character.
12. `getChar()` maps character -> its 8 font rows.
13. Color is applied where index is marked.
14. Final string printed to terminal.

## 3) Mental Model of Font Mapping

- Printable ASCII range is from 32 to 126.
- Each glyph uses 8 rows.
- In the font file, each glyph block takes 9 lines total:
  - 8 art rows
  - 1 separator line
- Mapping formula used by `getChar()`:

```go
startLine := (int(c)-32)*9 + 1
```

Why `+1`?
- Because of leading structure in the font text layout used in this implementation.

## 4) Core Function Cards (What To Say In Audit)

### main()
- Orchestrates whole program: parse args -> read font -> normalize input -> prepare color -> render -> print.

### parseArgs()
- Accepts:
  - `go run . "text"`
  - `go run . --color=<color> "text"`
  - `go run . --color=<color> <substring> "text"`
- Returns usage error for unsupported combinations.

### getColorCode(name)
- Maps color names to ANSI escapes.
- Unknown color returns empty string (no color applied).

### coloredPositions(text, substring)
- Finds every occurrence of substring in text.
- Marks all character indexes belonging to matches in `map[int]bool`.

### asciiArtColor(input, line, colored, colorCode)
- Splits input by real newline.
- Triple loop:
  - each text part (line)
  - rows `0..7`
  - each rune in part
- Uses `getChar()` to fetch glyph row.
- Applies color if current index is marked.
- Maintains `offset` and `pos` so color indexes match original full text.

### asciiArt(input, line)
- Non-color renderer (same rendering idea, no ANSI wrapping).
- Still useful for understanding baseline logic.

### getChar(c, line)
- Converts rune to starting line in font.
- Returns 8 rows slice for that character.
- Out-of-range guard returns 8 empty strings.

## 5) Rewrite Templates (Practice From Memory)

### A) Character lookup template

```go
func getChar(c rune, lines []string) []string {
    start := (int(c)-32)*9 + 1
    if start < 0 || start+8 > len(lines) {
        return make([]string, 8)
    }
    return lines[start : start+8]
}
```

### B) Basic ASCII renderer template (no color)

```go
func asciiArt(input string, lines []string) string {
    parts := strings.Split(input, "\n")
    var out string

    for i, part := range parts {
        if part == "" {
            if i < len(parts)-1 {
                out += "\n"
            }
            continue
        }

        for row := 0; row < 8; row++ {
            for _, ch := range part {
                glyph := getChar(ch, lines)
                out += glyph[row]
            }
            out += "\n"
        }
    }
    return out
}
```

### C) Substring mark template

```go
func coloredPositions(text, sub string) map[int]bool {
    m := make(map[int]bool)
    if sub == "" {
        return m
    }
    for i := 0; i <= len(text)-len(sub); i++ {
        if text[i:i+len(sub)] == sub {
            for j := 0; j < len(sub); j++ {
                m[i+j] = true
            }
        }
    }
    return m
}
```

## 6) 10 Likely Audit Questions + Short Answers

1. Why split input by newline?
- To render multi-line text independently, each with 8-row ASCII blocks.

2. Why 8 rows?
- Each glyph in the banner font is 8 lines tall.

3. Why multiply by 9 in index formula?
- Font layout stores 8 glyph lines + 1 separator per character block.

4. Why use `offset` in color rendering?
- To keep character index mapping consistent across split lines.

5. Why reset `pos := offset` on each row?
- Same source character must map to same text index in all 8 rows.

6. What happens on invalid args?
- Usage message is printed; program returns.

7. What happens if font file read fails?
- Error is printed; program returns.

8. What if color name is unknown?
- Empty color code, so output remains uncolored.

9. What if char is out of supported ASCII range?
- `getChar()` guard avoids panic and returns blank glyph rows.

10. Which font is currently used?
- `standard.txt` (hardcoded in current `main.go`).

## 7) Drill Plan (15-20 min Daily)

1. Rewrite `getChar()` from memory.
2. Rewrite non-color `asciiArt()` from memory.
3. Explain index formula out loud with one example (`'A'`, `' '` etc).
4. Add and remove newline test inputs mentally (`"Hi"`, `"Hi\n"`, `"\n"`).
5. Rewrite `parseArgs()` accepted cases.
6. Rewrite `coloredPositions()` from memory.
7. Explain `offset` vs `pos` without looking at code.

## 8) Extra Notes For Your Current Project

- Current runtime path always loads `standard.txt`.
- `shadow.txt` and `thinkertoy.txt` are present, but not selected by argument in this code version.
- Tests currently focus on baseline `asciiArt()` behavior and newline behavior.

---

If you can explain every section above out loud and rewrite sections 5A/5B/5C, you are in very strong shape for a coding audit.
