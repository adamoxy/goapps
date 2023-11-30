-- query.sql

-- name: CreateUser :one
-- Creates a new user
INSERT INTO users (name, phone_number) VALUES ($1, $2) RETURNING id;

-- name: GenerateOTP :exec
-- Generates a new OTP for a user
UPDATE users SET otp = $1, otp_expiration_time = $2 WHERE phone_number = $3;

-- name: GetUserOTP :one
-- Retrieves a user's OTP and expiration time
SELECT otp, otp_expiration_time FROM users WHERE phone_number = $1;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1;

