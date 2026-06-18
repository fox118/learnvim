#!/usr/bin/env bash
set -euo pipefail
start="2026-06-18"
today=$(date +%F)
day=$(( ( $(date -d "$today" +%s) - $(date -d "$start" +%s) ) / 86400 + 1 ))
if [ "$day" -lt 1 ]; then day=1; fi
if [ "$day" -gt 28 ]; then day=28; fi
printf 'challenges/day%02d.md\n' "$day"
