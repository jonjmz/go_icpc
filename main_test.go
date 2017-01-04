package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

// Tests the solution
func TestSolution(t *testing.T) {
	// Prepare input and output
	inFile, err := os.Open("in")
	if err != nil {
		t.Error(err)
	}
	defer inFile.Close()
	resultBuffer := new(bytes.Buffer)
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(resultBuffer)

	// Solve Problem
	Solution(reader, writer)

	// Prepare Buffers for testing
	desiredFile, err := os.Open("out")
	if err != nil {
		t.Error(err)
	}
	defer desiredFile.Close()
	desiredScanner := bufio.NewScanner(desiredFile)
	resultScanner := bufio.NewScanner(resultBuffer)

	// Compare output
	for {
		moreDesired := desiredScanner.Scan()
		moreResult := resultScanner.Scan()
		if moreDesired || moreResult {
			if desiredScanner.Text() != resultScanner.Text() {
				t.Errorf("%-20s VS %s", desiredScanner.Text(), resultScanner.Text())
			}
		} else {
			break
		}
	}

}
