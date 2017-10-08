package backpack

import (
	"math/rand"
	"fmt"
	"strings"
	"time"
)

type Gene struct {
	chromosomes []bool
}

func (gene *Gene) GetLength() int {
	return len(gene.chromosomes)
}

func (gene *Gene) Crossover(other *Gene, splitPoint int) (first, second Gene) {
	first = Gene{}
	second = Gene{}

	first.chromosomes = append(first.chromosomes, gene.chromosomes[:splitPoint]...)
	first.chromosomes = append(first.chromosomes, other.chromosomes[splitPoint:]...)
	second.chromosomes = append(second.chromosomes, other.chromosomes[:splitPoint]...)
	second.chromosomes = append(second.chromosomes, gene.chromosomes[splitPoint:]...)

	return
}

func (gene *Gene) Mutation() *Gene {
	if rand.Float64() > 0.5 {
		mutationIndex := rand.Intn(len(gene.chromosomes))
		gene.chromosomes[mutationIndex] = !gene.chromosomes[mutationIndex]
	}
	return gene
}

func GetInitialPopulation(populationSize, chromosomesSize int) []Gene {
	initialPopulation := make([]Gene, populationSize)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < populationSize; i++ {
		chromosomes := make([]bool, chromosomesSize)
		for j := 0; j < chromosomesSize; j++ {
			if rand.Float64() >= 0.5 {
				chromosomes[j] = true
			} else {
				chromosomes[j] = false
			}
		}
		initialPopulation[i] = Gene{chromosomes: chromosomes}
	}
	return initialPopulation
}

func (gene *Gene) GetHealth(items []Item) float64 {
	sum := 0.0
	for i, contains := range gene.chromosomes {
		if contains {
			sum += items[i].GetPrice()
		}
	}
	return sum
}

func (gene *Gene) GetDecease(items []Item) float64 {
	sum := 0.0
	for i, contains := range gene.chromosomes {
		if contains {
			sum += items[i].GetWeight()
		}
	}
	return sum
}

func (gene Gene) String() string {
	stringValues := make([]string, len(gene.chromosomes))
	for i, value := range gene.chromosomes {
		if value {
			stringValues[i] = "1"
		} else {
			stringValues[i] = "0"
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(stringValues, ""))
}

func GetMinimalGene(items []Item, maxWeight float64) *Gene {
	gene := &Gene{}
	gene.chromosomes = make([]bool, len(items))
	for i, v := range items {
		if v.GetWeight() < maxWeight {
			gene.chromosomes[i] = true
			return gene
		}
	}
	panic("All items are heavy")
}
