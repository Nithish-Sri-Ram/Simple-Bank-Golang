-- don't mess with the comments - they are ment to say how to generate the files - each word has it's meaning 

-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;


-- This is to block the concurrent accecc of sinle account - when one transaction is going on - because we are using for update account - we could block the access from other account  
-- To avoid deadlock - after doing the required debugging we found that - updating only changes the account balance. The account id will never be changed because it's the primary key of accounts table
-- So if i'm telling that I'm selecting this account for update - it means it's primary key won't be touched so postgres would not need to aquire the transaction lock so the dead lock can be avoided  
-- so instead of just mentioning 'FOR UPDATE' we can do select FOR NO KEY UPDATE  
-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;



-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;


-- Instead of doing get and update function seperately - we are performing one single update function  
-- name: AddAccountBalance :one
UPDATE accounts 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;