# Day 22 - VS Code Navigation Task

Date: 2026-07-09
Timebox: 10 minutes

## Target Commands

- VS Code search
- Vim editing

## Challenge

Use file search to find the broken note under `input/`, ignore the decoy file, and save the fixed file to `work/day22/found-note.md`.

Start from these file(s):

- `challenges/day22/input/archive/readme.md`
- `challenges/day22/input/deep/found-note.md`

Expected output path(s):

- `work/day22/found-note.md`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day22`.
2. Copy the input file(s) into `work/day22/`.
3. Edit only the files in `work/day22/`.
4. Run `go run ./cmd/check day22`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day22/`.
- [ ] `go run ./cmd/check day22` passes.
- [ ] I added a note to `notes/progress-log.md`.
