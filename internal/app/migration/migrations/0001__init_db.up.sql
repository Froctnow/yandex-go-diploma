CREATE SCHEMA gophermart;

CREATE TABLE gophermart.users
(
    id         uuid           DEFAULT gen_random_uuid() PRIMARY KEY,
    login      TEXT UNIQUE              NOT NULL,
    password   TEXT                     NOT NULL,
    balance    NUMERIC(10, 2) DEFAULT 0 NOT NULL,
    created_at TIMESTAMP      DEFAULT NOW()
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
    accrual     NUMERIC(10, 2),
    status      gophermart.order_status NOT NULL DEFAULT 'NEW'::gophermart.order_status,
    uploaded_at TIMESTAMP WITH TIME ZONE         DEFAULT NOW()
);

CREATE TABLE gophermart.transactions
(
    id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    order_number TEXT           NOT NULL,
    user_id      uuid           NOT NULL REFERENCES gophermart.users (id),
    sum          NUMERIC(10, 2) NOT NULL,
    processed_at TIMESTAMP        DEFAULT NOW()
);
