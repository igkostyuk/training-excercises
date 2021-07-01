-- +migrate Down
DROP TABLE IF EXISTS book cascade;
DROP TABLE IF EXISTS author cascade;
