
-- +migrate Up
CREATE TABLE IF NOT EXISTS user (
    id int(15) AUTO_INCREMENT,
    user_name varchar(100),
    PRIMARY KEY (id));

-- +migrate Down
DROP TABLE IF EXISTS user;
