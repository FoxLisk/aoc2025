/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/foxlisk/aoc2025/internal/utils"
	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Day[9] daily")
		// part1_day9()
		part2_day9()
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}

type point struct {
	x int
	y int
}

func (p point) intersects(l line) bool {
	if l.isHorizontal() {
		panic("I'm pretty sure testing for horizontal intersection is a bug")
	} else {
		// is vertical; we have no diagonals
		y1, y2 := l.p1.y, l.p2.y
		// fmt.Printf(" y1, y2 := %v, %v", y1, y2)
		if y1 > y2 {
			y1, y2 = y2, y1
			// fmt.Printf(" - swapped - ")
		}
		// fmt.Printf("\n")

		return p.x == l.p1.x && y1 <= p.y && p.y <= y2
	}
}

type line struct {
	p1 point
	p2 point
}

func (l line) isHorizontal() bool {
	return l.p1.y == l.p2.y
}

// creates a new line, such that the first point is to the left or bottom of the second point
func newLine(p1, p2 point) line {
	if p1.x != p2.x && p1.y != p2.y {
		panic("Lines must be horizontal or vertical")
	}
	if p1.x == p2.x && p1.y == p2.y {
		panic("Please don't include degenerate lines (aka points)")
	}
	if p1.x > p2.x {
		p1, p2 = p2, p1
	}
	if p1.x == p2.x && p1.y > p2.y {
		p1, p2 = p2, p1
	}

	return line{
		p1: p1, p2: p2,
	}
}

type corners struct {
	p1   point
	p2   point
	area int
}

type grid struct {
	points          []point
	horizontalLines []line
	verticalLines   []line
	// for display purposes
	width        int
	height       int
	finalized    bool
	pointsInside map[point]bool
}

func newGrid(cap int) grid {
	return grid{
		points:          make([]point, 0, cap),
		horizontalLines: make([]line, 0),
		verticalLines:   make([]line, 0),
		pointsInside:    make(map[point]bool),
	}
}

func (g *grid) addPoint(p point) {
	if g.finalized {
		panic("Cannot add points to finalized grids")
	}
	if len(g.points) > 0 {
		line := newLine(g.points[len(g.points)-1], p)
		g.addLine(line)

	}
	g.points = append(g.points, p)
	if p.x > g.width {
		g.width = p.x + 1
	}
	if p.y > g.height {
		g.height = p.y + 1
	}
}

func (g *grid) addLine(l line) {
	if g.finalized {
		panic("Cannot add lines to finalized grids")
	}
	if l.isHorizontal() {
		g.horizontalLines = append(g.horizontalLines, l)
	} else {
		g.verticalLines = append(g.verticalLines, l)
	}
}

func (g *grid) doneAddingPoints() {
	l := newLine(g.points[0], g.points[len(g.points)-1])
	g.addLine(l)
	// pad display
	g.width += 1
	g.height += 1
	slices.SortFunc(g.verticalLines, func(a, b line) int {
		return cmp.Compare(a.p1.x, b.p1.x)
	})
	g.finalized = true
}

func (g grid) pointIsInside(p point) bool {
	if v, in := g.pointsInside[p]; in {
		return v
	}
	if len(g.points) != len(g.horizontalLines)+len(g.verticalLines) {
		panic("Please finalize grid before asking about interior")
	}
	if debug {
		fmt.Println("Checking if point", p, "is inside")
	}

	valid := false
	var previousLine *line
	for _, l := range g.verticalLines {
		if debug {
			fmt.Println(" Thinking about vertical line ", l)
		}
		if l.p1.x > p.x {
			// g.verticalLines is sorted, so this means we're into irrelevant lines
			break
		}
		testPoint := point{x: l.p1.x, y: p.y}
		if debug {
			fmt.Print("  Checking if ", testPoint, " intersects", l)
		}
		if testPoint.intersects(l) {
			if debug {
				fmt.Print(" - it does!")
			}

			if !linesExtendVertically(previousLine, &l, g.horizontalLines) {
				valid = !valid
			} else {
				if debug {
					fmt.Print(" ...but the lines extend vertically")
				}

			}
			previousLine = &l

		}
		if debug {
			fmt.Printf(" valid now %v \n", valid)
		}
	}
	// clumsy way of handling the farthest-right line on this row
	if previousLine != nil && p.intersects(*previousLine) {
		valid = true
	}
	g.pointsInside[p] = valid
	return valid
}

func linesExtendVertically(l1, l2 *line, horizontalLines []line) bool {
	if l1 == nil {
		return false
	}
	if l1.isHorizontal() && l2.isHorizontal() {
		panic("Only vertical lines can extend vertically")
	}
	if l1.p1.x > l2.p1.x {
		panic("Please provide lines in left-to-right order")
	}
	yToCheck := -1
	if l1.p1.y == l2.p1.y {
		return false // this is a U shape
	} else if l1.p1.y == l2.p2.y {
		yToCheck = l1.p1.y // this is a |_  shape
		//             |
	} else if l1.p2.y == l2.p1.y {
		yToCheck = l1.p2.y // this is a  _| shape
		//           |
	} else if l1.p2.y == l2.p2.y {
		return false // this is an n shape
	}
	for _, l := range horizontalLines {
		if !l.isHorizontal() {
			panic("Please only pass horizontal lines!!!11")

		}
		if l.p1.y == yToCheck {
			return true
		}
	}
	return false
}

func (g *grid) printGrid() {
	textGrid := make([][]byte, 0, g.height)
	for range g.height {
		row := slices.Repeat([]byte{'.'}, g.width)
		textGrid = append(textGrid, row)
	}
	for _, p := range g.points {
		textGrid[p.y][p.x] = '#'
	}
	for y := range textGrid {
		for x := range textGrid[0] {
			if textGrid[y][x] != '.' {
				continue
			}
			p := point{x: x, y: y}
			if g.pointIsInside(p) {
				textGrid[y][x] = 'X'
			}

		}

	}
	for _, s := range textGrid {
		fmt.Println(string(s))
	}
}

func rectanglePerimeter(p1, p2 point) int {
	return 2 * (iAbs(p1.x-p2.x) + iAbs(p1.y-p2.y))
}

func (g *grid) rectangleIsInside(p1, p2 point) bool {
	minX := min(p1.x, p2.x)
	maxX := max(p1.x, p2.x)
	minY := min(p1.y, p2.y)
	maxY := max(p1.y, p2.y)
	for x := minX; x <= maxX; x++ {
		for _, p := range []point{{x: x, y: minY}, {x: x, y: maxY}} {
			if !g.pointIsInside(p) {
				return false
			}
		}
	}
	for y := minY; y <= maxY; y++ {
		for _, p := range []point{{x: maxX, y: y}, {x: minX, y: y}} {
			if !g.pointIsInside(p) {
				return false
			}
		}
	}
	return true
}

func part1_day9() {
	lines, err := utils.ReadLines("inputs/9")
	utils.Check(err)
	points := make([]point, 0, len(lines))
	for _, l := range lines {
		parts := strings.Split(l, ",")
		x, err := strconv.Atoi(parts[0])
		utils.Check(err)
		y, err := strconv.Atoi(parts[1])
		utils.Check(err)
		points = append(points, point{x: x, y: y})
	}
	fmt.Println(points)
	maxArea := 0
	var bestPoint *corners
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			area := rectangleSize(p1, p2)
			if area > maxArea {
				maxArea = area
				bestPoint = &corners{p1: p1, p2: p2, area: area}
			}
		}
	}
	fmt.Println(bestPoint)
}

func part2_day9() {
	lines, err := utils.ReadLines("inputs/9")
	utils.Check(err)
	g := newGrid(len(lines))
	for _, l := range lines {
		parts := strings.Split(l, ",")
		x, err := strconv.Atoi(parts[0])
		utils.Check(err)
		y, err := strconv.Atoi(parts[1])
		utils.Check(err)
		g.addPoint(point{x: x, y: y})
	}
	g.doneAddingPoints()
	if debug {
		g.printGrid()
	}
	type candidateRect struct {
		p1   point
		p2   point
		area int
	}
	var candidates []candidateRect
	for i, p1 := range g.points {
		for _, p2 := range g.points[i:] {
			area := rectangleSize(p1, p2)
			candidates = append(candidates, candidateRect{p1: p1, p2: p2, area: area})

		}
	}
	slices.SortFunc(candidates, func(a, b candidateRect) int { return -cmp.Compare(a.area, b.area) })
	for _, candidate := range candidates {
		fmt.Println("Testing candidate of perimeter", rectanglePerimeter(candidate.p1, candidate.p2))
		if g.rectangleIsInside(candidate.p1, candidate.p2) {
			fmt.Println("The password is", candidate.area)
			break
		}
	}
}

// its DUMB
func rectangleSize(p1, p2 point) int {
	xd := iAbs(p2.x-p1.x) + 1
	yd := iAbs(p2.y-p1.y) + 1
	return xd * yd
}

// fuuuuuck this language
func iAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
