# Day 08 - Fix a Go Function Name

Date: 2026-06-25
Timebox: 10 minutes

## Target Commands

- search
- word motions
- `cw`
- `*`
- n

## Challenge

Fix the mismatched Go function name and save the package files under `work/day08/`.

Start from these file(s):

- `challenges/day08/input/go.mod`
- `challenges/day08/input/greet.go`
- `challenges/day08/input/greet_test.go`

Expected output path(s):

- `work/day08/go.mod`
- `work/day08/greet.go`
- `work/day08/greet_test.go`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day08`.
2. Copy the input file(s) into `work/day08/`.
3. Edit only the files in `work/day08/`.
4. Run `go run ./cmd/check day08`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day08/`.
- [ ] `go run ./cmd/check day08` passes.
- [ ] I added a note to `notes/progress-log.md`.
