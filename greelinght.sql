CREATE DATABASE greelight;

-- Create a user called greelight with a password of your choosing
CREATE ROLE greelight WITH LOGIN PASSWORD 'password';
-- Create extension for case-sensitive character string type
CREATE EXTENSION IF NOT EXISTS citext;