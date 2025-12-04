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

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day three",
	Long:  `no thank you.`,
	Run:   run_day3,
}

func init() {
	rootCmd.AddCommand(day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run_day3(cmd *cobra.Command, args []string) {
	part1_day3()
	part2_day3()
}

func parse_into_joltages(filename string) [][]uint64 {
	lines, err := utils.ReadLines(filename)
	utils.Check(err)
	allJoltages := make([][]uint64, 0, len(lines))
	for _, line := range lines {
		joltages := make([]uint64, 0, len(line))
		for _, c := range line {
			num, err := strconv.ParseUint(string(c), 10, 64)
			utils.Check(err)
			joltages = append(joltages, num)
		}
		allJoltages = append(allJoltages, joltages)
	}
	return allJoltages
}

func find_maximum_joltage_for_bank_part_1(joltages []uint64) uint64 {
	var first_digit uint64 = 0
	first_index := len(joltages)
	for i := len(joltages) - 2; i >= 0; i-- {
		if joltages[i] >= first_digit {
			first_digit = joltages[i]
			first_index = i
		}
	}
	var second_digit uint64 = 0
	for i := len(joltages) - 1; i > first_index; i-- {
		second_digit = max(second_digit, joltages[i])
	}
	return first_digit*10 + second_digit
}

func find_maximum_joltage_for_bank_part_2(joltages []uint64) uint64 {
	digits := make([]uint64, 0, 12)
	for len(digits) < 12 {
		digits_remaining := 12 - len(digits)
		var best_digit uint64 = 0
		best_index := len(joltages)
		for i := len(joltages) - digits_remaining; i >= 0; i-- {
			if joltages[i] >= best_digit {
				best_digit = joltages[i]
				best_index = i
			}
		}
		digits = append(digits, best_digit)
		joltages = joltages[best_index+1:]
	}

	digitString := ""
	for _, d := range digits {
		digitString += strconv.FormatUint(d, 10)
	}
	num, err := strconv.ParseUint(digitString, 10, 64)
	utils.Check(err)
	return num
}

func part1_day3() {
	joltageBanks := parse_into_joltages("inputs/3_1_ex")
	thePassword := 0
	for _, bank := range joltageBanks {
		fmt.Println(bank)
		maxJoltage := find_maximum_joltage_for_bank_part_1(bank)
		fmt.Println("Best: ", maxJoltage)
		thePassword += int(maxJoltage)
	}
	fmt.Println("The password is:", thePassword)
}

func part2_day3() {
	joltageBanks := parse_into_joltages("inputs/3_1")
	thePassword := 0
	for _, bank := range joltageBanks {
		fmt.Println(bank)
		maxJoltage := find_maximum_joltage_for_bank_part_2(bank)
		fmt.Println("Best: ", maxJoltage)
		thePassword += int(maxJoltage)
	}
	fmt.Println("The password is:", thePassword)
}
