package randomNumberGenerator

import (
	"time"
	"math/rand"
	"math"
)

type RNG struct {
	seed int64
	rand *rand.Rand
}

func NewRNG() *RNG {
	rng := RNG{}

	// Set the seed to the current time. This can be updated later by the user.
	rng.seed = time.Now().UTC().UnixNano()
	rng.rand = rand.New(rand.NewSource(rng.seed))

	return &rng
}

func (rng *RNG) GetSeed() int64 {
	return rng.seed
}

func (rng *RNG) SetSeed(seed int64) {
	rng.seed = seed
}

func (rng *RNG) Uniform() float64 {
	return rng.rand.Float64()
}

func (rng *RNG) UniformRange(a, b float64) float64 {
	return a + rng.Uniform() * (b - a)
}

func (rng *RNG) Normal(mean, stddev float64) float64 {

	var r, x float64

	for r >= 1 || r == 0 {
		x = rng.UniformRange(-1.0, 1.0)
		y := rng.UniformRange(-1.0, 1.0)
		r = x*x + y*y
	}

	result := x * math.Sqrt(-2 * math.Log(r) / r)

	return mean + stddev * result

}

func (rng *RNG) Percentage() int {
	return rng.rand.Intn(100)
}

func (rng *RNG) Range(n int) int {
	return rng.rand.Intn(n)
}

func (rng *RNG) GetWeightedEntity(values map[int]int) int {
	// First up, get the total weight value from the map
	totalWeight := 0
	for weight := range values {
		totalWeight += weight
	}

	// Next, get a random integer in the range of the total weight
	r := rng.Range(totalWeight)

	for weight, value := range values {
		r -= weight
		if r <= 0 {
			return value
		}
	}

	return -1
}

