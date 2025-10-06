-- +goose Up
-- +goose StatementBegin
CREATE TABLE account.accounts
(
    id                 uuid PRIMARY KEY NOT NULL,
    username           text             NOT NULL,
    password_hash      text             NOT NULL,
    email              text             NOT NULL,
    refresh_token_hash text             NOT NULL,
    registration_date  timestamptz      NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account.accounts;
-- +goose StatementEnd
