-- +goose Up
CREATE TABLE budget (
                       id int NOT NULL PRIMARY KEY,
                       username text,
                       balance int,
                       updated_at timestamp
);

-- +goose Down
DROP TABLE budget
