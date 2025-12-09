/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day fve",
	Long:  `asdf`,
	Run:   run_day5,
}

func init() {
	rootCmd.AddCommand(day5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run_day5(cmd *cobra.Command, args []string) {

	ranges, ids := parse_input("inputs/5")
	part1_day5(ranges, ids)
	part2_day5(ranges)
}

func part1_day5(ranges []*utils.IdRange, ids []int) {
	valid := 0
	for _, id := range ids {
		if is_fresh(id, ranges) {
			valid += 1
		}
	}
	fmt.Println("Fresh ingredients:", valid)
}

func printRanges(ranges []*utils.IdRange) {
	fmt.Println("Ranges:")
	for _, r := range ranges {
		fmt.Printf("  IdRange{ Start: %d, End: %d }\n", r.Start, r.End)
	}
}

func part2_day5(ranges []*utils.IdRange) {
	slices.SortFunc(ranges, func(a, b *utils.IdRange) int {
		return cmp.Compare(a.Start, b.Start)
	})
	printRanges(ranges)
	count := len(ranges)

	// this is probably a pretty cringe way of making sure we are thorough but idc really
	for {
		ranges = mergeRanges(ranges)
		if len(ranges) == count {
			break
		}
		count = len(ranges)
		slices.SortFunc(ranges, func(a, b *utils.IdRange) int {
			return cmp.Compare(a.Start, b.Start)
		})
		printRanges(ranges)
	}

	totalValid := 0
	for _, r := range ranges {
		totalValid += 1 + int(r.End) - int(r.Start)
	}
	fmt.Println("Total valid:", totalValid)
}

func mergeRanges(ranges []*utils.IdRange) []*utils.IdRange {
	var mergedRanges []*utils.IdRange
	mergedRanges = append(mergedRanges, ranges[0])
	for i, r := range ranges {
		if i == 0 {
			continue // lol
		}
		subsumed := false

		for _, merged := range mergedRanges {
			// if r.End == merged.Start - 1, we want to merge them; that's like
			// (1, 3) (4, 5) which should be merged into (1, 5)
			if r.Start <= merged.Start && r.End >= merged.Start-1 {
				merged.Start = r.Start
				subsumed = true
			}
			// similarly: consider the ranges in the other order, (4, 5) and (1, 3)
			if r.End >= merged.End && r.Start <= merged.End+1 {
				merged.End = r.End
				subsumed = true
			}
			if r.Start >= merged.Start && r.End <= merged.End {
				// the case where the range is totally contained inside the other
				subsumed = true
			}
			if subsumed {
				break
			}
		}
		if !subsumed {
			mergedRanges = append(mergedRanges, r)
		}
	}
	return mergedRanges
}

func is_fresh(id int, ranges []*utils.IdRange) bool {
	for _, r := range ranges {
		if id >= int(r.Start) && id <= int(r.End) {
			return true
		}
	}
	return false
}

func parse_input(filename string) ([]*utils.IdRange, []int) {
	lines, err := utils.ReadLines(filename)
	utils.Check(err)
	state := "ranges"
	var ranges []*utils.IdRange
	var ids []int
	for _, line := range lines {
		if state == "ranges" {
			if line == "" {
				state = "ids"
				continue
			}
			thisRange, err := utils.IdRangeFromString(line)
			utils.Check(err)
			ranges = append(ranges, thisRange)
		} else if state == "ids" {
			i, err := strconv.Atoi(line)
			utils.Check(err)
			ids = append(ids, i)
		}
	}
	return ranges, ids
}
