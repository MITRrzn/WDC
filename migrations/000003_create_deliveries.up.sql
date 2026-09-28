CREATE TABLE deliveries
(
    id                    BIGSERIAL PRIMARY KEY,
    event_id              BIGINT      NOT NULL,
    endpoint_id           BIGINT      NOT NULL,

    status                TEXT        NOT NULL DEFAULT 'pending',
    attempts              INTEGER     NOT NULL DEFAULT 0,
    next_attempt_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_error            TEXT,
    processing_started_at TIMESTAMPTZ,

    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    delivered_at          TIMESTAMPTZ,

    CONSTRAINT deliveries_event_id_fk
        FOREIGN KEY (event_id)
            REFERENCES events (id),

    CONSTRAINT deliveries_endpoint_id_fk
        FOREIGN KEY (endpoint_id)
            REFERENCES endpoints (id),

    CONSTRAINT deliveries_status_check
        CHECK (status IN ('pending', 'processing', 'delivered', 'failed')),

    CONSTRAINT deliveries_attempts_check
        CHECK (attempts >= 0)
);

CREATE INDEX idx_deliveries_status_next_attempt_at
    ON deliveries (status, next_attempt_at);

CREATE INDEX idx_deliveries_event_id
    ON deliveries (event_id);

CREATE INDEX idx_deliveries_endpoint_id
    ON deliveries (endpoint_id);