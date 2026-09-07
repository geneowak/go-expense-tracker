-- +goose Up
CREATE TABLE expenses(
    id uuid PRIMARY KEY,
    item_name varchar(255) NOT NULL,
    category_name varchar(255) NOT NULL,
    quantity int NOT NULL,
    unit_cost int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE expenses;
