-- +goose Up
-- +goose StatementBegin
CREATE TABLE balance.account_balance
(
    id          UUID           NOT NULL PRIMARY KEY,
    user_id     UUID           NOT NULL,
    currency_id INTEGER        NOT NULL REFERENCES balance.currency (id),
    amount      DECIMAL(30, 8) NOT NULL,
    updated_at  TIMESTAMPTZ    NOT NULL,
    CONSTRAINT UNIQUE (user_id, currency_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE balance.account_balance;
-- +goose StatementEnd
