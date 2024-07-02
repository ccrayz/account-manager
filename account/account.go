package account

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Account struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

func NewAccount(address string, privateKey string) *Account {
	return &Account{
		Address:    address,
		PrivateKey: privateKey,
	}
}

func (a *Account) ShowAccount() {
	fmt.Println("addr", a.Address, "private", a.PrivateKey)
}

func GetAccounts(mnemonic string, rangekey int) []Account {
	accounts := make([]Account, rangekey)
	for i := 0; i < rangekey; i++ {
		account := *getAccount(getAccountToMnemonic(mnemonic), uint32(i))
		accounts[i] = account
	}
	return accounts
}

func getAccount(account *bip32.Key, accountIndex uint32) *Account {
	change, _ := account.NewChildKey(accountIndex)

	addressIndex, err := change.NewChildKey(0)
	if err != nil {
		log.Fatalf("Failed to derive address index key: %v", err)
	}
	privateKey, err := crypto.ToECDSA(addressIndex.Key)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return NewAccount(address.Hex(), hexutil.Encode(crypto.FromECDSA(privateKey)))
}

func getAccountToMnemonic(mnemonic string) *bip32.Key {
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatalf("Failed to create master key: %v", err)
	}
	purpose, _ := masterKey.NewChildKey(44 | bip32.FirstHardenedChild)
	coinType, _ := purpose.NewChildKey(60 | bip32.FirstHardenedChild)
	account, _ := coinType.NewChildKey(0 | bip32.FirstHardenedChild)

	return account
}
