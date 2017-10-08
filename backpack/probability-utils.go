package backpack

import (
	"math/rand"
	"fmt"
	"math"
)

type Object interface{}

// Segment of probability ruler
type Segment struct {
	gene  Gene
	start float64
	end   float64
}

// Accepts closure for probability calculation and list of objects
func GetObjectProbabilityRule(getProbability func(gene Gene) float64, genes []Gene) (int, Gene) {
	segments := []Segment{}
	var start float64
	for _, gene := range genes {
		probability := getProbability(gene)
		if math.IsNaN(probability) {
			probability = 0
		}
		segments = append(segments, Segment{gene, start, start + probability})
		start += probability
	}
	random := rand.Float64() * start
	for index, segment := range segments {
		if segment.isInSegment(random) {
			return index, segment.gene
		}
	}
	fmt.Println(segments)
	panic("Can not choose gene")
}

// Checks if number is in segment
func (segment *Segment) isInSegment(value float64) bool {
	return value >= segment.start && value <= segment.end
}
