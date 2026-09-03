-- name: CreateUser :one
INSERT INTO
    users(
        id,
        email,
        hashed_password,
        created_at,
        updated_at
    )
VALUES
    (uuidv7(), $1, $2, NOW(), NOW())
RETURNING
    *;
