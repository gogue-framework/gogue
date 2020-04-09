package randomnumbergenerator

import (
	"math/rand"
	"time"
)

// DiceRoller contains various methods of rolling different types of dice. It uses a Gogue RNG to determine randomness
type DiceRoller struct {
	rng RNG
}

// NewDiceRoller creates a new DiceRoller. It sets the seed for the contained RNG to the current UNIX timestamp
func NewDiceRoller() *DiceRoller {
	diceRoller := DiceRoller{}

	// Set the seed to the current time. This can be updated later by the user.
	diceRoller.rng.seed = time.Now().UTC().UnixNano()
	diceRoller.rng.rand = rand.New(rand.NewSource(diceRoller.rng.seed))

	return &diceRoller
}

// GetSeed returns the RNG seed
func (diceRoller *DiceRoller) GetSeed() int64 {
	return diceRoller.rng.seed
}

// SetSeed sets the RNG seed
func (diceRoller *DiceRoller) SetSeed(seed int64) {
	diceRoller.rng.seed = seed
}

// RollNSidedDie rolls a die with N sides
func (diceRoller *DiceRoller) RollNSidedDie(n int) int {
	return diceRoller.rng.Range(0, n) + 1
}

// RollNSidedDieOpenEnded rolls a die with N sides. If the maximum value of the die is rolled, the value is accumulated,
// and the die is rolled again. Any time the max die value is rolled, this process is repeated. This is useful in
// situations where something should be improbable, rather than impossible.
func (diceRoller *DiceRoller) RollNSidedDieOpenEnded(n int) int {
	roll := diceRoller.rng.Range(0, n) + 1
	totalRoll := roll

	for roll == n {
		roll = diceRoller.rng.Range(0, n) + 1
		totalRoll += roll
	}

	return totalRoll
}
