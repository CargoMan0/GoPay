-- +goose Up
-- +goose StatementBegin
CREATE TYPE balance.operation_type AS ENUM ('PAYMENT', 'WITHDRAWAL');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE balance.operation_type;
-- +goose StatementEnd
