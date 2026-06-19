# Day 11 - Rename a Field

Date: 2026-06-28
Timebox: 10 minutes

## Target Commands

- n
- N
- `ciw`
- careful replacement

## Challenge

Rename the `User` field to `Username` without changing unrelated words, then save the package files under `work/day11/`.

Start from these file(s):

- `challenges/day11/input/go.mod`
- `challenges/day11/input/profile.go`
- `challenges/day11/input/profile_test.go`

Expected output path(s):

- `work/day11/go.mod`
- `work/day11/profile.go`
- `work/day11/profile_test.go`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day11`.
2. Copy the input file(s) into `work/day11/`.
3. Edit only the files in `work/day11/`.
4. Run `go run ./cmd/check day11`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day11/`.
- [ ] `go run ./cmd/check day11` passes.
- [ ] I added a note to `notes/progress-log.md`.
