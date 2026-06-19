package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	day, err := dayArg(os.Args[1:])
	if err != nil {
		return err
	}
	if !validDay(day) {
		return fmt.Errorf("invalid day %q; use a value like day01", day)
	}

	expectedDir := filepath.Join("challenges", day, "expected")
	workDir := filepath.Join("work", day)

	if _, err := os.Stat(expectedDir); err != nil {
		return fmt.Errorf("expected files for %s were not found: %w", day, err)
	}
	if _, err := os.Stat(workDir); err != nil {
		return fmt.Errorf("work folder %q was not found; create it and copy the input files first", workDir)
	}

	expectedFiles, err := regularFiles(expectedDir)
	if err != nil {
		return err
	}
	workFiles, err := regularFiles(workDir)
	if err != nil {
		return err
	}

	var failures []string
	for _, rel := range expectedFiles {
		expectedPath := filepath.Join(expectedDir, rel)
		workPath := filepath.Join(workDir, rel)

		expected, err := os.ReadFile(expectedPath)
		if err != nil {
			return err
		}
		actual, err := os.ReadFile(workPath)
		if err != nil {
			failures = append(failures, fmt.Sprintf("missing %s", workPath))
			continue
		}
		if !bytes.Equal(normalizeNewlines(actual), normalizeNewlines(expected)) {
			failures = append(failures, fmt.Sprintf("content mismatch in %s", workPath))
		}
	}

	if len(failures) > 0 {
		sort.Strings(failures)
		return fmt.Errorf("check failed for %s:\n- %s", day, strings.Join(failures, "\n- "))
	}

	if containsGoFiles(workFiles) {
		if err := runGoTest(workDir); err != nil {
			return err
		}
	}

	fmt.Printf("PASS %s\n", day)
	return nil
}

func dayArg(args []string) (string, error) {
	switch len(args) {
	case 0:
		return todayDay(), nil
	case 1:
		return args[0], nil
	default:
		return "", errors.New("usage: go run ./cmd/check [day01]")
	}
}

func todayDay() string {
	entries, err := os.ReadDir("work")
	if err != nil {
		return "day01"
	}

	day := "day01"
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() && validDay(name) && name > day {
			day = name
		}
	}
	return day
}

func validDay(day string) bool {
	if len(day) != len("day01") || !strings.HasPrefix(day, "day") {
		return false
	}
	return day[3] >= '0' && day[3] <= '9' && day[4] >= '0' && day[4] <= '9'
}

func regularFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files = append(files, filepath.ToSlash(rel))
		return nil
	})
	sort.Strings(files)
	return files, err
}

func normalizeNewlines(content []byte) []byte {
	return bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
}

func containsGoFiles(files []string) bool {
	for _, file := range files {
		if strings.HasSuffix(file, ".go") {
			return true
		}
	}
	return false
}

func runGoTest(workDir string) error {
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = workDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go test failed in %s:\n%s", workDir, strings.TrimSpace(string(output)))
	}
	return nil
}
