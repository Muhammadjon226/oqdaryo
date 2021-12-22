DROP TABLE IF EXISTS comments;

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    email VARCHAR(128),
    comment TEXT
)