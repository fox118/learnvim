# Day 06 - Extract Important Lines

Date: 2026-06-23
Timebox: 10 minutes

## Target Commands

- `/`
- n
- Visual mode
- yank
- paste

## Challenge

Copy only the error lines from a noisier `log.txt` into `work/day06/errors.txt`. Leave the original log alone if you want.

Start from these file(s):

- `challenges/day06/input/log.txt`

Expected output path(s):

- `work/day06/errors.txt`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day06`.
2. Copy the input file(s) into `work/day06/`.
3. Edit only the files in `work/day06/`.
4. Run `go run ./cmd/check day06`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day06/`.
- [ ] `go run ./cmd/check day06` passes.
- [ ] I added a note to `notes/progress-log.md`.
