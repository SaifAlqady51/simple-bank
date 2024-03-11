package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTX execute a function within database transaction
func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error {

	// start a new transaction
	tx, err := store.db.BeginTx(ctx, nil)

	// return err if exists
	if err != nil {
		return err
	}
	// create new Query from tx db
	q := New(tx)

	// execute fn function that takes query and returns error if exists
	err = fn(q)

	if err != nil {
		// if any error occurs during the transaction rollback the changes
		if rbErr := tx.Rollback(); rbErr != nil {
			// print out the error
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	// if transaction succedded Commit the transation
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer
	FromAccount Account
	ToAccount   Account
	FromEntry   Entry
	ToEntry     Entry
}

//TransferTx performs a money transfer from one account another

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	// initiat result variable
	var result TransferTxResult

	// check if store.execTX would returns an error
	err := store.execTX(ctx, func(q *Queries) error {
		var err error
		// get result.Transfer from Create Transfer option or get the error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		// return err if exists
		if err != nil {
			return err
		}
		// get FromEntry prop from CreateEntry function and pass its valut to  result.FromEntry *TransferTxResult
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		// get ToEntry prop from CreateEntry function and pass its valut to  result.ToEntry *TransferTxResult
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountId, -arg.Amount, arg.ToAccountId, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountId, arg.Amount, arg.FromAccountId, -arg.Amount)
		}
		return nil
	})
	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountID1 int64, amount1 int64, accountID2 int64, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return

}
