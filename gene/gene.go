package gene

import (
	"github.com/ajoshi-nuwm/ai-lab-2-genethic-go/backpack"
	"math/rand"
	"fmt"
	"strings"
)

type Gene struct {
	Chromosomes []bool
}

func (gene *Gene) Health(items []*backpack.Item) float64 {
	price := 0.0
	for i, contains := range gene.Chromosomes {
		if contains {
			price += items[i].GetPrice()
		}
	}
	return price
}

func (gene *Gene) Crossover(other *Gene, splitPoint int) (first, second *Gene) {
	first = &Gene{}
	second = &Gene{}

	first.Chromosomes = append(first.Chromosomes, gene.Chromosomes[:splitPoint]...)
	first.Chromosomes = append(first.Chromosomes, other.Chromosomes[splitPoint:]...)

	second.Chromosomes = append(second.Chromosomes, other.Chromosomes[:splitPoint]...)
	second.Chromosomes = append(second.Chromosomes, gene.Chromosomes[splitPoint:]...)

	return
}

func (gene *Gene) Mutation() *Gene {
	mutationIndex := rand.Intn(len(gene.Chromosomes))
	gene.Chromosomes[mutationIndex] = !gene.Chromosomes[mutationIndex]
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
		initialPopulation[i] = &Gene{Chromosomes: chromosomes}
	}
	return initialPopulation
}

func GetNewGeneration(genes []*Gene, items []*backpack.Item, backPack backpack.Backpack) []*Gene {

}

func (gene Gene) String() string {
	stringValues := make([]string, len(gene.Chromosomes))
	for i, value := range gene.Chromosomes {
		if value {
			stringValues[i] = "1"
		} else {
			stringValues[i] = "0"
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(stringValues, ""))
}
