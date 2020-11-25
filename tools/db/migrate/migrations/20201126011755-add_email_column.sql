
-- +migrate Up
ALTER TABLE user ADD email varchar(100);

-- +migrate Down
ALTER TABLE user DROP COLUMN email;
