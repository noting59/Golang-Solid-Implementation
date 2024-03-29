CREATE schema test

CREATE TABLE test.product (id SERIAL PRIMARY KEY, name VARCHAR, price float);
CREATE TABLE test.cart (id SERIAL PRIMARY KEY, productId INTEGER, userId INTEGER, createdAt TIMESTAMP);

INSERT INTO test.product (id, name, price) VALUES (1, 'Corsair GS600 600 Watt PSU', 120);

CREATE TABLE test.user (id SERIAL PRIMARY KEY, name VARCHAR, email VARCHAR, cardToken VARCHAR default NULL);

CREATE TABLE test.order (id SERIAL PRIMARY KEY, name VARCHAR, productId int, price float, status VARCHAR, userId INTEGER, formToken VARCHAR default NULL, createdAt TIMESTAMP , updatedAt TIMESTAMP);
ALTER SEQUENCE test.order_id_seq RESTART WITH 200;

INSERT INTO test.user (name, email) VALUES ('Vlad', 'v.pistun@gmail.com');