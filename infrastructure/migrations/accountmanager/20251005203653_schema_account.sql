-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA account;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA account;
-- +goose StatementEnd
