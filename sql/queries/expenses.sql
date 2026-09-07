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
