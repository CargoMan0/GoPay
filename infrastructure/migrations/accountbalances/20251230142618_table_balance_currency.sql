-- +goose Up
-- +goose StatementBegin
CREATE TABLE balance.currency
(
    id         SERIAL      NOT NULL PRIMARY KEY,
    code       VARCHAR(3)  NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE balance.currency;
-- +goose StatementEnd
