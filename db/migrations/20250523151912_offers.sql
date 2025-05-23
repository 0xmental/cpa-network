-- +goose Up
-- +goose StatementBegin
CREATE TABLE offers (
    id SERIAL PRIMARY KEY,
    target_url TEXT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT FALSE,
    redirect_domain VARCHAR(255) NOT NULL,
    conversion_type SMALLINT NOT NULL,
    payout JSONB NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS offers;
-- +goose StatementEnd
