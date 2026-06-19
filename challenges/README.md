# Challenge Design

Each day now works like a timeboxed editing task with feedback. The challenge files describe what to fix, `input/` contains the broken starting files, and `expected/` contains the checker answers.

## Challenge Shape

Each challenge has:

- `challenges/dayXX.md`: the daily instructions.
- `challenges/dayXX/input/`: broken files to copy into your work folder.
- `challenges/dayXX/expected/`: expected files used by the checker.
- `work/dayXX/`: your edited files. This folder is ignored by git.

Daily workflow:

1. Open `challenges/day03.md`.
2. Create `work/day03`.
3. Copy the file(s) from `challenges/day03/input/` into `work/day03/`.
4. Edit the work files using the day’s target Vim commands.
5. Run `./check.sh`, or `./check.sh day03` to check a specific day.

The checker compares `work/dayXX/` against `challenges/dayXX/expected/`. If the work folder contains Go files, it also runs `go test ./...` inside that folder.

## Daily Challenges

### Day 01: Fix a Short Note

Fix a Markdown note with duplicated words, missing punctuation, inconsistent headings, and rough bullet text.

Focus: Normal mode, Insert mode, `h/j/k/l`, `x`, `i`, `Esc`, `:w`.

### Day 02: Clean a Todo List

Reorder a todo list, delete completed junk items, and add missing handoff steps. Store the fixed list in a new file.

Focus: line movement, `dd`, `p`, `u`, save/quit.

### Day 03: Repair a Paragraph

Fix a prose paragraph with repeated words, misspellings, punctuation gaps, and capitalization issues. Yank the final first sentence into a separate file.

Focus: word movement, `w`, `b`, `e`, `dw`, `cw`, `yy`, `p`.

### Day 04: Normalize Headings

Edit a Markdown file where headings use inconsistent levels. Convert the structure into one title, three sections, and clean bullet lists.

Focus: start/end of line, `0`, `$`, `I`, `A`, `dd`, `p`.

### Day 05: Fix CSV-like Text

Clean a plain-text list of names and emails. Remove duplicate rows and fix spacing around commas.

Focus: repeatable edits, `.`, `f`, `;`, small substitutions.

### Day 06: Extract Important Lines

From a noisier log-style text file, copy only the error lines into a requested output file while leaving the original file unchanged.

Focus: search, `/`, `n`, Visual mode basics, yank, paste.

### Day 07: Week 1 Review

Complete a mixed Markdown cleanup that uses deleting, changing, yanking, pasting, and undoing across a fuller review note. The expected file checks the final content, not which commands were used.

Focus: survival commands from the first week.

### Day 08: Fix a Go Function Name

Edit a small Go file where function names and call sites do not match.

Focus: search, word motions, `cw`, `*`, `n`.

### Day 09: Add a Missing Return

Update a simple Go function with a missing return value. Keep the change narrow and save the corrected file to the target path.

Focus: insert at line ends, `$`, `A`, braces navigation by search.

### Day 10: Repair a Go Test

Fix table-driven Go test cases where expected values are wrong or labels are unclear.

Focus: moving by words and lines, changing quoted text, repeat with `.`.

### Day 11: Rename a Field

Rename one struct field and update its usages in a small Go file. Avoid changing unrelated words that happen to look similar.

Focus: search, `n`, `N`, `ciw`, careful replacement.

### Day 12: Format a Markdown Changelog

Turn a messy changelog into consistent sections and bullet points. Yank the final latest-version section to a separate file.

Focus: paragraphs, line operations, Visual selection, yank.

### Day 13: Fix Quoted Strings

Edit text with broken quotes, missing punctuation, and inconsistent capitalization. The expected result is exact text.

Focus: text objects like `ci"`, `di"`, and repeated edits.

### Day 14: Week 2 Review

Complete a mixed challenge with one Markdown file and one Go file. Save both outputs to the required paths.

Focus: operators plus motions, text objects, search, repeat.

### Day 15: Replace Config Values

Update a config-like file by changing only specific active keys. Do not touch commented examples.

Focus: search precision, `:s`, ranges, avoiding accidental replacements.

### Day 16: Fix a Markdown Table Without Seeing the Answer

Repair values across multiple columns in a Markdown table.

Focus: columns, `f`, `t`, `r`, `cw`.

### Day 17: Add a Go Error Case

Add a missing branch to a small Go function and update the related test case. Checker uses `go test`.

Focus: copying nearby patterns, `yap`, `p`, targeted edits.

### Day 18: Extract a Go Helper

Move repeated logic into a helper function in a tiny Go file. This is still a small edit, not a full refactor.

Focus: yanking blocks, pasting, changing identifiers.

### Day 19: Fix Imports

Edit a Go file with unused and missing imports.

Focus: deleting lines, inserting lines, keeping file structure clean.

### Day 20: Search and Replace Review

Fix a document where one product name changed, but only in active content and not in historical notes.

Focus: controlled substitution, confirmation, `n`, `N`.

### Day 21: Week 3 Review

Complete a realistic repo-style task: fix one note, one config snippet, and one small Go function. Save all outputs under one day folder.

Focus: switching files, search, text objects, repeat.

### Day 22: VS Code Navigation Task

Use VS Code file search to find a broken file among a small decoy, then fix it with Vim motions. The challenge should mention the file name pattern but not the exact path.

Focus: VS Code search plus Vim editing.

### Day 23: Terminal and Checker Task

Edit a file, save it to the target path, then run a checker from the integrated terminal. The challenge should require reading the checker output and making any remaining correction.

Focus: edit, save, terminal loop.

### Day 24: Multi-file Markdown Cleanup

Apply the same cleanup rules to two short Markdown files. A future checker verifies both outputs.

Focus: buffers/tabs in VS Code, repeatable command patterns.

### Day 25: Go Behavior Change

Change a tiny Go function to meet a new requirement described in prose. Tests should define the expected behavior.

Focus: reading code, targeted edits, running tests.

### Day 26: Yank Required Content

Fix a file and also save a specific extracted section into a second file. The checker verifies both the final file and extracted content.

Focus: Visual mode, `y`, registers if desired, paste to a new file.

### Day 27: Final Mixed Repair

Complete a larger but still timeboxed mixed challenge with broken Markdown, broken Go code, and a required extracted summary.

Focus: combining the strongest habits from the month.

### Day 28: Personal Workflow Challenge

Update the personal cheatsheet with commands actually used during the month, then fix one final small file. The checker verifies the final file, while the cheatsheet remains personal.

Focus: sustainable Vim-in-VS-Code workflow.

## Checker Notes

The Go checker can:

- Compare edited files against golden expected files.
- Check that required output paths exist.
- Verify extracted yank files.
- Run `go test` for Go challenges.
- Print clear, friendly failure messages.

For now, text validation is intentionally exact apart from line ending differences. If exact matching feels too strict later, the checker can grow day-specific validation rules.
