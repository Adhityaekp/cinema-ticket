ALTER TABLE users
ADD COLUMN email_verified_at TIMESTAMP NULL,
ADD COLUMN email_verification_token VARCHAR(255) NULL,
ADD COLUMN email_verification_expires_at TIMESTAMP NULL;

CREATE INDEX idx_users_email_verification_token
ON users(email_verification_token);