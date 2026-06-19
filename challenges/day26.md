# Day 26 - Yank Required Content

Date: 2026-07-13
Timebox: 10 minutes

## Target Commands

- Visual mode
- y
- registers
- paste to a new file

## Challenge

Fix the file at `work/day26/guide.md` and save only the Commands section to `work/day26/commands.md`.

Start from these file(s):

- `challenges/day26/input/guide.md`

Expected output path(s):

- `work/day26/guide.md`
- `work/day26/commands.md`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day26`.
2. Copy the input file(s) into `work/day26/`.
3. Edit only the files in `work/day26/`.
4. Run `go run ./cmd/check day26`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day26/`.
- [ ] `go run ./cmd/check day26` passes.
- [ ] I added a note to `notes/progress-log.md`.
