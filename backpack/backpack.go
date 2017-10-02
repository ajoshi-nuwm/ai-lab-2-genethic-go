/*
This package contains context of the problem.

backpack.go contains
 */
package backpack

import (
	"github.com/ajoshi-nuwm/ai-lab-2-genethic-go/gene"
	"math/rand"
	"fmt"
)

type Backpack struct {
	TotalWeight float64
	items       []*Item
	genes       []*gene.Gene
}

func NewBackPack(totalWeight float64, items []*Item, genes []*gene.Gene) *Backpack {
	return &Backpack{totalWeight, items, genes}
}

func (backpack *Backpack) NextGeneration() {
	fmt.Println("before")
	fmt.Println(backpack.genes)
	splitPoint := rand.Intn(backpack.genes[0].GetLength())
	parents := backpack.getCurrentParents()
	newGeneration := make([]*gene.Gene, 0)

	fmt.Println("parents")
	fmt.Println(parents)
	for i := 0; i < len(parents); i += 2 {
		children1, children2 := parents[i/2].Crossover(parents[i/2+1], splitPoint)
		newGeneration = append(newGeneration, children1)
		newGeneration = append(newGeneration, children2)
	}

	for _, parent := range parents {
		if rand.Float64() > 0.5 {
			newGeneration = append(newGeneration, parent.Mutation())
		} else {
			newGeneration = append(newGeneration, parent)
		}
	}

	backpack.genes = newGeneration
	fmt.Println("after")
	fmt.Println(backpack.genes)
}

func (backpack *Backpack) getCurrentParents() []*gene.Gene {
	parents := make([]*gene.Gene, 0)
	averageHealth := backpack.getAverageHealth()

	for _, currentGene := range backpack.genes {
		if currentGene.GetHealth(backpack.getItemsAsContextValues()) > averageHealth {
			parents = append(parents, currentGene)
		}
	}
	return parents
}

func (backpack *Backpack) getAverageHealth() float64 {
	sum := 0.0
	for _, currentGene := range backpack.genes {
		sum += currentGene.GetHealth(backpack.getItemsAsContextValues())
	}
	return sum / float64(len(backpack.items))
}

func (backpack *Backpack) getItemsAsContextValues() []gene.ContextValue {
	contextValues := make([]gene.ContextValue, len(backpack.items))
	for i, item := range backpack.items {
		contextValues[i] = *item
	}
	return contextValues
}
