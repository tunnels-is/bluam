package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

// takes a string and checks for a number of prefixes and suffixes then removes them
// and returns the string with trimed spaces
func toPlainDomain(s string) string {
	// Check prefixes 1st then another pass for suffixes
	switch {
	case strings.HasPrefix(s, ":"):
		return ""
	case strings.HasPrefix(s, "["):
		return ""
	case strings.HasPrefix(s, "#"):
		return ""
	case strings.HasPrefix(s, "!"):
		return ""
	case strings.HasPrefix(s, "*"):
		s = strings.TrimSpace(s[1:])
	case strings.HasPrefix(s, "||"):
		s = strings.TrimSpace(s[2:])
	case strings.HasPrefix(s, "0.0.0.0"):
		s = strings.TrimSpace(s[len("0.0.0.0"):])
	case strings.HasPrefix(s, "127.0.0.1"):
		s = strings.TrimSpace(s[len("127.0.0.1"):])
	default:
		return strings.TrimSpace(s)
	}

	if strings.HasSuffix(s, "^") {
		s = strings.TrimSpace(s[0 : len(s)-1])
	}

	return s
}

// take a directory path and returns false if doesn't exist, true if it does
func CheckPathExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// Finds the executalbe path
func FindExePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return ex
}

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
