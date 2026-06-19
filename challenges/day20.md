# Day 20 - Search and Replace Review

Date: 2026-07-07
Timebox: 10 minutes

## Target Commands

- controlled substitution
- confirmation
- n
- N

## Challenge

Rename the product throughout active release-note content only, leaving historical notes unchanged. Save the result to `work/day20/release.md`.

Start from these file(s):

- `challenges/day20/input/release.md`

Expected output path(s):

- `work/day20/release.md`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day20`.
2. Copy the input file(s) into `work/day20/`.
3. Edit only the files in `work/day20/`.
4. Run `go run ./cmd/check day20`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day20/`.
- [ ] `go run ./cmd/check day20` passes.
- [ ] I added a note to `notes/progress-log.md`.
