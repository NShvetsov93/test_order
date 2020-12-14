
-- +migrate Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE
);

CREATE TABLE products_q (
    product_id int UNIQUE,
    quantity int,
    CONSTRAINT fk_products
      FOREIGN KEY(product_id)
	     REFERENCES products(id)
);

CREATE TABLE orders_limit (
    id SERIAL PRIMARY KEY,
    quantity int
);

-- +migrate Down
DROP TABLE IF EXISTS orders_limit;

ALTER TABLE IF EXISTS products_q
DROP CONSTRAINT IF EXISTS fk_products;

DROP TABLE IF EXISTS products_q;
DROP TABLE IF EXISTS products;
