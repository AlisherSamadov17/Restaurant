CREATE TABLE users (
    id uuid PRIMARY KEY not null,
    first_name VARCHAR(20) default null,
    last_name VARCHAR(20) default null,
    email VARCHAR(40) not null,
    phone VARCHAR(20)
);

CREATE TABLE orders (
    id uuid PRIMARY KEY not null,
    amount VARCHAR(60),
    user_id uuid REFERENCES users(id),
    created_at TIMESTAMP
);

CREATE TABLE products (
    id uuid PRIMARY KEY not null,
    price INT,
    name VARCHAR(40)
);

CREATE TABLE order_products (
    id uuid PRIMARY KEY not null,
    order_id uuid REFERENCES orders(id),
    product_id uuid REFERENCES products(id),
    quantity INT,
    price INT
);
