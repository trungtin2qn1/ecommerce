DROP TABLE IF EXISTS buyers;
CREATE TABLE buyers (
    id            BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    name VARCHAR(50),
    bio VARCHAR(100),
    gender INT,
    dob TIMESTAMP,
    membership INT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(email)
);
CREATE INDEX buyers_email ON buyers(email);

DROP TABLE IF EXISTS buyer_phone_numbers;
CREATE TABLE buyer_phone_numbers (
    id            BIGSERIAL PRIMARY KEY,
    id_buyer          BIGINT NOT NULL,
    phone VARCHAR(15),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE INDEX buyer_phone_numbers_id_buyer ON buyer_phone_numbers(id_buyer);

DROP TABLE IF EXISTS buyer_addresses;
CREATE TABLE buyer_addresses (
    id            BIGSERIAL PRIMARY KEY,
    id_buyer          BIGINT NOT NULL,
    country VARCHAR(20),
    province VARCHAR(20),
    city VARCHAR(20),
    district VARCHAR(20),
    ward VARCHAR(20),
    street VARCHAR(20),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE INDEX buyer_addresses_id_buyer ON buyer_addresses(id_buyer);

DROP TABLE IF EXISTS rel_buyer_addresses_phone_numbers;
CREATE TABLE rel_buyer_addresses_phone_numbers (
    id            BIGSERIAL PRIMARY KEY,
    id_phone_number          BIGINT NOT NULL,
    id_address          BIGINT NOT NULL
);
CREATE INDEX rel_buyer_addresses_phone_numbers_id_address ON rel_buyer_addresses_phone_numbers(id_address);
CREATE INDEX rel_buyer_addresses_phone_numbers_id_phone_number ON rel_buyer_addresses_phone_numbers(id_phone_number);

DROP TABLE IF EXISTS sellers;
CREATE TABLE sellers (
    id            BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(email)
);
CREATE INDEX sellers_email ON sellers(email);

DROP TABLE IF EXISTS shops;
CREATE TABLE shops (
    id            BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    country VARCHAR(20),
    province VARCHAR(20),
    city VARCHAR(20),
    district VARCHAR(20),
    ward VARCHAR(20),
    street VARCHAR(20),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

DROP TABLE IF EXISTS rel_shop_sellers;
CREATE TABLE rel_shop_sellers (
    id            BIGSERIAL PRIMARY KEY,
    id_seller          BIGINT NOT NULL,
    id_shop          BIGINT NOT NULL
);
CREATE INDEX rel_shop_sellers_id_seller ON rel_shop_sellers(id_seller);
CREATE INDEX rel_shop_sellers_id_shop ON rel_shop_sellers(id_shop);

DROP TABLE IF EXISTS items;
CREATE TABLE items (
    id            BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    id_shop          BIGINT NOT NULL
);
CREATE INDEX items_id_shop ON items(id_shop);

DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
    id            BIGSERIAL PRIMARY KEY,
    status INT,
    id_buyer         BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE INDEX orders_id_buyer ON orders(id_buyer);

DROP TABLE IF EXISTS rel_order_items;
CREATE TABLE rel_order_items (
    id            BIGSERIAL PRIMARY KEY,
    status INT,
    id_order          BIGINT NOT NULL,
    id_item          BIGINT NOT NULL
);
CREATE INDEX rel_order_items_id_order ON rel_order_items(id_order);
CREATE INDEX rel_order_items_id_item ON rel_order_items(id_item);