package tezos

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc1 = `
{
  "tx": {
    "destination": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "amount": "1560",
    "gasLimit": "15385",
    "kind": "transaction",
    "operationResultStatus": "applied",
    "blockHash": "BMJYDJk9wRpxQuuqcFS7MZivqShtrgG18eY5K6rSDBpx5vcJLB2",
    "fee": "0",
    "operationResultConsumedGas": "10207",
    "counter": "1383819",
    "blockLevel": 791441,
    "operationResultErrors": null,
    "blockTimestamp": "2020-01-22T16:34:22Z",
    "parameters": null,
    "source": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "insertedTimestamp": "2020-01-22 16:34:56.015281 UTC"
  },
  "op": {
    "opHash": "ooBAC2ynR5LfU9L2KEon8Z3ujmwEVyB9si1rrppBCmmEn4Mr3bJ",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BMJYDJk9wRpxQuuqcFS7MZivqShtrgG18eY5K6rSDBpx5vcJLB2",
    "blockLevel": 791441,
    "blockTimestamp": "2020-01-22T16:34:22Z",
    "insertedTimestamp": "2020-01-22 16:34:49.405793 UTC"
  }
}
`

const transferSrc2 = `
{
  "tx": {
    "destination": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "amount": "1751",
    "gasLimit": "15385",
    "kind": "transaction",
    "operationResultStatus": "applied",
    "blockHash": "BLJKc6f6SpFs3Jr7WMp2ekx5jQQyzHWN6SWHDo2AJ41HJoPKTF2",
    "fee": "0",
    "operationResultConsumedGas": "10207",
    "counter": "1383094",
    "blockLevel": 788725,
    "operationResultErrors": null,
    "blockTimestamp": "2020-01-20T18:54:52Z",
    "parameters": null,
    "source": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "insertedTimestamp": "2020-01-20 18:56:34.828193 UTC"
  },
  "op": {
    "opHash": "ookuQhVYopxrg8FtnfNASxpMhmnNhBqYaK3PJQXDpP7sDCJAZwf",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BLJKc6f6SpFs3Jr7WMp2ekx5jQQyzHWN6SWHDo2AJ41HJoPKTF2",
    "blockLevel": 788725,
    "blockTimestamp": "2020-01-20T18:54:52Z",
    "insertedTimestamp": "2020-01-20 18:56:28.193751 UTC"
  }
}
`

const transferSrc3 = `
{
  "tx": {
    "destination": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "amount": "1751",
    "gasLimit": "15385",
    "kind": "transaction",
    "operationResultStatus": "backtracked",
    "blockHash": "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
    "fee": "0",
    "operationResultConsumedGas": "10207",
    "counter": "1382930",
    "blockLevel": 788568,
    "operationResultErrors": null,
    "blockTimestamp": "2020-01-20T16:16:32Z",
    "parameters": null,
    "source": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "insertedTimestamp": "2020-01-20 16:19:06.938114 UTC"
  },
  "op": {
    "opHash": "ooXN845juCMcQuqeodwBJWNhY17A5HKyWRxbcwS1m6TfqCqjM8q",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
    "blockLevel": 788568,
    "blockTimestamp": "2020-01-20T16:16:32Z",
    "insertedTimestamp": "2020-01-20 16:18:59.855515 UTC"
  }
}
`

const transferSrc4 = `
{
  "tx": {
    "destination": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "amount": "1751",
    "gasLimit": "15385",
    "kind": "delegation",
    "operationResultStatus": "backtracked",
    "blockHash": "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
    "fee": "0",
    "operationResultConsumedGas": "10207",
    "counter": "1382930",
    "blockLevel": 788568,
    "operationResultErrors": null,
    "blockTimestamp": "2020-01-20T16:16:32Z",
    "parameters": null,
    "source": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "insertedTimestamp": "2020-01-20 16:19:06.938114 UTC"
  },
  "op": {
    "opHash": "ooXN845juCMcQuqeodwBJWNhY17A5HKyWRxbcwS1m6TfqCqjM8q",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
    "blockLevel": 788568,
    "blockTimestamp": "2020-01-20T16:16:32Z",
    "insertedTimestamp": "2020-01-20 16:18:59.855515 UTC"
  }
}
`

var transfer1 = blockatlas.Tx{
	ID:    "ooBAC2ynR5LfU9L2KEon8Z3ujmwEVyB9si1rrppBCmmEn4Mr3bJ",
	Coin:  coin.XTZ,
	Date:  1579710862,
	From:  "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	To:    "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	Fee:   "0",
	Block: 791441,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1560"),
		Symbol:   coin.Coins[coin.XTZ].Symbol,
		Decimals: coin.Coins[coin.XTZ].Decimals,
	},
	Status: blockatlas.StatusCompleted,
}

var transfer2 = blockatlas.Tx{
	ID:    "ookuQhVYopxrg8FtnfNASxpMhmnNhBqYaK3PJQXDpP7sDCJAZwf",
	Coin:  coin.XTZ,
	Date:  1579546492,
	From:  "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	To:    "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	Fee:   "0",
	Block: 788725,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1751"),
		Symbol:   coin.Coins[coin.XTZ].Symbol,
		Decimals: coin.Coins[coin.XTZ].Decimals,
	},
	Status: blockatlas.StatusCompleted,
}

var transfer3 = blockatlas.Tx{
	ID:    "ooXN845juCMcQuqeodwBJWNhY17A5HKyWRxbcwS1m6TfqCqjM8q",
	Coin:  coin.XTZ,
	Date:  1579536992,
	From:  "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	To:    "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	Fee:   "0",
	Block: 788568,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1751"),
		Symbol:   coin.Coins[coin.XTZ].Symbol,
		Decimals: coin.Coins[coin.XTZ].Decimals,
	},
	Status: blockatlas.StatusFailed,
	Error:  "transaction failed",
}

func TestNormalizeTx(t *testing.T) {
	tests := []struct {
		name   string
		srcTx  string
		wantTx blockatlas.Tx
		wantOk bool
	}{
		{"transfer 1", transferSrc1, transfer1, true},
		{"transfer 2", transferSrc2, transfer2, true},
		{"transfer 3", transferSrc3, transfer3, true},
		{"transfer 4", transferSrc4, blockatlas.Tx{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var transaction Transaction
			err := json.Unmarshal([]byte(tt.srcTx), &transaction)
			assert.Nil(t, err)
			gotTx, gotOk := NormalizeTx(transaction)
			assert.Equal(t, gotOk, tt.wantOk, "transfer ok result don't equal")
			assert.Equal(t, gotTx, tt.wantTx, "transfer don't equal")
		})
	}
}

const validatorSrc = `
[
	{"pkh":"tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m","rolls":3726}
]
`

var validator = blockatlas.Validator{
	Status: true,
	ID:     "tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m",
	Details: blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
		MinimumAmount: blockatlas.Amount("0"),
		Type:          blockatlas.DelegationTypeDelegate,
	},
}

func TestNormalizeValidator(t *testing.T) {
	var v []Validator
	err := json.Unmarshal([]byte(validatorSrc), &v)
	assert.Nil(t, err)
	result := normalizeValidator(v[0])
	assert.Equal(t, validator, result)
}

const delegationsSrc = `
{
  "op": {
    "signature": "sigvHd2YBByFXU8nL4CZKSTYXNdMapMsJw1f239YRRjgz9NvrTyA6iGnpBDhi9kCB4zMHysrg9H4jxcpPH975WiQtEmkMjb5",
    "blockUuid": "4b292c55-41ba-4383-a1d6-03fb71b88f41",
    "opHash": "opGphHGNEZZN5rF78yxwe9BJydxYA2yqxECnZR6s6HcxXtCg8Sj",
    "uuid": "e4ec0e07-1601-4da3-bd92-090a820ed369",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BLkscXpE63gajVzmgBS7fQx63hERKQRCZFGtMXdYY6WPThHyji7",
    "protocol": "PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS",
    "branch": "BKqtLegZfdPR3USyYYcMpedB59W5eUBuFZAVpVMPpFgEvMcZjr1",
    "blockLevel": 791778,
    "blockTimestamp": "2020-01-22T22:13:38Z",
    "insertedTimestamp": "2020-01-22 22:14:05.937406 UTC"
  },
  "delegation": {
    "storageLimit": "257",
    "delegate": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "opUuid": "e4ec0e07-1601-4da3-bd92-090a820ed369",
    "uuid": "6459fcd9-5eee-4999-ac4d-92330b9eaab3",
    "gasLimit": "10600",
    "kind": "delegation",
    "operationResultStatus": "applied",
    "fee": "1500",
    "operationResultUuid": "791f6ec7-ecec-43d5-82ca-a1497be0188c",
    "operationResultConsumedGas": "10000",
    "counter": "2409130",
    "operationResultErrors": null,
    "source": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "insertedTimestamp": "2020-01-22 22:15:12.038586 UTC",
    "metadataUuid": "85b42f50-1a89-421c-bcb0-a06926941bc4"
  }
}`

var validator1 = blockatlas.StakeValidator{
	ID:     "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "stake.fish",
		Description: "Leading validator for Proof of Stake blockchains. Stake your cryptocurrencies with us. We know validating.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tezos/validators/assets/tz2fcnbrerxtattnx6iimr1uj5jsdxvdhm93/logo.png",
		Website:     "https://stake.fish/",
	},
	Details: getDetails(),
}

var validatorMap = blockatlas.ValidatorMap{
	"tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93": validator1,
}

var delegationsBalance = "68995611927"

var delegation = blockatlas.DelegationsPage{
	{
		Delegator: validator1,
		Value:     delegationsBalance,
		Status:    blockatlas.DelegationStatusActive,
	},
}

func TestNormalizeDelegations(t *testing.T) {
	var txDelegation TxDelegation
	err := json.Unmarshal([]byte(delegationsSrc), &txDelegation)
	assert.NoError(t, err)
	assert.NotNil(t, txDelegation)
	result, err := NormalizeDelegation(txDelegation, delegationsBalance, validatorMap)
	assert.NoError(t, err)
	assert.Equal(t, delegation, result)
}
