CREATE SCHEMA gophermart;

CREATE TABLE gophermart.users
(
    id         uuid      DEFAULT gen_random_uuid() PRIMARY KEY,
    login      TEXT UNIQUE NOT NULL,
    password   TEXT        NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TYPE gophermart.order_status AS ENUM (
    'NEW',
    'PROCESSING',
    'INVALID',
    'PROCESSED'
    );

CREATE TABLE gophermart.orders
(
    number      TEXT PRIMARY KEY,
    user_id     uuid                    NOT NULL REFERENCES gophermart.users (id),
    accrual     INTEGER,
    status      gophermart.order_status NOT NULL DEFAULT 'NEW'::gophermart.order_status,
    uploaded_at TIMESTAMP WITH TIME ZONE         DEFAULT NOW()
);
