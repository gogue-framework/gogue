package randomNumberGenerator

import (
	"time"
	"math/rand"
)

type DiceRoller struct {
	rng RNG
}

func NewDiceRoller() *DiceRoller {
	diceRoller := DiceRoller{}

	// Set the seed to the current time. This can be updated later by the user.
	diceRoller.rng.seed = time.Now().UTC().UnixNano()
	diceRoller.rng.rand = rand.New(rand.NewSource(diceRoller.rng.seed))

	return &diceRoller
}

func (diceRoller *DiceRoller) GetSeed() int64 {
	return diceRoller.rng.seed
}

func (diceRoller *DiceRoller) SetSeed(seed int64) {
	diceRoller.rng.seed = seed
}

func (diceRoller *DiceRoller) RollNSidedDie(n int) int {
	return diceRoller.rng.Range(n) + 1
}

func (diceRoller *DiceRoller) RollNSidedDieOpenEnded(n int) int {
	roll := diceRoller.rng.Range(n) + 1
	totalRoll := roll

	for roll == n {
		roll = diceRoller.rng.Range(n) + 1
		totalRoll += roll
	}

	return totalRoll
}