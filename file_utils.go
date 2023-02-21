package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func parseInputFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// we're not writing to the file, so the iife-esque thing shouldn't
	// actually be needed, but it keeps editors from complaining about not
	// handling the error.
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// having an i counter sucks, but we do need to treat the first line in
	// the file like a special snowflake and I don't know of a better
	// option
	i := 0
	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// trim whitespace from start and end for sanity
		txt := strings.TrimSpace(scanner.Text())

		// skip commented lines that start with #
		if strings.HasPrefix(txt, "#") {
			continue
		}

		// skip any blank lines
		if 0 == len(txt) {
			continue
		}

		// first line should just be the gateway for the default route, so
		// we'll add the default network
		if 0 == i {
			lines = append(lines, "0.0.0.0/0 "+txt)
		} else {
			lines = append(lines, txt)
		}
		i += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
