CREATE TABLE item (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255),
    chrt_id INTEGER,
    track_number VARCHAR(255),
    price INTEGER,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INTEGER,
    size VARCHAR(10),
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR(255),
    status INTEGER
);

CREATE TABLE delivery_info (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    phone VARCHAR(20),
    zip VARCHAR(20),
    city VARCHAR(255),
    address VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255));

CREATE TABLE payment_info (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) UNIQUE,
    transaction VARCHAR(255),
    request_id VARCHAR(255),
    currency VARCHAR(3),
    provider VARCHAR(255),
    amount INTEGER,
    payment_dt VARCHAR(255),
    bank VARCHAR(255),
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) UNIQUE,
    track_number VARCHAR(255),
    entry VARCHAR(255),
    locale VARCHAR(5),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(255),
    sm_id INTEGER,
    date_created TIMESTAMP,
    oof_shard VARCHAR(255)
);