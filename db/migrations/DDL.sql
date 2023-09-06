CREATE TABLE customers (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) UNIQUE NOT NULL,
    address VARCHAR(100) NOT NULL
);

CREATE TABLE employees (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) UNIQUE NOT NULL,
    price INT NOT NULL DEFAULT 0,
    unit VARCHAR(10) NOT NULL
);

CREATE TABLE orders (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL,
    order_detail_id UUID NOT NULL,
    entry_date DATE NOT NULL DEFAULT CURRENT_DATE,
    finish_date DATE NOT NULL
);


CREATE TABLE order_details (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    order_id UUID NOT NULL,
    quantity INT NOT NULL,
    subtotal INT NOT NULL DEFAULT 0,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE transactions (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID NOT NULL,
    bill_date DATE NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
)

ALTER TABLE orders ADD CONSTRAINT fk_order_detail FOREIGN KEY (order_detail_id) REFERENCES order_details(id)