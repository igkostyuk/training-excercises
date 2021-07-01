-- Here we will assume that book has only one author

CREATE TABLE IF NOT EXISTS book
(
    id        text NOT NULL PRIMARY KEY,
    title     text DEFAULT '',
    isbn      text DEFAULT '',
    author_id text references author (id)
);

CREATE TABLE IF NOT EXISTS author
(
    id         text NOT NULL PRIMARY KEY,
    first_name text DEFAULT '',
    last_name  text DEFAULT ''
);

