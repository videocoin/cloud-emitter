package bcops

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetBCAuth initialize block chain auth with private key
func GetBCAuth(client *ethclient.Client, privKey *keystore.Key) (*bind.TransactOpts, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, fmt.Errorf("eth client required")
	}

	if privKey == nil {
		return nil, fmt.Errorf("private key required")
	}

	auth := bind.NewKeyedTransactor(privKey.PrivateKey)
	auth.Nonce = nil
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

// LoadBcPrivKeys loads blockchain private keys
// Used in all cloud-blockchain communication
func LoadBcPrivKeys(filename string, password string) (*keystore.Key, error) {
	encrypted, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	decrypted, err := keystore.DecryptKey(encrypted, password)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// RandStringRunes generates a random string with only a-zA-Z
func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
