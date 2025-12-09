/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day 2",
	Long:  `asdfkjasldkjf.`,
	Run:   run_2,
}

func init() {
	rootCmd.AddCommand(day2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run_2(cmd *cobra.Command, args []string) {
	part1_day2()
	part2_day2()
}

func read_ranges(filename string) []*utils.IdRange {
	var ranges []*utils.IdRange
	data, err := os.ReadFile(filename)
	utils.Check(err)
	data_s := string(data)
	data_s = strings.Trim(data_s, "\n")
	splitted := strings.Split(data_s, ",")
	for _, s := range splitted {
		myRange, err := utils.IdRangeFromString(s)
		utils.Check(err)
		ranges = append(ranges, myRange)
	}
	return ranges
}

func total_things_to_check(ranges []utils.IdRange) int {
	total := 0
	for _, thing := range ranges {
		total += int(thing.End) - int(thing.Start) + 1
	}
	return total
}

func id_is_invalid_1(num uint64) bool {
	s := strconv.FormatUint(num, 10)
	if len(s)%2 == 1 {
		return false
	}

	if s[:len(s)/2] == s[len(s)/2:] {
		return true
	}

	return false
}

func id_is_invalid_2(num uint64) bool {
	s := []rune(strconv.FormatUint(num, 10))
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		// we have a divisor; chop the string up into i-ths and see if they are all identical
		pieces := slices.Chunk(s, i)
		var first []rune
		allSame := true
		for piece := range pieces {
			if first == nil {
				first = piece
			}
			if !slices.Equal(first, piece) {
				allSame = false
				break
			}
		}
		if allSame {
			return true
		}
	}

	return false
}

func invalid_ids_in_range(a_range *utils.IdRange, checker func(uint64) bool) []uint64 {
	var invalidIds []uint64
	candidate := a_range.Start
	for candidate <= a_range.End {
		if checker(candidate) {
			invalidIds = append(invalidIds, candidate)
		}
		candidate += 1
	}
	return invalidIds
}

func part1_day2() {
	ranges := read_ranges("inputs/2_1_ex")
	var allInvalidIds []uint64
	for _, a_range := range ranges {
		invalidOnes := invalid_ids_in_range(a_range, id_is_invalid_1)
		allInvalidIds = append(allInvalidIds, invalidOnes...)
	}
	fmt.Println("All invalid IDs: ", allInvalidIds)
	the_password := 0
	for _, invalidId := range allInvalidIds {
		the_password += int(invalidId)
	}
	fmt.Println("The password is: ", the_password)

	// all_ranges := read_ranges("inputs/2_1")

}

func part2_day2() {
	ranges := read_ranges("inputs/2_1")
	var allInvalidIds []uint64
	for _, a_range := range ranges {
		invalidOnes := invalid_ids_in_range(a_range, id_is_invalid_2)
		allInvalidIds = append(allInvalidIds, invalidOnes...)
	}
	fmt.Println("All invalid IDs: ", allInvalidIds)
	the_password := 0
	for _, invalidId := range allInvalidIds {
		the_password += int(invalidId)
	}
	fmt.Println("The password is: ", the_password)

}
