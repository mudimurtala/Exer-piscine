# Mock Auditor Session - ASCII Art

How to use this file:
1. Read one prompt.
2. Answer out loud without checking code.
3. Compare with the model answer.
4. Score yourself from 0 to 2:
- 0 = incorrect
- 1 = partly correct
- 2 = clear and complete

Target: 24/30+ before audit.

## Prompt 1
Question:
Where does execution start, and what are the first 3 major operations?

Model answer:
Execution starts in main(). First: parse command arguments with parseArgs(). Second: read standard.txt font file and split it into lines. Third: normalize text by replacing literal \\n with real newlines before rendering.

## Prompt 2
Question:
What argument formats are valid?

Model answer:
Valid forms are:
- go run . "text"
- go run . --color=<color> "text"
- go run . --color=<color> <substring> "text"
Any other combination returns usage error.

## Prompt 3
Question:
Why does getChar() use (int(c)-32)*9 + 1 ?

Model answer:
Printable ASCII starts at 32. Each glyph in banner data occupies 9 lines in layout terms (8 glyph rows + separator), so we jump by 9 per character and then slice 8 rows for the glyph.

## Prompt 4
Question:
Explain the rendering loops in simple words.

Model answer:
The renderer uses three loops: for each text line, for each of 8 ASCII-art rows, and for each character in that line. For every character-row pair, it fetches glyph data from the font and appends it to output.

## Prompt 5
Question:
How does the project support multiline text input?

Model answer:
It converts literal \\n from CLI input into actual newline characters, then splits input by newline and renders each part separately.

## Prompt 6
Question:
How does color get applied only to a substring?

Model answer:
coloredPositions() scans full text for all substring matches and stores matched character indexes in map[int]bool. During rendering, if current character index is marked and a valid color code exists, that glyph row is wrapped in ANSI color/reset.

## Prompt 7
Question:
Why are offset and pos both needed in asciiArtColor()?

Model answer:
offset tracks where each split part begins in the original full text. pos tracks current character index while iterating inside a row. Resetting pos to offset each row ensures the same source character keeps the same index across all 8 rows.

## Prompt 8
Question:
What happens if a color name is invalid?

Model answer:
getColorCode() returns an empty string. Rendering still works, but no ANSI color is added.

## Prompt 9
Question:
What happens if a character is outside supported range?

Model answer:
getChar() guards indexes. If out of range, it returns 8 empty strings instead of panicking.

## Prompt 10
Question:
Why does main print usage in one case and file error in another?

Model answer:
Usage output appears when parseArgs() returns an argument/format error. File error appears when standard.txt cannot be read.

## Prompt 11
Question:
What is tested in main_test.go now?

Model answer:
Tests cover parseArgs valid/invalid cases, ANSI color code mapping, overlapping substring detection, getChar out-of-range behavior, asciiArt newline behavior, asciiArtColor ANSI insertion behavior, and end-to-end main output for invalid/plain/colored inputs.

## Prompt 12
Question:
How would you manually verify colored output quickly?

Model answer:
Run:
- go run . --color=red H "Hi"
You should see ANSI-colored glyphs where H appears. If color is unknown, output should still print but without color.

## Prompt 13
Question:
What are shadow.txt and thinkertoy.txt currently used for?

Model answer:
They are present as alternate fonts, but current main implementation always reads standard.txt.

## Prompt 14
Question:
Why might tests fail on some machines if run from a different directory?

Model answer:
Tests read standard.txt with a relative path, so they should run from the ascii-art module directory where that file exists.

## Prompt 15
Question:
Give a 30-second project summary for an auditor.

Model answer:
This Go program converts input text into banner-style ASCII art using standard.txt. It parses CLI args, optionally accepts color mode with full-string or substring coloring, normalizes multiline input via \\n conversion, maps each character to 8 glyph rows with ASCII-based indexing, and prints rendered output to stdout. Tests validate parsing, rendering, color behavior, and main-level integration.

## Lightning Rewrite Drills

1. Rewrite parseArgs() accepted branches from memory.
2. Rewrite getChar() safely with out-of-range guard.
3. Rewrite coloredPositions() with overlap support.
4. Rewrite basic renderer triple-loop (line -> row -> rune).
5. Explain offset/pos without looking at code.

## Self-Evaluation Grid

- Prompt 1: __/2
- Prompt 2: __/2
- Prompt 3: __/2
- Prompt 4: __/2
- Prompt 5: __/2
- Prompt 6: __/2
- Prompt 7: __/2
- Prompt 8: __/2
- Prompt 9: __/2
- Prompt 10: __/2
- Prompt 11: __/2
- Prompt 12: __/2
- Prompt 13: __/2
- Prompt 14: __/2
- Prompt 15: __/2

Total: __/30

Readiness guide:
- 0-17: Keep drilling core flow and mapping formula.
- 18-23: Good progress; tighten wording and confidence.
- 24-30: Audit-ready level for explanation and rewrites.
