-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    create_at DATE,
    update_at DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
