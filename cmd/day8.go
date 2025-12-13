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

var debug bool = false

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		part1_day8()
	},
}

type pair struct {
	v1 *vertex
	v2 *vertex
}

type pairDistance struct {
	p        pair
	distance int
}

type junctionGrid struct {
	vertices []*vertex

	distances []pairDistance
	// this is a list of hashsets, with each set representing a connected set of pairs
	circuits []map[*vertex]bool
}

func part1_day8() {
	fmt.Println("hi day 8")
	grid := parse_input_day8("inputs/8")
	// for i, whatever := range grid.vertices {
	// 	fmt.Printf("  vertex: %s - indexed: %s\n", whatever, grid.vertices[i])
	// }
	// for i, whatever := range grid.vertices {
	// 	fmt.Printf("  vertex: %s - indexed: %s\n", whatever, grid.vertices[i])
	// }
	grid.orderByEuclideanDistance()
	grid.addConnections(1000)
	slices.SortFunc(grid.circuits, func(a, b map[*vertex]bool) int {
		return -cmp.Compare(len(a), len(b))
	})
	pw := 1
	for _, c := range grid.circuits[:3] {
		pw *= len(c)
	}
	fmt.Println("The password is", pw)

}

// gross! and probably unnecessary!
func vertexCmp(v1, v2 *vertex) int {
	if v1.x < v2.x {
		return -1
	} else if v1.x > v2.x {
		return 1
	}
	if v1.y < v2.y {
		return -1
	} else if v1.y > v2.y {
		return 1
	}
	if v1.z < v2.z {
		return -1
	} else if v1.z > v2.z {
		return 1
	}
	return 0
}

func newPair(v1, v2 *vertex) pair {
	if vertexCmp(v1, v2) == 1 {
		// fmt.Printf("Swapping vertices: before (%p, %p) | ", v1, v2)
		v1, v2 = v2, v1
		// fmt.Printf("after (%p, %p)\n", v1, v2)
	}

	return pair{v1: v1, v2: v2}
}

func (p pairDistance) String() string {
	return fmt.Sprintf("%s - %s | %d", p.p.v1, p.p.v2, p.distance)
}

func (g *junctionGrid) printCircuits() {
	fmt.Println("Circuits: ")
	for _, c := range g.circuits {
		if len(c) == 0 {
			continue
		}
		if debug {
			fmt.Print("  {")
			for k, _ := range c {
				fmt.Printf("%p, ", k)
			}
			fmt.Print("}\n")
		}
	}
}

// uhhh let's see if i can do this the dumb way and survive
func (g *junctionGrid) orderByEuclideanDistance() {
	g.calculateDistances()
	slices.SortFunc(g.distances, func(a, b pairDistance) int {
		return cmp.Compare(a.distance, b.distance)
	})

}

// all distance calculations takes ~1s on full input so i think we're okay
func (g *junctionGrid) calculateDistances() {
	for i, v1 := range g.vertices {
		for _, v2 := range g.vertices[i+1:] {
			d := v1.distance(v2)
			if debug {
				// fmt.Printf("Creating NEW pair with vertices %s and %s\n", v1, v2)
			}
			p := newPair(v1, v2)
			g.distances = append(g.distances, pairDistance{p: p, distance: d})
		}
	}
}

func (g *junctionGrid) addConnections(amount int) {
	for _, info := range g.distances[:amount] {
		if debug {
			fmt.Printf("Adding connection: %s - %s\n", info.p.v1, info.p.v2)
		}
		g.addConnection(&info.p)
		if debug {
			g.printCircuits()
		}
	}
}

func (g *junctionGrid) addConnection(p *pair) {
	v1, v2 := p.v1, p.v2
	// how do i do type aliases someone please help me my family is dying
	var found1 *map[*vertex]bool
	var found2 *map[*vertex]bool
	for i, circuit := range g.circuits {
		// N.B. we're not catching the case where both ends of a pair are already in the same circuit
		hasV1 := circuit[v1]
		if debug {
			fmt.Printf("    is %s in %v? %v\n", v1, circuit, hasV1)
		}
		hasV2 := circuit[v2]
		if hasV1 {
			if debug {
				fmt.Printf("Found %s in circuit %d\n", v1, i)
			}
			if found1 != nil {
				panic(fmt.Sprintf("Vertex %s is somehow already in 2 circuits", v1))
			}
			found1 = &circuit
			// one end of our pair is in here, so add the other pair
			// this line modifies circuit to make the 2nd branch always true if we don't cache
			// the presence of the key above.
			circuit[v2] = true
		}
		if hasV2 {
			if debug {
				fmt.Printf("Found %s in circuit %d\n", v2, i)
			}
			if found2 != nil {
				panic(fmt.Sprintf("Vertex %s is somehow already in 2 circuits", v2))
			}
			found2 = &circuit
			circuit[v1] = true
		}
		// it *should* be safe to break early here, but i guess we can keep looping to catch bad data
	}
	if found1 != nil && found2 != nil && found1 != found2 {
		if debug {
			fmt.Println("Found homes for both circuits; merging")
		}
		// merge the circuits; it doesnt matter which one we pick
		// i'm just gonna drain the other circuit & leave it in the list of circuits, which is annoying but should be correct
		// ok why do i have to dereference the pointer to iterate the map?
		for k, v := range *found2 {
			if !v {
				panic(fmt.Sprintf("Found false value for key %s in circuit", k))
			}
			(*found1)[k] = true
		}
		// why the FUCK is clear a magical global namespace function?
		clear(*found2)
	}
	if found1 == nil && found2 == nil {
		if debug {
			fmt.Println("Found no homes; creating new circuit")
		}
		// okay fine we need to create a new circuit
		c := make(map[*vertex]bool)
		c[v1] = true
		c[v2] = true
		g.circuits = append(g.circuits, c)
	}
}

func parse_input_day8(filename string) junctionGrid {
	lines, err := utils.ReadLines(filename)
	utils.Check(err)
	vertices := make([]*vertex, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			panic(fmt.Sprintf("Invalid input line: `%s`", line))
		}
		x, err := strconv.Atoi(parts[0])
		utils.Check(err)
		y, err := strconv.Atoi(parts[1])
		utils.Check(err)
		z, err := strconv.Atoi(parts[2])
		utils.Check(err)
		v := &vertex{x: x, y: y, z: z}
		if debug {
			fmt.Println("Adding vertex", v)
		}
		vertices = append(vertices, v)
	}
	return junctionGrid{vertices: vertices}
}

type vertex struct {
	x int
	y int
	z int
}

func (v *vertex) String() string {
	return fmt.Sprintf("(%d, %d, %d) (%p)", v.x, v.y, v.z, v)
}

// returns square of distance because that sorts the same as distance
func (v *vertex) distance(other *vertex) int {
	// i'm going to guess that go is smart enough to optimize this code
	radicand := (other.x-v.x)*(other.x-v.x) + (other.y-v.y)*(other.y-v.y) + (other.z-v.z)*(other.z-v.z)
	return radicand

}

func init() {
	rootCmd.AddCommand(day8Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day8Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day8Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
