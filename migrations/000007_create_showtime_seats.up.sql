CREATE TABLE showtime_seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    showtime_id UUID NOT NULL,
    seat_id UUID NOT NULL,

    status VARCHAR(20) NOT NULL DEFAULT 'AVAILABLE'
        CHECK (
            status IN (
                'AVAILABLE',
                'HELD',
                'SOLD'
            )
        ),

    held_until TIMESTAMP NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_showtime_seats_showtime
        FOREIGN KEY (showtime_id)
        REFERENCES showtimes(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT fk_showtime_seats_seat
        FOREIGN KEY (seat_id)
        REFERENCES seats(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT uq_showtime_seat
        UNIQUE (showtime_id, seat_id)
);

CREATE INDEX idx_showtime_seats_showtime_id
    ON showtime_seats(showtime_id);

CREATE INDEX idx_showtime_seats_status
    ON showtime_seats(showtime_id, status);
