package main

import (
	"log"
	"strconv"

	"github.com/ccrayz/account-manager/account"
	"github.com/ccrayz/account-manager/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	mnemonic := config.Get("eth.mnemonic")
	if mnemonic == "" {
		log.Fatalf("Failed to get mnemonic: %v", mnemonic)
	}

	countStr := config.Get("eth.count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatalf("Failed to convert count to integer: %v", err)
	}

	accounts := account.GetAccounts(mnemonic, count)
	for _, a := range accounts {
		a.ShowAccount()
	}
}
