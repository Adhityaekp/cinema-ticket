CREATE TABLE studios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    cinema_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_studios_cinema
        FOREIGN KEY (cinema_id)
        REFERENCES cinemas(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT uq_studios_cinema_name
        UNIQUE (cinema_id, name)
);

CREATE INDEX idx_studios_cinema_id ON studios(cinema_id);