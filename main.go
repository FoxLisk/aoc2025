/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/foxlisk/aoc2025/cmd"
)

func main() {

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