package blackjack

func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	dealerCardVal := ParseCard(dealerCard)
	cardsSumValue := card1Val + card2Val

	switch {
	case cardsSumValue == 21:
		if dealerCardVal >= 10 {
			return "S"
		} else {
			return "W"
		}
	case cardsSumValue >= 17 && cardsSumValue <= 20:
		return "S"
	case cardsSumValue >= 12 && cardsSumValue <= 16:
		if dealerCardVal >= 7 {
			return "H"
		} else {
			return "S"
		}
	case cardsSumValue <= 11:
		return "H"
	case card1 == "ace", card2 == "ace":
		return "P"
	default:
		return ""
	}
}
