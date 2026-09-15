CREATE TABLE movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    title VARCHAR(200) NOT NULL,
    description TEXT,
    duration_minutes INT NOT NULL CHECK (duration_minutes > 0),
    genre VARCHAR(100),
    age_rating VARCHAR(20),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movies_title ON movies(title);