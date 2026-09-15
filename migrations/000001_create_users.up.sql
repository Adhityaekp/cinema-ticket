CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name VARCHAR(100) NOT NULL,
        email VARCHAR(150) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        role VARCHAR(20) NOT NULL DEFAULT 'CUSTOMER' CHECK (role IN ('CUSTOMER', 'ADMIN')),
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE INDEX idx_users_email ON users (email);

CREATE INDEX idx_users_is_active ON users (is_active);