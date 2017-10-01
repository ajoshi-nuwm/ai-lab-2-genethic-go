package gene

import (
	"github.com/ajoshi-nuwm/ai-lab-2-genethic-go/backpack"
	"math/rand"
)

type GeneImpl struct {
	chromosomes []bool
}

func (gene *GeneImpl) Health(items []*backpack.Item) float64 {
	price := 0.0
	for i, contains := range gene.chromosomes {
		if contains {
			price += items[i].GetPrice()
		}
	}
	return price
}

func (gene *GeneImpl) Crossover(other *GeneImpl, splitPoint int) (first, second *Gene) {
	first = &GeneImpl{}
	second = &GeneImpl{}

	first.chromosomes = append(first.chromosomes, gene.chromosomes[:splitPoint]...)
	first.chromosomes = append(first.chromosomes, other.chromosomes[splitPoint:]...)

	second.chromosomes = append(second.chromosomes, other.chromosomes[:splitPoint]...)
	second.chromosomes = append(second.chromosomes, gene.chromosomes[splitPoint:]...)

	return
}

func (gene *GeneImpl) Mutation() *Gene {
	mutationIndex := rand.Intn(len(gene.chromosomes))
	gene.chromosomes[mutationIndex] = !gene.chromosomes[mutationIndex]
	return gene
}
