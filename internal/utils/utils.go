package utils

import (
	"bufio"
	"os"
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