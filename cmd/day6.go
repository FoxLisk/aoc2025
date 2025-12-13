/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"
)

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run_day6,
}

func init() {
	rootCmd.AddCommand(day6Cmd)

}

func run_day6(cmd *cobra.Command, args []string) {
	part2_day6()
}

func part1_day6() {
	lines, err := utils.ReadLines("inputs/6")
	utils.Check(err)
	cols := make([][]int, len(strings.Fields(lines[0])))
	// fmt.Println("cols:", cols)
	ops := strings.Fields(lines[len(lines)-1])
	for _, line := range lines[:len(lines)-1] {
		splat := strings.Fields(line)
		for i, s := range splat {
			parsed, err := strconv.Atoi(s)
			utils.Check(err)
			cols[i] = append(cols[i], parsed)
		}
	}
	// fmt.Println("cols:", cols)
	// fmt.Println("ops:", ops)a
	thePassword := 0
	for i, op := range ops {
		var acc int
		var reducer func(int, int) int
		if op == "+" {
			acc = 0
			reducer = func(a int, b int) int { return a + b }
		} else if op == "*" {
			acc = 1
			reducer = func(a int, b int) int { return a * b }
		}
		val := utils.Reduce(cols[i], acc, reducer)
		thePassword += val
		// fmt.Println("doing", op, "to ", cols[i], " - got ", val)
	}
	fmt.Println("The password is ", thePassword)

}

type problem struct {
	op   rune
	nums [][]byte // maybe?
}

func (p *problem) display() {
	fmt.Println(string(p.op))
	for _, col := range p.nums {
		s := ""
		for _, c := range col {
			s = s + string(c)
		}
		fmt.Println(s)
	}
}

func (p *problem) calculate() int {
	nums := make([]int, 0, len(p.nums))
	for _, n := range p.nums {
		s := strings.Trim(string(n), " ")
		if s == "" {
			continue
		}
		parsed, err := strconv.Atoi(s)
		utils.Check(err)
		nums = append(nums, parsed)
	}
	if p.op == '+' {
		return utils.Reduce(nums, 0, func(a, b int) int { return a + b })
	} else if p.op == '*' {
		return utils.Reduce(nums, 1, func(a, b int) int { return a * b })
	} else {
		panic("uh uh")
	}
}

func part2_day6() {
	lines, err := utils.ReadLines("inputs/6")
	utils.Check(err)
	// cols := make([][]int, len(strings.Fields(lines[0])))
	ops := lines[len(lines)-1]
	var currentProblem *problem
	var problems []*problem

	for i, c := range ops {
		if c != ' ' {
			if currentProblem != nil {
				problems = append(problems, currentProblem)
			}
			currentProblem = &problem{
				op:   c,
				nums: make([][]byte, 0),
			}
		}
		var col []byte
		for _, line := range lines[:len(lines)-1] {
			col = append(col, line[i])
		}
		fmt.Println("Col", i, ": ", string(col))
		currentProblem.nums = append(currentProblem.nums, col)
	}
	problems = append(problems, currentProblem)
	thePassword := 0
	for _, p := range problems {
		p.display()
		val := p.calculate()
		fmt.Println(" =", val)
		thePassword += val
	}
	fmt.Println("password:", thePassword)
}
