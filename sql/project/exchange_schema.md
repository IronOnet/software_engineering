## Exchange Schema

- Orders:
    order_id
    side
    product_id
    price
    order_type
    quantity
    created_at
    updated_at


- Product:
    product_id
    name

- OrderBook:
    ob_id
    product_id
    depth
    created_at
    updated_at

- Member:
    member_id
    first_name
    last_name
    email
    password


- Trade:
    - trade_id
    - produc_id
    - order_id



```sql

CREATE DATABASE crypto_db

```

```sql

CREATE TABLE orders
(
  order_id SMALLINT UNSIGNED,
  product_id SMALLINT UNSIGNED,
  price FLOAT,
  quantity FLOAT,
  order_side VARCHAR(5),
  member_id SMALLINT UNSIGNED
  created_at Date,
  updated_at Date,
  constraint pk_orders PRIMARY KEY (order_id),
  constraint fk_product FOREIGN KEY REFERENCES products(product_id),
  constraint fk_member FOREIGN KEY REFERENCES members(member_id),
);

CREATE TABLE products(
  product_id SMALLINT UNSIGNED,
  product_name VARCHAR(20)
  product_short_name VARCHAR(10)
  description TEXT,
  created_at Date,
  updated_at Date,
  CONSTRAINT pk_products PRIMARY KEY (product_id)
);

CREATE TABLE members(
  member_id SMALLINT UNSIGNED,
  first_name VARCHAR(20),
  last_name VARCHAR(20),
  email VARCHAR(20),
  password_hash VARCHAR(30),
  created_at Date,
  updated_at Date,
  CONSTRAINT pk_members PRIMARY KEY (member_id),
);

CREATE TABLE trades(
  trade_id SMALLINT UNSIGNED,
  product_id SMALLINT UNSIGNED,
  order_sell SMALLINT UNSIGNED,
  order_buy SMALLINT UNSIGNED,
  created_at Date,
  updated_at Date
)
