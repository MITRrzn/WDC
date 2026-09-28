CREATE TABLE delivery_attempts
(
    id             BIGSERIAL PRIMARY KEY,
    delivery_id    BIGINT      NOT NULL,
    attempt_number INTEGER     NOT NULL,

    started_at     TIMESTAMPTZ NOT NULL,
    finished_at    TIMESTAMPTZ,

    status_code    INTEGER,
    duration_ms    BIGINT,
    error          TEXT,

    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT delivery_attempts_delivery_id_fk
        FOREIGN KEY (delivery_id)
            REFERENCES deliveries (id),

    CONSTRAINT delivery_attempts_attempt_number_check
        CHECK (attempt_number > 0),

    CONSTRAINT delivery_attempts_status_code_check
        CHECK (
            status_code IS NULL
                OR status_code BETWEEN 100 AND 599
            ),

    CONSTRAINT delivery_attempts_duration_ms_check
        CHECK (
            duration_ms IS NULL
                OR duration_ms >= 0
            )
);

CREATE INDEX idx_delivery_attempts_delivery_id
    ON delivery_attempts (delivery_id);