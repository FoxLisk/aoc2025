/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/foxlisk/aoc2025/cmd"
)

type T struct {
	x int
}

func main() {
	xes := []*T{{x: 1}, &T{x: 2}}
	for _, t := range xes {
		fmt.Printf("T(%d) (%p)\n", t.x, t)
	}
	for _, t := range xes {
		fmt.Printf("T(%d) (%p)\n", t.x, t)
	}
	if true {
		cmd.Execute()
	}
}

type I interface {
	f()
	make() I
}

type T1 struct {
	//
}

func (t1 T1) f() {

}

func handle(i I) {

}
