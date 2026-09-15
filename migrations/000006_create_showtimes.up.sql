CREATE TABLE showtimes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    movie_id UUID NOT NULL,
    cinema_id UUID NOT NULL,
    studio_id UUID NOT NULL,

    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,

    price NUMERIC(12, 2) NOT NULL CHECK (price >= 0),

    status VARCHAR(20) NOT NULL DEFAULT 'SCHEDULED'
        CHECK (
            status IN (
                'SCHEDULED',
                'ONGOING',
                'COMPLETED',
                'CANCELLED'
            )
        ),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_showtimes_movie
        FOREIGN KEY (movie_id)
        REFERENCES movies(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT fk_showtimes_cinema
        FOREIGN KEY (cinema_id)
        REFERENCES cinemas(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT fk_showtimes_studio
        FOREIGN KEY (studio_id)
        REFERENCES studios(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT chk_showtimes_time
        CHECK (end_time > start_time)
);

CREATE INDEX idx_showtimes_movie_id ON showtimes(movie_id);
CREATE INDEX idx_showtimes_cinema_id ON showtimes(cinema_id);
CREATE INDEX idx_showtimes_studio_id ON showtimes(studio_id);
CREATE INDEX idx_showtimes_start_time ON showtimes(start_time);
CREATE INDEX idx_showtimes_status ON showtimes(status);