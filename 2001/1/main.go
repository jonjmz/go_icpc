package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	for scanner.Scan() {
		interval := evaluate(scanner.Text())
		out.WriteString(interval.String())
		out.WriteString("\n")
		// fmt.Fprintf(out, "Devision by zero\n")
	}
	out.Flush()
}
func evaluate(text string) (r Interval) {
	var a, b Interval
	var operand rune
	var i, start, end int
	runes := []rune(text)
	// Part one Left operator
	// Expecting spaces, unary operators, or Intervals
	for i = 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			// Space is ignored
			continue
		} else if runes[i] == '-' {
			// Unary negate, next should be '['
			start = i + 1
			for end = start; end < len(runes); end++ {
				if runes[end] == ']' {
					break
				}
			}
			a.fromString(string(runes[start : end+1]))
			r = negate(a)
			return
		} else if runes[i] == '[' {
			start = i
			for end = start; end < len(runes); end++ {
				if runes[end] == ']' {
					break
				}
			}
			a.fromString(string(runes[start : end+1]))
			i = end + 1
			break
		}
	}
	// Part two Operator
	// Expecting spaces, operators
	for ; i < len(runes); i++ {
		if runes[i] == ' ' {
			// Space is ignored
			continue
		} else {
			operand = runes[i]
			break
		}
	}
	// Part Three Right interval
	// Expecting spaces, or Interval
	for ; i < len(runes); i++ {
		if runes[i] == ' ' {
			// Space is ignored
			continue
		} else if runes[i] == '[' {
			start = i
			for end = start; end < len(runes); end++ {
				if runes[end] == ']' {
					break
				}
			}
			b.fromString(string(runes[start : end+1]))
			break
		}
	}
	fmt.Printf("%v %c %v", a, operand, b)
	switch operand {
	case '+':
		r = add(a, b)
	case '-':
		r = subtract(a, b)
	case '*':
		r = multiply(a, b)
	case '/':
		r = divide(a, b)
	}
	fmt.Printf(" = %v\n", r)
	return
}

func subtract(a, b Interval) (r Interval) {
	r.a = a.min() - b.max()
	r.b = a.max() - b.min()
	return
}
func add(a, b Interval) (r Interval) {
	r.a = a.min() + b.min()
	r.b = a.max() + b.max()
	return
}
func multiply(a, b Interval) (r Interval) {
	r.a = math.Min(math.Min(
		a.max()*b.min(),
		a.min()*b.max()),
		a.min()*b.min())
	r.b = math.Max(math.Max(
		a.max()*b.min(),
		a.min()*b.max()),
		a.max()*b.max())
	return
}
func divide(a, b Interval) (r Interval) {
	r.a = math.Min(math.Min(
		a.max()/b.min(),
		a.min()/b.max()),
		a.min()/b.min())
	r.b = math.Max(math.Max(
		a.max()/b.min(),
		a.min()/b.max()),
		a.max()/b.max())
	return
}
func negate(a Interval) (r Interval) {
	r.a = -a.min()
	r.b = -a.max()
	return
}

// Interval represents an interval in interval arithmatic
type Interval struct {
	a float64
	b float64
}

func (i *Interval) min() float64 {
	return math.Min(i.a, i.b)
}
func (i *Interval) max() float64 {
	return math.Max(i.a, i.b)
}

// ToString converts Interval to string
func (i *Interval) String() string {
	return fmt.Sprintf("[%.3f,%.3f]", i.min(), i.max())
}

// FromString parses an interval from string
func (i *Interval) fromString(text string) {
	nums := strings.Split(text[1:len(text)-1], ",")
	i.a, _ = strconv.ParseFloat(nums[0], 64)
	i.b, _ = strconv.ParseFloat(nums[1], 64)
	return
}
