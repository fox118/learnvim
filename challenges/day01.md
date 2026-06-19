# Day 01 - Fix a Short Note

Date: 2026-06-18
Timebox: 10 minutes

## Target Commands

- Normal mode
- Insert mode
- `h/j/k/l`
- `x`
- `i`
- `Esc`
- `:w`

## Challenge

Fix the Markdown note by cleaning headings, sentences, and bullets, then save the corrected file to `work/day01/note.md`.

The starting note has repeated words, a typo, missing punctuation, inconsistent capitalization, and one line plus one bullet that should be removed.

Start from these file(s):

- `challenges/day01/input/note.md`

Expected output path(s):

- `work/day01/note.md`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day01`.
2. Copy the input file(s) into `work/day01/`.
3. Edit only the files in `work/day01/`.
4. Run `go run ./cmd/check day01`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day01/`.
- [ ] `go run ./cmd/check day01` passes.
- [ ] I added a note to `notes/progress-log.md`.
