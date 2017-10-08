/*
This package contains context of the problem.

backpack.go contains
 */
package backpack

import (
	"math/rand"
	"fmt"
)

type Backpack struct {
	TotalWeight float64
	items       []Item
	genes       []Gene
}

func NewBackPack(totalWeight float64, items []Item, genes []Gene) *Backpack {
	return &Backpack{totalWeight, items, genes}
}

func (backpack *Backpack) NextGeneration() {
	newGeneration := make([]Gene, 0)
	parents := backpack.getCurrentParents()

	for i := 0; i < len(parents); i += 2 {
		children1, children2 := parents[i].Crossover(&parents[i+1], rand.Intn(parents[i].GetLength()))
		newGeneration = append(newGeneration, children1, children2)
		newGeneration = append(newGeneration, *parents[i].Mutation(), *parents[i+1].Mutation())
	}
	backpack.genes = newGeneration
}

func (backpack *Backpack) getCurrentParents() []Gene {
	genesCopy := make([]Gene, len(backpack.genes))
	parents := make([]Gene, 0)
	copy(genesCopy, backpack.genes)

	for i := 0; i < len(backpack.genes)/2; i++ {
		parentIndex, parent := backpack.getNextParent(genesCopy)
		genesCopy = cutGeneFromSlice(genesCopy, parentIndex)
		parents = append(parents, parent)
	}
	return parents
}

func (backpack *Backpack) getNextParent(genes []Gene) (int, Gene) {
	return GetObjectProbabilityRule(func(gene Gene) float64 {
		if backpack.TotalWeight < gene.GetDecease(backpack.items) {
			return 1 / gene.GetDecease(backpack.items)
		}
		return gene.GetHealth(backpack.items) / gene.GetDecease(backpack.items)
	}, genes)
}

func (backpack *Backpack) PrintSolution() {
	var bestGene *Gene = GetMinimalGene(backpack.items, backpack.TotalWeight)
	fmt.Println(backpack.TotalWeight)
	for i := 0; i < 100000; i++ {
		backpack.NextGeneration()

		for _, gene := range backpack.genes {
			if gene.GetDecease(backpack.items) <= backpack.TotalWeight && gene.GetHealth(backpack.items) > bestGene.GetHealth(backpack.items) {
				fmt.Println(gene)
				fmt.Println(gene.GetHealth(backpack.items))
				fmt.Println(gene.GetDecease(backpack.items))
				bestGene = &gene
			}
		}
	}
}

func cutGeneFromSlice(genes []Gene, cutPoint int) []Gene {
	switch cutPoint {
	case 0:
		return genes[1:]
	case len(genes) - 1:
		return genes[:len(genes)-2]
	default:
		return append(genes[:cutPoint], genes[cutPoint+1:]...)
	}
}
