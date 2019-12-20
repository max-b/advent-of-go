package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadInputToLines takes a path and reads it line by line into an array
func ReadInputToLines(path string) ([]string, error) {
	var lines []string

	f, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("reading standard input:", err)
		return lines, err
	}

	if err := f.Close(); err != nil {
		return lines, err
	}

	return lines, nil
}
