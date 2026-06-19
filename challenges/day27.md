# Day 27 - Final Mixed Repair

Date: 2026-07-14
Timebox: 10 minutes

## Target Commands

- month review
- Markdown
- Go
- extraction

## Challenge

Fix the Markdown and Go files, then extract the summary to `work/day27/summary.md`.

Start from these file(s):

- `challenges/day27/input/report.md`
- `challenges/day27/input/go.mod`
- `challenges/day27/input/count.go`
- `challenges/day27/input/count_test.go`

Expected output path(s):

- `work/day27/report.md`
- `work/day27/summary.md`
- `work/day27/go.mod`
- `work/day27/count.go`
- `work/day27/count_test.go`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day27`.
2. Copy the input file(s) into `work/day27/`.
3. Edit only the files in `work/day27/`.
4. Run `go run ./cmd/check day27`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day27/`.
- [ ] `go run ./cmd/check day27` passes.
- [ ] I added a note to `notes/progress-log.md`.
