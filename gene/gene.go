package gene

import (
	"math/rand"
	"fmt"
	"strings"
)

type Gene struct {
	chromosomes []bool
}

type ContextValue interface {
	GetValue() float64
}

func (gene *Gene) GetLength() int {
	return len(gene.chromosomes)
}

func (gene *Gene) Crossover(other *Gene, splitPoint int) (first, second *Gene) {
	first = &Gene{}
	second = &Gene{}

	first.chromosomes = append(first.chromosomes, gene.chromosomes[:splitPoint]...)
	first.chromosomes = append(first.chromosomes, other.chromosomes[splitPoint:]...)
	second.chromosomes = append(second.chromosomes, other.chromosomes[:splitPoint]...)
	second.chromosomes = append(second.chromosomes, gene.chromosomes[splitPoint:]...)

	return
}

func (gene *Gene) Mutation() *Gene {
	mutationIndex := rand.Intn(len(gene.chromosomes))
	gene.chromosomes[mutationIndex] = !gene.chromosomes[mutationIndex]
	return gene
}

func GetInitialPopulation(populationSize, chromosomesSize int) []*Gene {
	initialPopulation := make([]*Gene, populationSize)
	for i := 0; i < populationSize; i++ {
		chromosomes := make([]bool, chromosomesSize)
		for j := 0; j < chromosomesSize; j++ {
			if rand.Float64() >= 0.5 {
				chromosomes[j] = true
			} else {
				chromosomes[j] = false
			}
		}
		initialPopulation[i] = &Gene{chromosomes: chromosomes}
	}
	return initialPopulation
}

func (gene *Gene) GetHealth(contextValues []ContextValue) float64 {
	sum := 0.0
	for i, contains := range gene.chromosomes {
		if contains {
			sum += contextValues[i].GetValue()
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
