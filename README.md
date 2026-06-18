# Learn Vim Challenges

A 28-day, 10-minutes-per-day Vim practice repository for the Linear project "Learn Vim".

Start date: 2026-06-18
End date: 2026-07-15

This repo is designed to be edited with Vim or VS Code + the VSCodeVim extension. Each daily challenge is intentionally small: open the day's markdown file, follow the instructions, edit the practice text, and record one observation in `notes/progress-log.md`.

## Daily workflow

1. Open today's file, for example:
   - Terminal Vim: `vim challenges/day01.md`
   - VS Code: open the repository, enable VSCodeVim, then edit `challenges/day01.md`.
2. Spend 10 focused minutes only.
3. Avoid arrow keys unless you are truly stuck.
4. Use the target commands listed in the challenge.
5. Add a short note to `notes/progress-log.md`.
6. Commit if you want a visible practice history:
   `git add . && git commit -m "practice: complete day 01 vim challenge"`

## Weekly focus

- Week 1: core survival skills, `vimtutor`, movement, insert/normal mode, save/quit, delete/change/yank.
- Week 2: Vim grammar, operators + motions, text objects, search, repeat, substitution.
- Week 3: VS Code + Vim, keybinding conflicts, IDE escape hatches, real editing sessions.
- Week 4: integration, speed, personal cheatsheet, sustainable workflow.

## VS Code setup

Install the VSCodeVim extension:
https://marketplace.visualstudio.com/items?itemName=vscodevim.vim

Suggested starter settings are in `.vscode/settings.json`. Treat them as a learning baseline, not final truth.

## Definition of done

By the end, you should be able to:

- Navigate and edit common text without arrow keys.
- Use operator + motion combinations like `dw`, `ci"`, `daw`, `yap`, `gg`, `G`, `/`, `n`, and `.`.
- Use Vim bindings inside VS Code without losing access to normal IDE features.
- Maintain a personal Vim-in-VS-Code cheatsheet based on commands you actually use.
