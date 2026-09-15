CREATE TABLE seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    studio_id UUID NOT NULL,
    seat_number VARCHAR(10) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_seats_studio
        FOREIGN KEY (studio_id)
        REFERENCES studios(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT uq_seats_studio_number
        UNIQUE (studio_id, seat_number)
);

CREATE INDEX idx_seats_studio_id ON seats(studio_id);