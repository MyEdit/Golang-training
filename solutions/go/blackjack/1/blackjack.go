package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var cardValue int = 0
    
	switch card {
		case "ace":
		    cardValue = 11
		case "two":
		    cardValue = 2
		case "three":
			cardValue = 3
        case "four":
			cardValue = 4
        case "five":
			cardValue = 5
        case "six":
			cardValue = 6
        case "seven":
			cardValue = 7
        case "eight":
			cardValue = 8
        case "nine":
			cardValue = 9
   		case "ten", "jack", "queen", "king":
		    cardValue = 10
		default:
        	cardValue = 0
	}

    return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	var card1Value int = ParseCard(card1)
    var card2Value int = ParseCard(card2)
    var dealerCardValue int = ParseCard(dealerCard)
    var cardsSum int =  card1Value + card2Value
    
	if card1Value == 11 && card2Value == 11 {
    	return "P"
    }

    if cardsSum == 21 {
        if dealerCardValue < 10 {
            return "W"
        } else {
            return "S"
        }
    }

    if cardsSum >= 17 && cardsSum <= 20 {
        return "S"
    }

    if cardsSum >= 12 && cardsSum <= 16 {
        if dealerCardValue >= 7 {
            return "H"
        }
        return "S"
    } else {
        return "H"
    }
}
