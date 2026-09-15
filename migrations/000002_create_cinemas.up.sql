CREATE TABLE cinemas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name VARCHAR(150) NOT NULL,
    city VARCHAR(100) NOT NULL,
    address TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_cinemas_city ON cinemas(city);