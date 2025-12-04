/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day 1",
	Long:  `blah`,
	Run:   run,
}

func init() {
	rootCmd.AddCommand(day1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



func run(cmd *cobra.Command, args []string) {
	part1()
	part2()
}

func part1() {
	fmt.Println("Starting part 1")
	lines, err := utils.ReadLines("inputs/1_1")
	utils.Check(err)
	password := 0
	value := 50
	for _, line := range lines {
		if line == "" {
			fmt.Println("empty line? continuing...")
			continue
		}
		rotation := parseLine(line)
		prev := value
		value = (value + rotation) % 100
		if value < 0 {
			value += 100
		}
		if value == 0 {
			password++
		}
		fmt.Printf("%d rotated %s (%d) -> %d\n", prev, line, rotation, value)
	}
	fmt.Println("The password is ", password)
}

func part2() {
	fmt.Println("Starting part 2")
	lines, err := utils.ReadLines("inputs/1_2")
	utils.Check(err)
	password := 0
	value := 50
	for _, line := range lines {
		if line == "" {
			fmt.Println("empty line? continuing...")
			continue
		}
		rotation := parseLine(line)
		full_rotations := utils.Abs(rotation) / 100
		password += full_rotations

		prev := value
		value = (prev + rotation) % 100
		// normalize to range
		if value < 0 {
			value += 100
		}
		passed_0 := false
		// if we passed 0 going left, our new value will be higher than our old value; etc
		if prev != 0 && value != 0 && ((rotation < 0 && value > prev) || (rotation > 0 && value < prev)) {
			password += 1
			passed_0 = true
		}
		if value == 0 {
			password += 1
		}
		fmt.Printf("Starting %d: rotated %d to %d: full_rotations %d: passed zero? %t\n    password: %d\n", prev, rotation, value, full_rotations, passed_0, password)
	}
	fmt.Println("The password is", password)
}

// returns rotation amount as a signed integer
func parseLine(line string) int {

	dir := line[0:1]
	rest := line[1:]
	val, err := strconv.Atoi(rest)
	utils.Check(err)

	switch dir {
	case "L":
		val *= -1
	case "R":
		//pass
	default:
		panic("Invalid rotation direction.")
	}
	return val
}
