# Day 25 - Go Behavior Change

Date: 2026-07-12
Timebox: 10 minutes

## Target Commands

- reading code
- targeted edits
- running tests

## Challenge

Change `Clamp` so values below zero become zero and values above ten become ten. Save the package files under `work/day25/`.

Start from these file(s):

- `challenges/day25/input/go.mod`
- `challenges/day25/input/clamp.go`
- `challenges/day25/input/clamp_test.go`

Expected output path(s):

- `work/day25/go.mod`
- `work/day25/clamp.go`
- `work/day25/clamp_test.go`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day25`.
2. Copy the input file(s) into `work/day25/`.
3. Edit only the files in `work/day25/`.
4. Run `go run ./cmd/check day25`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day25/`.
- [ ] `go run ./cmd/check day25` passes.
- [ ] I added a note to `notes/progress-log.md`.
