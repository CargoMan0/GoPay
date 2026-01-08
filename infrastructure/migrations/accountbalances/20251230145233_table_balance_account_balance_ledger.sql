-- +goose Up
-- +goose StatementBegin
CREATE TABLE balance.account_balance_operation
(
    id                 BIGSERIAL      NOT NULL PRIMARY KEY,
    account_balance_id UUID           NOT NULL REFERENCES balance.account_balance (id),
    type               balance.operation_type,
    amount             DECIMAL(30, 8) NOT NULL,
    related_id         UUID,
    status             balance.status NOT NULL,
    created_at         TIMESTAMPTZ    NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE balance.account_balance_operation;
-- +goose StatementEnd
