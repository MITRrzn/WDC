CREATE TABLE events
(
    id         BIGSERIAL PRIMARY KEY,
    type       TEXT        NOT NULL,
    payload    JSONB       NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);