package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Solution(reader, writer)
}

//Solution contains entire solution to the problem, nothing is done in main.
func Solution(in *bufio.Reader, out *bufio.Writer) {
	scanner := bufio.NewScanner(in)
	re := regexp.MustCompile(`([\d-]{9}(\d\d)) (\d) (\d+) (\d+)`)
	releaseDays := map[int]string{
		0: "July 23",
		1: "July 30",
		2: "August 6",
		3: "August 13",
		4: "August 20",
		5: "August 27",
		6: "September 3",
		7: "September 10",
		8: "September 17",
		9: "September 24",
	}
	var check float64
	var week int
	releaseCounts := make([]int, 10)
	releaseTotals := make([]float64, 10)
	for scanner.Scan() {
		parts := re.FindStringSubmatch(scanner.Text())
		ssn := parts[1]
		last2, _ := strconv.Atoi(parts[2])
		filing, _ := strconv.Atoi(parts[3])
		income, _ := strconv.Atoi(parts[4])
		liability, _ := strconv.Atoi(parts[5])

		switch {
		case filing == 1 || filing == 3:
			check = math.Min(math.Min(0.05*float64(income), float64(liability)), 300.0)
		case filing == 4:
			check = math.Min(math.Min(0.05*float64(income), float64(liability)), 500.0)
		case filing == 2 || filing == 5:
			check = math.Min(math.Min(0.05*float64(income), float64(liability)), 600.0)
		}
		out.WriteString(fmt.Sprintf("%s  $%-.2f\n", ssn, check))

		week = int(math.Floor(float64(last2) / 10.0))

		releaseCounts[week]++
		releaseTotals[week] += check
	}
	for i, releaseDate := range releaseDays {
		if releaseCounts[i] == 0 {
			continue
		}
		out.WriteString(fmt.Sprintf("%d  $%-.2f %s\n", releaseCounts[i], releaseTotals[i], releaseDate))
	}
	// Solution goes here
	out.Flush()
}
