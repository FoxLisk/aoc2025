/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"

	combinations "github.com/mxschmitt/golang-combinations"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "A brief description of your command",
	Long:  `asdf`,
	Run: func(cmd *cobra.Command, args []string) {
		raw, err := utils.ReadLines("inputs/10")
		utils.Check(err)
		machines := parse_machines(raw)
		part1_day10(machines)
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}

func part1_day10(machines []machine) {
	password := 0
	for _, m := range machines {
		soln := solve_machine_p1(m)
		password += soln
	}
	fmt.Println("Password:", password)
}

func solve_machine_p1(m machine) int {
	// pressing a button twice is never useful because it just perfectly undoes the first press.
	// so the only options are to press a button once or zero times
	// this means we can find the optimal solution by checking all possible combinations of button presses
	// that's 2^n, but n here is the number of buttons, and n <= 13, which is only like 8k, which is tiny

	// how do we represent this? naively it seems like you need all bitstrings, but i think it might be easier
	// to construct lists of indices instead of bitstrings?

	// i tried to write this myself but i sucked at it and gave up
	// for _, c := range combinations.combinations() {

	// }
	indices := make([]int, 0, len(m.buttons))
	for i := range m.buttons {
		indices = append(indices, i)
	}
	// having a 2nd parameter here that you set to 0 to mean all sucks a lot
	// but i think that is because golang sucks?
	best_soln := 100000
	for _, c := range combinations.Combinations(indices, 0) {
		if test_button_combination(m, c) {
			best_soln = min(best_soln, len(c))
		}
	}
	return best_soln
}

func test_button_combination(m machine, button_presses []int) bool {
	state := make([]bool, len(m.lights_soln))
	for _, button_index := range button_presses {
		button := m.buttons[button_index]
		for _, i := range button {
			state[i] = !state[i]
		}
	}
	return slices.Equal(state, m.lights_soln)
}

type button []int

type machine struct {
	lights_soln []bool
	buttons     []button
	joltages    []int
}

func parse_machines(lines []string) []machine {
	machines := make([]machine, 0, len(lines))
	for _, line := range lines {
		machines = append(machines, parse_machine(line))
	}
	return machines
}

func parse_machine(line string) machine {
	parts := strings.Fields(line)
	// this is a string copy which i feel compelled to point out
	lights := strings.Trim(parts[0], "[]")
	bits := make([]bool, 0, len(lights))
	for _, c := range lights {
		if c == '.' {
			bits = append(bits, false)
		} else if c == '#' {
			bits = append(bits, true)
		} else {
			panic(fmt.Sprintf("Invalid lights input in input line `%s`: `%s`", line, lights))
		}
	}
	joltages_s := strings.Split(strings.Trim(parts[len(parts)-1], "{}"), ",")
	joltages := make([]int, len(joltages_s))
	for _, c := range joltages_s {
		v, err := strconv.Atoi(c)
		utils.Check(err)
		joltages = append(joltages, v)
	}

	buttons := make([]button, 0, len(parts)-2)
	for _, button_s := range parts[1 : len(parts)-1] {
		index_strs := strings.Split(strings.Trim(button_s, "()"), ",")
		button := make([]int, 0, len(index_strs))
		for _, s := range index_strs {
			v, err := strconv.Atoi(s)
			utils.Check(err)
			button = append(button, v)
		}
		buttons = append(buttons, button)
	}
	return machine{
		lights_soln: bits,
		joltages:    joltages,
		buttons:     buttons,
	}

}
