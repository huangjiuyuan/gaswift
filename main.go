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
	fmt.Println(dna.Fitness(target))
	child :=dna.Crossover(partner)
	fmt.Printf("%s\n", child)
	child.Mutate(0.2)
	fmt.Printf("%s\n", child)
}
