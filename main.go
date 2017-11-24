package main

import (
	"fmt"

	ga "github.com/huangjiuyuan/gaswift/geneticalgo"
)

func main() {
	target := "Taylor Swift"
	dna := ga.InitDNA(len(target))
	partner := ga.InitDNA(len(target))
	fmt.Printf("%s, %s\n", dna.Genes, partner.Genes)
	fmt.Println(dna.CalcFitness(target))
	child := dna.Crossover(partner)
	fmt.Printf("%s\n", child.Genes)
	child.Mutate(0.2)
	fmt.Printf("%s\n", child.Genes)

	p := ga.InitPopulation(target, 0.01, 2000)
	for {
		p.NaturalSelection()
		p.Generate()
		p.CalcFitness(target)
		fmt.Println(p.Best())
		fmt.Println(p.AverageFitness())
		fmt.Println(p.GetGenerations())
		if p.IsFinished() {
			break
		}
	}
}
