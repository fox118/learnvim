# Day 19 - Fix Imports

Date: 2026-07-06
Timebox: 10 minutes

## Target Commands

- deleting lines
- inserting lines
- file structure

## Challenge

Fix the imports so the code builds, then save the package files under `work/day19/`.

Start from these file(s):

- `challenges/day19/input/go.mod`
- `challenges/day19/input/format.go`
- `challenges/day19/input/format_test.go`

Expected output path(s):

- `work/day19/go.mod`
- `work/day19/format.go`
- `work/day19/format_test.go`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day19`.
2. Copy the input file(s) into `work/day19/`.
3. Edit only the files in `work/day19/`.
4. Run `go run ./cmd/check day19`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day19/`.
- [ ] `go run ./cmd/check day19` passes.
- [ ] I added a note to `notes/progress-log.md`.
