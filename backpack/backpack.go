/*
This package contains context of the problem.

backpack.go contains
 */
package backpack

import (
	"github.com/ajoshi-nuwm/ai-lab-2-genethic-go/gene"
	"math/rand"
	"fmt"
	"github.com/ajoshi-nuwm/ai-lab-2-genethic-go/utils"
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

	genesCopy := make([]*gene.Gene, len(backpack.genes))
	copy(genesCopy, backpack.genes)
	backpack.getNextParent(genesCopy)
	return backpack.genes[:len(backpack.genes) / 2]
}

func (backpack *Backpack) getNextParent(genes []*gene.Gene) {
	objects := backpack.getGenesAsObjects()
	for _, currentGene := range backpack.genes {
		rule := util.GetObjectProbailityRule(func(object util.Object) float64 {
			return currentGene.GetHealth(backpack.getItemsAsContextValues()) / currentGene.GetDecease(backpack.getItemsAsContextValues())
		}, objects)

		currentGene, ok := rule.(gene.Gene)
		fmt.Println(rule)
		if ok {
			fmt.Println("HERE WE GO")
			fmt.Println(currentGene)
		}

	}
}

func (backpack *Backpack) getItemsAsContextValues() []gene.ContextValue {
	contextValues := make([]gene.ContextValue, len(backpack.items))
	for i, item := range backpack.items {
		contextValues[i] = *item
	}
	return contextValues
}

func (backpack *Backpack) getGenesAsObjects() []util.Object {
	objects := make([]util.Object, len(backpack.genes))
	for i, gene := range backpack.genes {
		objects[i] = *gene
	}
	return objects
}
