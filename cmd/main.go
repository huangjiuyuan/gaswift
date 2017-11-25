package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/huangjiuyuan/gaswift"
)

func main() {
	var target string
	if len(os.Args) > 1 {
		target = strings.Join(os.Args[1:], " ")
	} else {
		target = "I knew you were trouble when you walked in."
	}
	fmt.Println(target)

	p := gaswift.InitPopulation(target, 0.01, 1000)
	for {
		p.NaturalSelection()
		p.Generate()
		p.CalcPopulationFitness(target)
		fmt.Println(p.Best())
		fmt.Println(p.AverageFitness())
		fmt.Println(p.GetGenerations())
		if p.IsFinished() {
			break
		}
	}
}
