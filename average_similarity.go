package anomaly

import (
	"math"
	"math/rand"
)

const vectorsSize = 1024

// AverageSimilarity computes surpise by calculation the average cosine
// similarity across all Vectors
type AverageSimilarity struct {
	vectors       [][]float32
	begin, length int
}

// NewAverageSimilarity creates a new average similarity surprise engine
func NewAverageSimilarity(width int, rnd *rand.Rand) Network {
	return &AverageSimilarity{
		vectors: make([][]float32, vectorsSize),
	}
}

// Train computes the surprise with average similarity
func (a *AverageSimilarity) Train(input []float32) float32 {
	sum, c := 0.0, a.begin
	for i := 0; i < a.length; i++ {
		sum += math.Abs(Similarity(input, a.vectors[c]))
		c = (c + 1) % vectorsSize
	}
	averageSimilarity := float32(sum / float64(a.length))

	if a.length < vectorsSize {
		a.vectors[a.begin+a.length] = input
		a.length++
	} else {
		a.vectors[a.begin] = input
		a.begin = (a.begin + 1) % vectorsSize
	}

	return averageSimilarity
}
