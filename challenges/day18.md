# Day 18 - Extract a Go Helper

Date: 2026-07-05
Timebox: 10 minutes

## Target Commands

- yanking blocks
- pasting
- changing identifiers

## Challenge

Extract the repeated trim/lowercase logic into `cleanName`, then save the package files under `work/day18/`.

Start from these file(s):

- `challenges/day18/input/go.mod`
- `challenges/day18/input/names.go`
- `challenges/day18/input/names_test.go`

Expected output path(s):

- `work/day18/go.mod`
- `work/day18/names.go`
- `work/day18/names_test.go`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day18`.
2. Copy the input file(s) into `work/day18/`.
3. Edit only the files in `work/day18/`.
4. Run `go run ./cmd/check day18`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day18/`.
- [ ] `go run ./cmd/check day18` passes.
- [ ] I added a note to `notes/progress-log.md`.
