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
    expenses
WHERE
    (
        sqlc.narg('start_date')::timestamptz IS NULL
        OR created_at >= sqlc.narg('start_date')
    )
    AND (
        sqlc.narg('end_date')::timestamptz IS NULL
        OR created_at <= sqlc.narg('end_date')
    )
ORDER BY
    created_at DESC;

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
    unit_cost = $4,
    updated_at = NOW()
WHERE
    id = $5
RETURNING
    *;

-- name: DeleteExpense :exec
DELETE FROM
    expenses
WHERE
    id = $1;
