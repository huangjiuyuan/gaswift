package gaswift

import (
	"log"
	"math/rand"
)

type DNA struct {
	Genes []byte
	Fitness float32
}

func InitDNA(num int) DNA {
	d := new(DNA)
	d.Genes = make([]byte, num)
	for i := 0; i < len(d.Genes); i++ {
		d.Genes[i] = byte(rand.Intn(96) + 32)
	}
	return *d
}

func (d *DNA) CalcFitness(target string) float32 {
	score := 0
	for i := 0; i < len(d.Genes); i++ {
		if d.Genes[i] == target[i] {
			score++
		}
	}
	d.Fitness = float32(score) / float32(len(target))
	return d.Fitness
}

func (d *DNA) Crossover(partner DNA) DNA {	
	child := InitDNA(len(d.Genes))
	mid := rand.Intn(len(d.Genes))
	for i := 0; i < len(d.Genes); i++ {
		if i > mid {
			child.Genes[i] = d.Genes[i]
		} else {
			child.Genes[i] = partner.Genes[i]
		}
	}
	return child
}

func (d *DNA) Mutate(rate float32) {
	if rate < 0 || rate > 1 {
		log.Fatalf("Mutation rate %f must be between 0 and 1", rate)
	}
	for i := 0; i < len(d.Genes); i++ {
		if rand.Float32() < rate {
			d.Genes[i] = byte(rand.Intn(96) + 32)
		}
	}
}
