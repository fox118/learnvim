# Day 24 - Multi-file Markdown Cleanup

Date: 2026-07-11
Timebox: 10 minutes

## Target Commands

- buffers
- repeatable command patterns

## Challenge

Apply the same cleanup rules to both Markdown files and save them under `work/day24/`.

Start from these file(s):

- `challenges/day24/input/alpha.md`
- `challenges/day24/input/beta.md`

Expected output path(s):

- `work/day24/alpha.md`
- `work/day24/beta.md`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day24`.
2. Copy the input file(s) into `work/day24/`.
3. Edit only the files in `work/day24/`.
4. Run `go run ./cmd/check day24`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day24/`.
- [ ] `go run ./cmd/check day24` passes.
- [ ] I added a note to `notes/progress-log.md`.
