package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/SaifAlqady51/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQuery.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGettransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQuery.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)

	require.WithinDuration(t, transfer1.CreatedAt, transfer1.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: transfer1.Amount,
	}

	transfer2, err := testQuery.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, arg.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestDeletetransfer(t *testing.T) {

	transfer1 := createRandomTransfer(t)

	err := testQuery.DeleteTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)

	transfer2, err := testQuery.GetTransfer(context.Background(), transfer1.ID)

	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)

}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQuery.ListTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
