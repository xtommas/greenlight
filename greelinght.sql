CREATE DATABASE greenlight;

-- Create a user called greelight with a password of your choosing
CREATE ROLE greenlight WITH LOGIN PASSWORD 'password';
-- Create extension for case-sensitive character string type
CREATE EXTENSION IF NOT EXISTS citext;