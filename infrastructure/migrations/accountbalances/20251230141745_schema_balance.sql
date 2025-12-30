-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA balance;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA balance;
-- +goose StatementEnd
