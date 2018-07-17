package backpack

import (
	"math/rand"
	"fmt"
	"math"
)

// Segment of probability ruler
type Segment struct {
	start float64
	end   float64
}

// Accepts closure for probability calculation and list of objects
func GetObjectProbabilityRule(len int, probability func(int) float64, target func(int)) {
	var segments []Segment
	var start float64
	for i := 0; i < len; i++ {
		probability := probability(i)
		if math.IsNaN(probability) {
			probability = 0
		}
		segments = append(segments, Segment{start, start + probability})
		start += probability
	}
	random := rand.Float64() * start
	for index, segment := range segments {
		if segment.isInSegment(random) {
			target(index)
			return
		}
	}
	fmt.Println(segments)
	panic("Can not choose object")
}

// Checks if number is in segment
func (segment *Segment) isInSegment(value float64) bool {
	return value >= segment.start && value <= segment.end
}
