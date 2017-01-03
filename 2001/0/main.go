package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Solution(reader, writer)
}

//Solution contains entire solution to the problem, nothing is done in main.
func Solution(in *bufio.Reader, out *bufio.Writer) {
	scanner := bufio.NewScanner(in)
	people := make(map[string]bool)
	// Part one, read in names
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		} else {
			people[scanner.Text()] = true
		}
	}
	// Part two read in luggage info
	for scanner.Scan() {
		val := people[scanner.Text()]
		if val {
			out.WriteString("OK\n")
		} else {
			out.WriteString("WARNING: passenger not on board\n")
		}
	}
	out.Flush()
}
