package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	m := float64((score - 10)) / 2
	return int(math.Floor(m))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	sum := 0
	values := []int{
		rollDice(),
		rollDice(),
		rollDice(),
		rollDice(),
	}

	slices.Sort(values)
	for i, v := range values {
		if i == 0 {
			continue
		}

		sum += v
	}

	return sum
}

func rollDice() int {
	return rand.Intn(6) + 1
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{}
	c.Charisma = Ability()
	c.Constitution = Ability()
	c.Dexterity = Ability()
	c.Intelligence = Ability()
	c.Strength = Ability()
	c.Wisdom = Ability()
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}
