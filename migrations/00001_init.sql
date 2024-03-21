-- +goose Up
-- +goose StatementBegin

CREATE TABLE goods (
    id BIGSERIAL PRIMARY KEY ,
    name TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE goods;

-- +goose StatementEnd
