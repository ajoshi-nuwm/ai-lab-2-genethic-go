package util

import "math/rand"

type Object interface{}

// Segment of probability ruler
type Segment struct {
	ruleObject Object
	start      float64
	end        float64
}

// Accepts closure for probability calculation and list of objects
func GetObjectProbailityRule(getProbability func(object Object) float64, objects []Object) Object {
	segments := []Segment{}
	var start float64
	for _, object := range objects {
		probability := getProbability(object)
		segments = append(segments, Segment{object, start, start + probability})
		start += probability
	}
	random := rand.Float64() * start
	for _, segment := range segments {
		if segment.isInSegment(random) {
			return segment.ruleObject
		}
	}
	return nil
}

// Checks if number is in segment
func (segment *Segment) isInSegment(value float64) bool {
	return value >= segment.start && value <= segment.end
}
