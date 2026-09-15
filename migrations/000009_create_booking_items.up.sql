CREATE TABLE
    booking_items (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        booking_id UUID NOT NULL,
        showtime_seat_id UUID NOT NULL,
        price NUMERIC(12, 2) NOT NULL CHECK (price >= 0),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_booking_items_booking FOREIGN KEY (booking_id) REFERENCES bookings (id) ON UPDATE CASCADE ON DELETE RESTRICT,
        CONSTRAINT fk_booking_items_showtime_seat FOREIGN KEY (showtime_seat_id) REFERENCES showtime_seats (id) ON UPDATE CASCADE ON DELETE RESTRICT,
        CONSTRAINT uq_booking_showtime_seat UNIQUE (booking_id, showtime_seat_id)
    );

CREATE INDEX idx_booking_items_booking_id ON booking_items (booking_id);

CREATE INDEX idx_booking_items_showtime_seat_id ON booking_items (showtime_seat_id);