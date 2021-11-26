package accounts

import (
	"math/big"
	"testing"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	t.Logf("Name: %s", Contract.Name)
	t.Logf("Description: %s", Contract.Description)
	t.Logf("Program hash: %s", Contract.ProgramHash.String())
	t.Logf("Hname: %s", Contract.Hname())
}

var dummyAssetID = [iotago.NativeTokenIDLength]byte{1, 2, 3}

func checkLedger(t *testing.T, state dict.Dict, cp string) *iscp.Assets {
	total := GetTotalAssets(state)
	// t.Logf("checkpoint '%s.%s':\n%s", curTest, cp, total.String())
	require.NotPanics(t, func() {
		mustCheckLedger(state, cp)
	})
	return total
}

var curTest = ""

func TestCreditDebit1(t *testing.T) {
	curTest = "TestCreditDebit1"
	state := dict.New()
	total := checkLedger(t, state, "cp0")

	require.True(t, total.Equals(iscp.NewEmptyAssets()))

	agentID1 := iscp.NewRandomAgentID()
	transfer := iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	CreditToAccount(state, agentID1, transfer)
	total = checkLedger(t, state, "cp1")

	require.NotNil(t, total)
	require.EqualValues(t, 1, len(total.Tokens))
	require.True(t, total.Equals(transfer))

	transfer.Iotas = 1
	CreditToAccount(state, agentID1, transfer)
	total = checkLedger(t, state, "cp2")

	expected := iscp.NewAssets(43, nil).AddToken(dummyAssetID, big.NewInt(4))
	require.True(t, expected.Equals(total))

	userAssets := GetAssets(state, agentID1)
	require.EqualValues(t, 43, userAssets.Iotas)
	require.Zero(t, userAssets.Tokens.MustSet()[dummyAssetID].Amount.Cmp(big.NewInt(4)))
	checkLedger(t, state, "cp2")

	DebitFromAccount(state, agentID1, expected)
	total = checkLedger(t, state, "cp3")
	expected = iscp.NewEmptyAssets()
	require.True(t, expected.Equals(total))
}

func TestCreditDebit2(t *testing.T) {
	curTest = "TestCreditDebit2"
	state := dict.New()
	total := checkLedger(t, state, "cp0")
	require.True(t, total.Equals(iscp.NewEmptyAssets()))

	agentID1 := iscp.NewRandomAgentID()
	transfer := iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	CreditToAccount(state, agentID1, transfer)
	total = checkLedger(t, state, "cp1")

	expected := transfer
	require.EqualValues(t, 1, len(total.Tokens))
	require.True(t, expected.Equals(total))

	transfer = iscp.NewEmptyAssets().AddToken(dummyAssetID, big.NewInt(2))
	DebitFromAccount(state, agentID1, transfer)
	total = checkLedger(t, state, "cp2")
	require.EqualValues(t, 0, len(total.Tokens))
	expected = iscp.NewAssets(42, nil)
	require.True(t, expected.Equals(total))

	require.True(t, util.IsZeroBigInt(GetTokenBalance(state, agentID1, &transfer.Tokens[0].ID)))
	bal1, ok := GetAccountBalances(state, agentID1)
	require.True(t, ok)
	require.True(t, total.Equals(bal1))
}

func TestCreditDebit3(t *testing.T) {
	curTest = "TestCreditDebit3"
	state := dict.New()
	total := checkLedger(t, state, "cp0")
	require.True(t, total.Equals(iscp.NewEmptyAssets()))

	agentID1 := iscp.NewRandomAgentID()
	transfer := iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	CreditToAccount(state, agentID1, transfer)
	total = checkLedger(t, state, "cp1")

	expected := transfer
	require.EqualValues(t, 1, len(total.Tokens))
	require.True(t, expected.Equals(total))

	transfer = iscp.NewEmptyAssets().AddToken(dummyAssetID, big.NewInt(100))
	require.Panics(t,
		func() {
			DebitFromAccount(state, agentID1, transfer)
		},
	)
	total = checkLedger(t, state, "cp2")

	require.EqualValues(t, 1, len(total.Tokens))
	expected = iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	require.True(t, expected.Equals(total))
}

func TestCreditDebit4(t *testing.T) {
	curTest = "TestCreditDebit4"
	state := dict.New()
	total := checkLedger(t, state, "cp0")
	require.True(t, total.Equals(iscp.NewEmptyAssets()))

	agentID1 := iscp.NewRandomAgentID()
	transfer := iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	CreditToAccount(state, agentID1, transfer)
	total = checkLedger(t, state, "cp1")

	expected := transfer
	require.EqualValues(t, 1, len(total.Tokens))
	require.True(t, expected.Equals(total))

	keys := getAccountsIntern(state).Keys()
	require.EqualValues(t, 1, len(keys))

	agentID2 := iscp.NewRandomAgentID()
	require.NotEqualValues(t, agentID1, agentID2)

	transfer = iscp.NewAssets(20, nil)
	ok := MoveBetweenAccounts(state, agentID1, agentID2, transfer)
	require.True(t, ok)
	total = checkLedger(t, state, "cp2")

	keys = getAccountsIntern(state).Keys()
	require.EqualValues(t, 2, len(keys))

	expected = iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	require.True(t, expected.Equals(total))

	bm1, ok := GetAccountBalances(state, agentID1)
	require.True(t, ok)
	expected = iscp.NewAssets(22, nil).AddToken(dummyAssetID, big.NewInt(2))
	require.True(t, expected.Equals(bm1))

	bm2, ok := GetAccountBalances(state, agentID2)
	require.True(t, ok)
	expected = iscp.NewAssets(20, nil)
	require.True(t, expected.Equals(bm2))
}

func TestCreditDebit5(t *testing.T) {
	curTest = "TestCreditDebit5"
	state := dict.New()
	total := checkLedger(t, state, "cp0")
	require.True(t, total.Equals(iscp.NewEmptyAssets()))

	agentID1 := iscp.NewRandomAgentID()
	transfer := iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	CreditToAccount(state, agentID1, transfer)
	total = checkLedger(t, state, "cp1")

	expected := transfer
	require.EqualValues(t, 1, len(total.Tokens))
	require.True(t, expected.Equals(total))

	keys := getAccountsIntern(state).Keys()
	require.EqualValues(t, 1, len(keys))

	agentID2 := iscp.NewRandomAgentID()
	require.NotEqualValues(t, agentID1, agentID2)

	transfer = iscp.NewAssets(50, nil)
	ok := MoveBetweenAccounts(state, agentID1, agentID2, transfer)
	require.False(t, ok)
	total = checkLedger(t, state, "cp2")

	keys = getAccountsIntern(state).Keys()
	require.EqualValues(t, 1, len(keys))

	expected = iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	require.True(t, expected.Equals(total))

	bm1, ok := GetAccountBalances(state, agentID1)
	require.True(t, ok)
	require.True(t, expected.Equals(bm1))

	_, ok = GetAccountBalances(state, agentID2)
	require.False(t, ok)
}

func TestCreditDebit6(t *testing.T) {
	curTest = "TestCreditDebit6"
	state := dict.New()
	total := checkLedger(t, state, "cp0")
	require.True(t, total.Equals(iscp.NewEmptyAssets()))

	agentID1 := iscp.NewRandomAgentID()
	transfer := iscp.NewAssets(42, nil).AddToken(dummyAssetID, big.NewInt(2))
	CreditToAccount(state, agentID1, transfer)
	checkLedger(t, state, "cp1")

	agentID2 := iscp.NewRandomAgentID()
	require.NotEqualValues(t, agentID1, agentID2)

	ok := MoveBetweenAccounts(state, agentID1, agentID2, transfer)
	require.True(t, ok)
	total = checkLedger(t, state, "cp2")

	keys := getAccountsIntern(state).Keys()
	require.EqualValues(t, 1, len(keys))

	_, ok = GetAccountBalances(state, agentID1)
	require.False(t, ok)

	bal2, ok := GetAccountBalances(state, agentID2)
	require.True(t, ok)
	require.True(t, total.Equals(bal2))
}

func TestCreditDebit7(t *testing.T) {
	curTest = "TestCreditDebit7"
	state := dict.New()
	total := checkLedger(t, state, "cp0")
	require.True(t, total.Equals(iscp.NewEmptyAssets()))

	agentID1 := iscp.NewRandomAgentID()
	transfer := iscp.NewEmptyAssets().AddToken(dummyAssetID, big.NewInt(2))
	CreditToAccount(state, agentID1, transfer)
	checkLedger(t, state, "cp1")

	debitTransfer := iscp.NewAssets(1, nil)
	// debit must fail
	require.Panics(t, func() {
		DebitFromAccount(state, agentID1, debitTransfer)
	})

	total = checkLedger(t, state, "cp1")
	require.True(t, transfer.Equals(total))
}
