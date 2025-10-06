-- +goose Up
-- +goose StatementBegin
CREATE TABLE account.account_events
(
    id         uuid        NOT NULL PRIMARY KEY,
    account_id uuid        NOT NULL,
    event_type text        NOT NULL,
    payload    jsonb       NOT NULL,
    created_at timestamptz NOT NULL,
    sent       boolean     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account.account_events;
-- +goose StatementEnd
