-- create orders table
CREATE TABLE IF NOT EXISTS orders (
    order_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_name VARCHAR(50),
    ordered_at TIMESTAMP
);

-- create items table
CREATE TABLE IF NOT EXISTS items (
    item_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    item_code VARCHAR(20),
    description VARCHAR(100),
    quantity NUMERIC CHECK(quantity > 0),
    order_id INTEGER REFERENCES orders(order_id) ON DELETE CASCADE
)