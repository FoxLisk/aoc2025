/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day f4r",
	Long:  `asdf.`,
	Run:   run_day4,
}

func init() {
	rootCmd.AddCommand(day4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type indices struct {
	i int
	j int
}

func newIndices(i, j int) indices {
	return indices{
		i: i,
		j: j,
	}
}

func run_day4(cmd *cobra.Command, args []string) {
	part1_day4()
	part2_day4()
}

func part1_day4() {
	grid := make_grid()
	thePassword := 0
	rows := len(grid)
	// ignoring empty inputs
	cols := len(grid[0])
	for i := range rows {
		for j := range cols {
			if grid[i][j] != "@"[0] {
				continue
			}
			// fmt.Printf("Looking at (%d, %d)\n", i, j)
			neighbours_to_check := adjacent_indices(newIndices(i, j), rows, cols)
			var neighbours_strs []string
			for _, n := range neighbours_to_check {
				neighbours_strs = append(neighbours_strs, fmt.Sprintf("(%d, %d)", n.i, n.j))
			}
			// fmt.Printf("   neighbours: %s\n", strings.Join(neighbours_strs, " "))
			total_ats := 0
			for _, neighbour := range neighbours_to_check {
				if grid[neighbour.i][neighbour.j] == "@"[0] {
					total_ats += 1
				}
			}
			if total_ats < 4 {
				fmt.Printf("Position (%d, %d) is good to go\n", i, j)
				thePassword += 1
			}
		}
	}
	fmt.Println("THe password is", thePassword)
}

func make_grid() [][]byte {
	grid_raw, err := utils.ReadLines("inputs/4")
	utils.Check(err)

	grid := make([][]byte, 0, len(grid_raw))
	for _, row_raw := range grid_raw {
		row := []byte(row_raw)
		grid = append(grid, row)
	}
	return grid
}

func part2_day4() {
	grid := make_grid()
	total_removable := 0
	for {
		removable := removable_indices(grid)
		if len(removable) == 0 {
			break
		}
		total_removable += len(removable)
		for _, loc := range removable {
			grid[loc.i][loc.j] = "."[0]
		}
	}
	fmt.Println("Total removable cells:", total_removable)
}

// im just gonna be pretty inefficient and not try to do anything as i go, rather just doing a bunch of extra loops
// if AoC wants me to be efficient they are free to give me difficult problems
func removable_indices(grid [][]byte) []indices {
	var removable []indices
	rows := len(grid)
	// ignoring empty inputs
	cols := len(grid[0])
	for i := range rows {
		for j := range cols {
			if grid[i][j] != "@"[0] {
				continue
			}
			// fmt.Printf("Looking at (%d, %d)\n", i, j)
			neighbours_to_check := adjacent_indices(newIndices(i, j), rows, cols)
			// var neighbours_strs []string
			// for _, n := range neighbours_to_check {
			// 	neighbours_strs = append(neighbours_strs, fmt.Sprintf("(%d, %d)", n.i, n.j))
			// }
			// fmt.Printf("   neighbours: %s\n", strings.Join(neighbours_strs, " "))
			total_ats := 0
			for _, neighbour := range neighbours_to_check {
				if grid[neighbour.i][neighbour.j] == "@"[0] {
					total_ats += 1
				}
			}
			if total_ats < 4 {
				removable = append(removable, newIndices(i, j))
			}
		}
	}
	return removable
}

func adjacent_indices(loc indices, rows int, cols int) []indices {
	neighbours := make([]indices, 0, 8)
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			i_candidate := loc.i + i
			j_candidate := loc.j + j
			if i_candidate < 0 || j_candidate < 0 || i_candidate >= rows || j_candidate >= cols {
				continue
			}
			neighbours = append(neighbours, indices{i: i_candidate, j: j_candidate})
		}
	}
	return neighbours
}
