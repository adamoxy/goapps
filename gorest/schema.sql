CREATE TABLE users (
    id SERIAL PRIMARY KEY,
  name text    NOT NULL,
  phone_number text    NOT NULL,
  otp text     NULL,
  otp_expiration_time  TIMESTAMP    NULL
);