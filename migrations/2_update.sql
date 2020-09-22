-- +migrate Up
AFTER TABLE users
ADD COLUMN test text;