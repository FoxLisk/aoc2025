/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "A brief description of your command",
	Long:  `x`,
	Run: func(cmd *cobra.Command, args []string) {
		// part1_day7()
		part2_day7()
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}

func part1_day7() {
	lines, err := utils.ReadLines("inputs/7_ex")
	utils.Check(err)
	var entryIndex = -1
	for i, c := range lines[0] {
		if c == 'S' {
			entryIndex = i
		}
	}
	if entryIndex == -1 {
		panic("no entrypoint found")
	}
	state := beamState{beams: map[int]bool{entryIndex: true}, width: len(lines[0])}
	for _, line := range lines {
		row := parseManifoldRow(line)
		state.moveThroughRow(row)
		state.displayAfterRow(row)
	}
	fmt.Println("Total splits:", state.totalSplits)
}

func part2_day7() {
	// copypaste, fucking blow me
	lines, err := utils.ReadLines("inputs/7")
	utils.Check(err)
	var entryIndex = -1
	for i, c := range lines[0] {
		if c == 'S' {
			entryIndex = i
		}
	}
	if entryIndex == -1 {
		panic("no entrypoint found")
	}
	state := beamState2{particleFuturesAtIndex: map[int]int{entryIndex: 1}, width: len(lines[0]), totalFutures: 1}
	for _, line := range lines {
		row := parseManifoldRow(line)
		state.moveThroughRow(row)
		state.displayAfterRow(row)
	}
	fmt.Println("Total futures:", state.totalFutures)
}

type beamState struct {
	width       int
	beams       map[int]bool
	totalSplits int
}

type beamState2 struct {
	width                  int
	particleFuturesAtIndex map[int]int
	totalFutures           int
}

func (bs *beamState) moveThroughRow(row manifoldRow) {
	for _, i := range row.splitters {
		if bs.beams[i] {
			bs.beams[i-1] = true
			bs.beams[i+1] = true
			bs.totalSplits = bs.totalSplits + 1
		}
		bs.beams[i] = false
	}
}

func (bs beamState) displayAfterRow(mr manifoldRow) {
	s := ""
	for i := range bs.width {
		if bs.beams[i] && slices.Contains(mr.splitters, i) {
			fmt.Println("ELKSJDFLKJSDF")
		}
		if bs.beams[i] {
			s = s + "|"
		} else if slices.Contains(mr.splitters, i) {
			s = s + "^"
		} else {
			s = s + "."
		}
	}
	fmt.Println(s)
}

func (bs *beamState2) moveThroughRow(row manifoldRow) {
	// at each splitter, each current timeline is split (i.e. incremented)
	// for each particle timeline that made it to this splitter. so if 1 particle
	// hits, timelines goes up by 1. if 2 particles hit, they each split,
	// and the total futures goes up by 2
	for _, i := range row.splitters {
		timelines := bs.particleFuturesAtIndex[i]
		bs.totalFutures = bs.totalFutures + timelines
		// each particle that hits this contributes 1 to each side
		bs.particleFuturesAtIndex[i-1] += timelines
		bs.particleFuturesAtIndex[i+1] += timelines
		bs.particleFuturesAtIndex[i] = 0 // but the ones that got here are cleared out
	}
}

func (bs beamState2) displayAfterRow(mr manifoldRow) {
	s := ""
	for i := range bs.width {
		if bs.particleFuturesAtIndex[i] > 0 && slices.Contains(mr.splitters, i) {
			fmt.Println("ELKSJDFLKJSDF")
		}
		if bs.particleFuturesAtIndex[i] > 0 {
			s = s + "|"
		} else if slices.Contains(mr.splitters, i) {
			s = s + "^"
		} else {
			s = s + "."
		}
	}
	fmt.Println(s)
}

type manifoldRow struct {
	splitters []int
}

func parseManifoldRow(row string) manifoldRow {
	mr := manifoldRow{}
	for i, c := range row {
		if c == '^' {
			mr.splitters = append(mr.splitters, i)
		}
	}
	return mr
}
