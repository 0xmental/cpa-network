-- +goose Up
-- +goose StatementBegin
CREATE TABLE partners (
     id SERIAL PRIMARY KEY,
     email VARCHAR(255) UNIQUE NOT NULL,
     pass TEXT NOT NULL,
     contact_info jsonb NOT NULL,
     withdraw_info jsonb,
     postback_url TEXT,
     is_active BOOLEAN DEFAULT TRUE,
     balance NUMERIC(18, 2) DEFAULT 0,
     created_at timestamptz NOT NULL ,
     updated_at timestamptz NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS partners;
-- +goose StatementEnd
