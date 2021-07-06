package tests

import (
	"testing"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"fmt"
)

func TestBox(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(10000000000)},
	}

	sim := backends.NewSimulatedBackend(gAlloc, 1000000)

	_, _, box, err := DeployBox(auth, sim)
	require.NoError(t, err)

	sim.Commit()


	val := big.NewInt(1234)

	transOpts := &bind.TransactOpts{
		From: auth.From,
		Signer: auth.Signer,
		GasLimit: 123456,
	}
	tx, err := box.Store(transOpts, val)
	require.NoError(t, err)

	fmt.Printf("%+v", tx)


	sim.Commit()

	actualVal, err := box.Retrieve(&bind.CallOpts{
		Pending: true,
	})
	require.NoError(t, err)

	require.Equal(t, val, actualVal)

}
