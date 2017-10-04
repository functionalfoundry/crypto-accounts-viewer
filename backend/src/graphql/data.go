package graphql

import (
	"../data"
)

func GetCurrencies() []data.Currency {
	return data.Currencies[:]
}

func GetExchanges() []data.Exchange {
	return data.Exchanges[:]
}
