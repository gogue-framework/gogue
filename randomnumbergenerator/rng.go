package randomnumbergenerator

import (
	"math"
	"math/rand"
	"time"
)

// RNG is a random number generator. It contains a seed, and the Go standardlib random number generator object. It is
// to generate a variety of random values.
type RNG struct {
	seed int64
	rand *rand.Rand
}

// NewRNG creates a new RNG. A seed is set to the current Unix timestamp.
func NewRNG() *RNG {
	rng := RNG{}

	// Set the seed to the current time. This can be updated later by the user.
	rng.seed = time.Now().UTC().UnixNano()
	rng.rand = rand.New(rand.NewSource(rng.seed))

	return &rng
}

// GetSeed returns the seed value for the RNG
func (rng *RNG) GetSeed() int64 {
	return rng.seed
}

// SetSeed sets the seed value for the RNG
func (rng *RNG) SetSeed(seed int64) {
	rng.seed = seed
}

// Uniform returns a uniform random value in the range [0.0, 1.0]
func (rng *RNG) Uniform() float64 {
	return rng.rand.Float64()
}

// UniformRange returns a uniform random value across the defined range
func (rng *RNG) UniformRange(a, b float64) float64 {
	return a + rng.Uniform()*(b-a)
}

// Normal returns a random value from a normal (Gaussian) distribution
func (rng *RNG) Normal(mean, stddev float64) float64 {

	var r, x float64

	for r >= 1 || r == 0 {
		x = rng.UniformRange(-1.0, 1.0)
		y := rng.UniformRange(-1.0, 1.0)
		r = x*x + y*y
	}

	result := x * math.Sqrt(-2*math.Log(r)/r)

	return mean + stddev*result

}

// Percentage returns a value in the range [0, 100]
func (rng *RNG) Percentage() int {
	return rng.rand.Intn(100)
}

// Range returns a value in the defined range
func (rng *RNG) Range(min, max int) int {
	if min == max {
		return min
	}

	return rng.rand.Intn(max-min) + min
}

// RangeNegative returns a value in the defined range, but allows negative values
func (rng *RNG) RangeNegative(min, max int) int {
	if min == max {
		return min
	}

	return rng.rand.Intn(max-min+1) + min
}

// GetWeightedEntity takes a weight map (a map of entity IDs with a weight associated with them). It then selects an
// entity based on the weights. An entity with a higher weight is more likely to be chosen over a lower weight entity.
func (rng *RNG) GetWeightedEntity(values map[int]int) int {
	// First up, get the total weight value from the map
	totalWeight := 0
	for weight := range values {
		totalWeight += weight
	}

	// Next, get a random integer in the range of the total weight
	r := rng.Range(0, totalWeight)

	for weight, value := range values {
		r -= value
		if r <= 0 {
			return weight
		}
	}

	return -1
}
