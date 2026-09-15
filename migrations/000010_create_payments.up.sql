CREATE TABLE
    payments (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        booking_id UUID NOT NULL UNIQUE,
        payment_method VARCHAR(50) NOT NULL,
        transaction_id VARCHAR(150) UNIQUE,
        amount NUMERIC(12, 2) NOT NULL CHECK (amount >= 0),
        status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (
            status IN ('PENDING', 'SUCCESS', 'FAILED', 'EXPIRED')
        ),
        paid_at TIMESTAMP NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_payments_booking FOREIGN KEY (booking_id) REFERENCES bookings (id) ON UPDATE CASCADE ON DELETE RESTRICT
    );

CREATE INDEX idx_payments_status ON payments (status);

CREATE INDEX idx_payments_transaction_id ON payments (transaction_id);