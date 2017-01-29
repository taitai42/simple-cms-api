-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS home (
    title varchar(55),
    image_url varchar (255),
    description text
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE home;