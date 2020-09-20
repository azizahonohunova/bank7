package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	card := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 30_000_03,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: -50_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 2020",
			Balance: 40_000_04,
			Active:  false,
		},
	}
	result := PaymentSources(cards)
	fmt.Println(result[0].Number)
	//Output: 5058 xxxx xxxx 8888

}
