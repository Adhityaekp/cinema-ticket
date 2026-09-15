CREATE TABLE
    refunds (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        booking_id UUID NOT NULL UNIQUE,
        refund_id VARCHAR(150) UNIQUE,
        amount NUMERIC(12, 2) NOT NULL CHECK (amount >= 0),
        reason TEXT NOT NULL,
        status VARCHAR(30) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'SUCCESS', 'FAILED')),
        refunded_at TIMESTAMP NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_refunds_booking FOREIGN KEY (booking_id) REFERENCES bookings (id) ON UPDATE CASCADE ON DELETE RESTRICT
    );

CREATE INDEX idx_refunds_status ON refunds (status);