## Project Overview

GoReloaded is a Go project for advanced text processing and transformation. It reads an input text file, applies a series of transformations (case changes, article corrections, number conversions, punctuation fixes), and writes the processed result to an output file.

## Features

- **Case Transformation Tags:** Supports tags like `(up)`, `(low)`, `(cap)` to change the case of preceding words.
- **Number Conversion Tags:** Converts hexadecimal and binary numbers to decimal using `(hex)` and `(bin)` markers.
- **Article Correction:** Automatically changes "a" to "an" before words starting with vowels or 'h'.
- **Punctuation Fixes:** Cleans up spaces before punctuation and ensures proper formatting.
- **Efficient Processing:** Handles all transformations in a single pipeline.

## Usage

```bash
go run . input.txt output.txt
```

- `input.txt`: The file containing the text to process.
- `output.txt`: The file where the processed text will be written.

## File Overview

### main.go

- Entry point of the application.
- Handles command-line arguments, reads input, processes text, and writes output.
- Calls transformation functions in sequence: `processMarkers`, `handleUpTag`, `fixArticles`, and `removeSpaceEfficient`.

### articles.go

- Contains `fixArticles()` function.
- Changes "a"/"A" to "an"/"An" before words starting with vowels or 'h'.

### case.go

- Implements `handleUpTag()` and related helpers.
- Processes tags like `(up)`, `(low)`, `(cap)` to transform the case of preceding words.
- Supports tags with optional counts, e.g., `(up,3)` to affect multiple words.

### hexbin.go

- Contains `processMarkers()` function.
- Converts preceding words from hexadecimal or binary to decimal when `(hex)` or `(bin)` markers are found.

### punctuation.go

- Contains `removeSpaceEfficient()` function.
- Removes unnecessary spaces before punctuation, ensures proper comma spacing, and cleans up ellipsis formatting.

### utils.go

- Placeholder for shared utility functions (currently empty).

### go.mod

- Go module definition for the project.

### sample.txt

- Example input file for testing punctuation and formatting.

### result.txt

- Example output file showing processed results.

## Example

**Input (sample.txt):**
```
Punctuation tests are ... kinda boring ,what do you think ?
```

**Output (result.txt):**
```
Punctuation tests are... kinda boring, what do you think?
```

## How It Works

1. **Read Input:** Loads the text from the input file.
2. **Split Words:** Breaks text into words for processing.
3. **Apply Transformations:** Sequentially applies marker conversions, case changes, article corrections, and punctuation fixes.
4. **Write Output:** Saves the final processed text to the output file.

## License

This project is for educational purposes.
