-- +migrate Up
CREATE TABLE album(id text not null, title text not null, artist text not null, price int not null, primary key(id(256)));

-- +migrate Down
DROP TABLE album;
