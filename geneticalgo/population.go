package geneticalgo

import (
	"math/rand"
)

type Population struct {
	MutationRate float32
	Population []DNA
	MatingPool []DNA
	Target string
	Generations int
	Finished bool
	PerfectScore float32
}

func InitPopulation(target string, rate float32, num int) *Population {
	p := new(Population)
	p.MutationRate = rate
	p.Population = make([]DNA, num)
	for i := 0; i < len(p.Population); i++ {
		p.Population[i] = InitDNA(len(target))
		p.Population[i].CalcFitness(target)
	}
	p.MatingPool = make([]DNA, 0)
	p.Finished = false
	p.PerfectScore = 1
	return p
}

func (p *Population) NaturalSelection() {
	p.MatingPool = make([]DNA, 0)
	var maxFitness float32 = 0

	for i := 0; i < len(p.Population); i++ {
		if p.Population[i].Fitness > maxFitness {
			maxFitness = p.Population[i].Fitness
		}
	}

	for i := 0; i < len(p.Population); i++ {
		n := int(p.Population[i].Fitness * 100)
		for j := 0; j < n; j++ {
			p.MatingPool = append(p.MatingPool, p.Population[i])
		}
	}
}

func (p *Population) Generate() {
	for i := 0; i < len(p.Population); i++ {
		a := rand.Intn(len(p.MatingPool))
		b := rand.Intn(len(p.MatingPool))
		partnerA := p.MatingPool[a]
		partnerB := p.MatingPool[b]
		child := partnerA.Crossover(partnerB)
		child.Mutate(p.MutationRate)
		p.Population[i] = child
	}
	p.Generations++
}

func (p *Population) Best() string {
	var best float32
	index := 0

	for i := 0; i < len(p.Population); i++ {
		if p.Population[i].Fitness > best {
			index = i
			best = p.Population[i].Fitness
		}
	}

	if best == p.PerfectScore {
		p.Finished = true
	}
	return string(p.Population[index].Genes)
}

func (p *Population) IsFinished() bool {
	return p.Finished
}

func (p *Population) GetGenerations() int {
	return p.Generations
}

func (p *Population) AverageFitness() float32 {
	var total float32
	for i := 0; i < len(p.Population); i++ {
		total += p.Population[i].Fitness
	}
	return total / float32(len(p.Population))
}
