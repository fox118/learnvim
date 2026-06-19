# Day 15 - Replace Config Values

Date: 2026-07-02
Timebox: 10 minutes

## Target Commands

- search precision
- `:s`
- ranges

## Challenge

Update only the active config values, leaving commented examples unchanged, and save the result to `work/day15/app.conf`.

Start from these file(s):

- `challenges/day15/input/app.conf`

Expected output path(s):

- `work/day15/app.conf`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day15`.
2. Copy the input file(s) into `work/day15/`.
3. Edit only the files in `work/day15/`.
4. Run `go run ./cmd/check day15`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day15/`.
- [ ] `go run ./cmd/check day15` passes.
- [ ] I added a note to `notes/progress-log.md`.
