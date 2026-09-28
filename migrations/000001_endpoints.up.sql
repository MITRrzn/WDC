CREATE TABLE endpoints
(
    id         BIGSERIAL PRIMARY KEY,
    url        TEXT        NOT NULL,
    secret     TEXT        NOT NULL,
    is_active  BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);