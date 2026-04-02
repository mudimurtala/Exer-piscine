package main

import (
    "bytes"
    "io"
    "os"
    "strings"
    "testing"
)

func loadFont(t *testing.T) []string {
    t.Helper()
    data, err := os.ReadFile("standard.txt")
    if err != nil {
        t.Fatalf("could not load standard.txt: %v", err)
    }
    return strings.Split(string(data), "\n")
}

func withArgs(t *testing.T, args []string, fn func()) {
    t.Helper()
    original := os.Args
    os.Args = args
    defer func() { os.Args = original }()
    fn()
}

func captureStdout(t *testing.T, fn func()) string {
    t.Helper()
    original := os.Stdout
    r, w, err := os.Pipe()
    if err != nil {
        t.Fatalf("pipe creation failed: %v", err)
    }
    os.Stdout = w

    fn()

    _ = w.Close()
    os.Stdout = original

    var buf bytes.Buffer
    if _, err := io.Copy(&buf, r); err != nil {
        t.Fatalf("could not read captured stdout: %v", err)
    }
    _ = r.Close()
    return buf.String()
}

func TestParseArgsNoArguments(t *testing.T) {
    withArgs(t, []string{"cmd"}, func() {
        _, _, _, err := parseArgs()
        if err == nil {
            t.Fatalf("expected error for no args")
        }
    })
}

func TestParseArgsPlainText(t *testing.T) {
    withArgs(t, []string{"cmd", "Hello"}, func() {
        color, sub, text, err := parseArgs()
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if color != "" || sub != "" || text != "Hello" {
            t.Fatalf("unexpected parse result: color=%q sub=%q text=%q", color, sub, text)
        }
    })
}

func TestParseArgsColorWholeString(t *testing.T) {
    withArgs(t, []string{"cmd", "--color=red", "Hello"}, func() {
        color, sub, text, err := parseArgs()
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if color != "red" || sub != "Hello" || text != "Hello" {
            t.Fatalf("unexpected parse result: color=%q sub=%q text=%q", color, sub, text)
        }
    })
}

func TestParseArgsColorSubstring(t *testing.T) {
    withArgs(t, []string{"cmd", "--color=blue", "lo", "Hello"}, func() {
        color, sub, text, err := parseArgs()
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if color != "blue" || sub != "lo" || text != "Hello" {
            t.Fatalf("unexpected parse result: color=%q sub=%q text=%q", color, sub, text)
        }
    })
}

func TestParseArgsInvalidCombinations(t *testing.T) {
    cases := [][]string{
        {"cmd", "a", "b"},
        {"cmd", "--color=green", "a", "b", "c"},
    }

    for _, args := range cases {
        args := args
        t.Run(strings.Join(args, "_"), func(t *testing.T) {
            withArgs(t, args, func() {
                _, _, _, err := parseArgs()
                if err == nil {
                    t.Fatalf("expected usage error for args: %v", args)
                }
            })
        })
    }
}

func TestGetColorCodeKnownAndUnknown(t *testing.T) {
    if got := getColorCode("red"); got != "\033[31m" {
        t.Fatalf("expected red ANSI code, got %q", got)
    }
    if got := getColorCode("unknown"); got != "" {
        t.Fatalf("expected empty code for unknown color, got %q", got)
    }
}

func TestColoredPositionsOverlap(t *testing.T) {
    positions := coloredPositions("banana", "ana")
    expected := []int{1, 2, 3, 4, 5}
    for _, idx := range expected {
        if !positions[idx] {
            t.Fatalf("expected position %d to be colored", idx)
        }
    }
    if positions[0] {
        t.Fatalf("did not expect position 0 to be colored")
    }
}

func TestGetCharOutOfRangeReturnsBlankGlyph(t *testing.T) {
    lines := loadFont(t)
    glyph := getChar(rune(10), lines)
    if len(glyph) != 8 {
        t.Fatalf("expected 8 rows, got %d", len(glyph))
    }
    for i, row := range glyph {
        if row != "" {
            t.Fatalf("expected empty row at %d, got %q", i, row)
        }
    }
}

func TestAsciiArtReturnsOutput(t *testing.T) {
    lines := loadFont(t)
    result := asciiArt("Hello", lines)
    if result == "" {
        t.Fatalf("expected output for Hello, got empty string")
    }
}

func TestAsciiArtNewlineProducesBlankLine(t *testing.T) {
    lines := loadFont(t)
    result := asciiArt("\n", lines)
    if result != "\n" {
        t.Fatalf("expected single newline, got %q", result)
    }
}

func TestAsciiArtTrailingNewlineNoExtraLines(t *testing.T) {
    lines := loadFont(t)
    withNewline := asciiArt("Hello\n", lines)
    withoutNewline := asciiArt("Hello", lines)
    if withNewline != withoutNewline {
        t.Fatalf("trailing newline changed output unexpectedly")
    }
}

func TestAsciiArtColorWithoutColorCodeMatchesAsciiArt(t *testing.T) {
    lines := loadFont(t)
    input := "Hi"
    colored := coloredPositions(input, input)
    got := asciiArtColor(input, lines, colored, "")
    want := asciiArt(input, lines)
    if got != want {
        t.Fatalf("expected uncolored output to match asciiArt")
    }
}

func TestAsciiArtColorAddsAnsiSequences(t *testing.T) {
    lines := loadFont(t)
    input := "A"
    colored := coloredPositions(input, input)
    got := asciiArtColor(input, lines, colored, "\033[31m")
    if !strings.Contains(got, "\033[31m") {
        t.Fatalf("expected output to include red ANSI code")
    }
    if !strings.Contains(got, "\033[0m") {
        t.Fatalf("expected output to include ANSI reset code")
    }
}

func TestMainInvalidArgsPrintsUsage(t *testing.T) {
    withArgs(t, []string{"cmd"}, func() {
        out := captureStdout(t, main)
        if !strings.Contains(out, "Usage: go run .") {
            t.Fatalf("expected usage output, got %q", out)
        }
    })
}

func TestMainPlainInputPrintsAsciiArt(t *testing.T) {
    withArgs(t, []string{"cmd", "Hi"}, func() {
        out := captureStdout(t, main)
        if out == "" {
            t.Fatalf("expected non-empty output")
        }
        if strings.Contains(out, "Usage: go run .") {
            t.Fatalf("did not expect usage output")
        }
    })
}

func TestMainColorInputPrintsAnsiOutput(t *testing.T) {
    withArgs(t, []string{"cmd", "--color=red", "H", "Hi"}, func() {
        out := captureStdout(t, main)
        if !strings.Contains(out, "\033[31m") {
            t.Fatalf("expected ANSI colored output")
        }
    })
}
