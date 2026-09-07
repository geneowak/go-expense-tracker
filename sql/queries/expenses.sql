-- name: CreateExpense :one
INSERT INTO
    expenses(
        id,
        item_name,
        category_name,
        quantity,
        unit_cost,
        user_id,
        created_at,
        updated_at
    )
VALUES
    (uuidv7(), $1, $2, $3, $4, $5, NOW(), NOW())
RETURNING
    *;

-- name: GetExpenses :many
SELECT
    *
FROM
    expenses;

-- name: GetExpenseById :one
SELECT
    *
FROM
    expenses
WHERE
    id = $1;

-- name: UpdateExpense :one
UPDATE
    expenses
SET
    item_name = $1,
    category_name = $2,
    quantity = $3,
    unit_cost = $4
WHERE
    id = $5
RETURNING
    *;
