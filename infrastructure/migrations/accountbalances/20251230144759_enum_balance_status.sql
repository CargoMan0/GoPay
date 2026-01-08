-- +goose Up
-- +goose StatementBegin
CREATE TYPE balance.status AS ENUM ('PENDING', 'FAILED', 'SUCCESS');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE balance.status;
-- +goose StatementEnd
