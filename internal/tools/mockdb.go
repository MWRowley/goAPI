package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Matthew": {
		AuthToken: "123ABC",
		Username:  "Matthew",
	},
	"Katie": {
		AuthToken: "qazwsx123",
		Username:  "Katie",
	},
	"Nick": {
		AuthToken: "foobar",
		Username:  "Nick",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Matthew": {
		Coins:    100,
		Username: "Matthew",
	},
	"Katie": {
		Coins:    500,
		Username: "Katie",
	},
	"Nick": {
		Coins:    0,
		Username: "Nick",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
