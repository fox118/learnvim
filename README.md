# Learn Vim Challenges

A 28-day, 10-minutes-per-day Vim practice repository for the Linear project "Learn Vim".

Start date: 2026-06-18
End date: 2026-07-15

This repo is designed to be edited with Vim or VS Code + the VSCodeVim extension. Each daily challenge is still timeboxed, but the fixtures include enough repeated edits to build muscle memory: read the day's instructions, copy the broken input into `work/dayXX/`, fix it, run the checker, and record one observation in `notes/progress-log.md`.

## Daily workflow

1. Open today's file, for example:
   - Terminal Vim: `vim challenges/day01.md`
   - VS Code: open the repository, enable VSCodeVim, then edit `challenges/day01.md`.
2. Create the work folder, for example: `mkdir -p work/day01`.
3. Copy the input file(s) from `challenges/day01/input/` into `work/day01/`.
4. Spend 10 focused minutes fixing only the files under `work/day01/`.
5. Run the checker: `./check.sh`.
6. Add a short note to `notes/progress-log.md`.
7. Commit if you want a visible practice history:
   `git add . && git commit -m "practice: complete day 01 vim challenge"`

## Weekly focus

- Week 1: core survival skills with checkable Markdown and text cleanup tasks.
- Week 2: Vim grammar, operators + motions, text objects, search, repeat, and first Go edits.
- Week 3: search, substitution, multi-file work, and Go test feedback.
- Week 4: VS Code integration, terminal checker loops, extraction/yank practice, and final review.

## VS Code setup

Install the VSCodeVim extension:
https://marketplace.visualstudio.com/items?itemName=vscodevim.vim

Suggested starter settings are in `.vscode/settings.json`. Treat them as a learning baseline, not final truth.

## Definition of done

By the end, you should be able to:

- Navigate and edit common text without arrow keys.
- Use operator + motion combinations like `dw`, `ci"`, `daw`, `yap`, `gg`, `G`, `/`, `n`, and `.`.
- Fix small Markdown, text, and Go files against expected output.
- Use Vim bindings inside VS Code without losing access to normal IDE features.
- Maintain a personal Vim-in-VS-Code cheatsheet based on commands you actually use.
