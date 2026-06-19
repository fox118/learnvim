# Day 03 - Repair a Paragraph

Date: 2026-06-20
Timebox: 10 minutes

## Target Commands

- w
- b
- e
- dw
- `cw`
- yy
- `p`

## Challenge

Fix the paragraph, including repeated words, capitalization, and punctuation, then save it to `work/day03/paragraph.md`. Yank the final first sentence into `work/day03/yank.txt`.

Start from these file(s):

- `challenges/day03/input/paragraph.md`

Expected output path(s):

- `work/day03/paragraph.md`
- `work/day03/yank.txt`

## Suggested Workflow

1. Create the work folder: `mkdir -p work/day03`.
2. Copy the input file(s) into `work/day03/`.
3. Edit only the files in `work/day03/`.
4. Run `go run ./cmd/check day03`.
5. If the checker reports a mismatch, reopen your work file and fix it.

## Done Checklist

- [ ] I practiced for 10 focused minutes.
- [ ] I avoided arrow keys unless stuck.
- [ ] I saved the requested file(s) under `work/day03/`.
- [ ] `go run ./cmd/check day03` passes.
- [ ] I added a note to `notes/progress-log.md`.
