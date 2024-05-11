CREATE SCHEMA gophermart;

CREATE TABLE gophermart.users
(
    id         uuid      DEFAULT gen_random_uuid() PRIMARY KEY,
    login      TEXT UNIQUE NOT NULL,
    password   TEXT        NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
