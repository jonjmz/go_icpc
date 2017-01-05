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
	// Solution goes here
	out.Flush()
}
