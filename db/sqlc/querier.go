// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	// Instead of doing get and update function seperately - we are performing one single update function
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	// don't mess with the comments - they are ment to say how to generate the files - each word has it's meaning
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	// This is to block the concurrent accecc of sinle account - when one transaction is going on - because we are using for update account - we could block the access from other account
	// To avoid deadlock - after doing the required debugging we found that - updating only changes the account balance. The account id will never be changed because it's the primary key of accounts table
	// So if i'm telling that I'm selecting this account for update - it means it's primary key won't be touched so postgres would not need to aquire the transaction lock so the dead lock can be avoided
	// so instead of just mentioning 'FOR UPDATE' we can do select FOR NO KEY UPDATE
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
