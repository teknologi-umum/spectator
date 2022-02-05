package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// loadEnvironment is a simple function to load environment variables.
// We have two source: one coming from the machine's environment, and
// one coming from the .env file.
//
// The job for providing default values falls back to the main.go function.
func loadEnvironment() error {
	// Check if there is an .env file available on the current directory
	// and load it.
	file, err := os.Open(".env")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}
	defer file.Close()

	// Read line by line from the file, and store it
	// in the environment variable.
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var env = line[:strings.Index(line, "=")]
		var value = line[strings.Index(line, "=")+1:]
		os.Setenv(env, value)
	}

	return nil
}
