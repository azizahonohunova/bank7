package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var PaymentCards []types.PaymentSource
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			PaymentCards = append(PaymentCards, types.PaymentSource{
				Type:    "card",
				Number:  string(card.PAN),
				Balance: card.Balance,
			})
		}
	}
	return PaymentCards
}
