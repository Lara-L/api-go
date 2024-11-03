package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"maria": {
		AuthToken: "123ABC",
		Username:  "maria",
	},
	"jose": {
		AuthToken: "def456",
		Username:  "jose",
	},
	"joao": {
		AuthToken: "qwe123",
		Username:  "joao",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"maria": {
		Coins:    100,
		Username: "maria",
	},
	"jose": {
		Coins:    200,
		Username: "jose",
	},
	"joao": {
		Coins:    50,
		Username: "joao",
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
