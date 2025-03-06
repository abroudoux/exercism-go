package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
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

func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10.0) / 2.0))
}

func Ability() int {
	var randNumbers []int
	for i := 0; i < 4; i++ {
		randNumbers = append(randNumbers, rand.Intn(6)+1)
	}

	sort.Ints(randNumbers)

	var sum int
	for i := 1; i < len(randNumbers); i++ {
		sum += randNumbers[i]
	}

	return sum
}

func GenerateCharacter() Character {
	constitutionScore := Ability()

	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitutionScore,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitutionScore),
	}
}
