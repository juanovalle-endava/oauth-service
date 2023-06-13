CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   user_id VARCHAR(50) NOT NULL,
   password VARCHAR(100) NOT NULL
);

INSERT INTO users (user_id, password) VALUES
    ('endava', 'secretpass'),
    ('user', 'password');

CREATE TABLE tokens (
   id SERIAL PRIMARY KEY,
   token text NOT NULL
);