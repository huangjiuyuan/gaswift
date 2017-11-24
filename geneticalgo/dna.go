package geneticalgo

import (
	"log"
	"math/rand"
)

type DNA struct {
	Genes []byte
}

func InitDNA(num int) *DNA {
	dna := new(DNA)
	dna.Genes = make([]byte, num)
	for i := 0; i < len(dna.Genes); i++ {
		dna.Genes[i] = byte(rand.Intn(96) + 32)
	}
	return dna
}

func (dna *DNA) Fitness(target string) float32 {
	score := 0
	for i := 0; i < len(dna.Genes); i++ {
		if dna.Genes[i] == target[i] {
			score++
		}
	}
	return float32(score) / float32(len(target))
}

func (dna *DNA) Crossover(partner *DNA) *DNA {	
	child := InitDNA(len(dna.Genes))
	mid := rand.Intn(len(dna.Genes))
	for i := 0; i < len(dna.Genes); i++ {
		if i > mid {
			child.Genes[i] = dna.Genes[i]
		} else {
			child.Genes[i] = partner.Genes[i]
		}
	}
	return child
}

func (dna *DNA) Mutate(rate float32) {
	if rate < 0 || rate > 1 {
		log.Fatalf("Mutation rate %f must be between 0 and 1", rate)
	}
	for i := 0; i < len(dna.Genes); i++ {
		if rand.Float32() < rate {
			dna.Genes[i] = byte(rand.Intn(96) + 32)
		}
	}
}
