package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// i'm just ignoring size because it's too much typing
	out := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}

	return out, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// this is hilarious
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type IdRange struct {
	Start uint64
	End   uint64
}

func IdRangeFromString(s string) (*IdRange, error) {
	parts := strings.Split(s, "-")
	var start, end uint64
	var err error
	if start, err = strconv.ParseUint(parts[0], 10, 64); err != nil {
		return nil, err
	}
	if end, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
		return nil, err
	}
	return &IdRange{Start: start, End: end}, nil
}
