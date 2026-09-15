CREATE TABLE
    bookings (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        booking_code VARCHAR(30) NOT NULL UNIQUE,
        user_id UUID NOT NULL,
        showtime_id UUID NOT NULL,
        total_amount NUMERIC(12, 2) NOT NULL CHECK (total_amount >= 0),
        status VARCHAR(30) NOT NULL DEFAULT 'PENDING' CHECK (
            status IN (
                'PENDING',
                'PAID',
                'EXPIRED',
                'CANCELLED',
                'REFUND_PENDING',
                'REFUNDED'
            )
        ),
        expires_at TIMESTAMP NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_bookings_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
        CONSTRAINT fk_bookings_showtime FOREIGN KEY (showtime_id) REFERENCES showtimes (id) ON UPDATE CASCADE ON DELETE RESTRICT
    );

CREATE INDEX idx_bookings_user_id ON bookings (user_id);

CREATE INDEX idx_bookings_showtime_id ON bookings (showtime_id);

CREATE INDEX idx_bookings_status ON bookings (status);

CREATE INDEX idx_bookings_expires_at ON bookings (expires_at);