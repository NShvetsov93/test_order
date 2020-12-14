
-- +migrate Up
INSERT INTO products (name) values('one');
INSERT INTO products (name) values('two');
INSERT INTO products (name) values('three');

INSERT INTO products_q (product_id,quantity) values(1,4);
INSERT INTO products_q (product_id,quantity) values(2,1);
INSERT INTO products_q (product_id,quantity) values(3,3);

INSERT INTO orders_limit (quantity) values(0);


-- +migrate Down
