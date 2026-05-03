

CREATE TABLE products (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE reviews (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    rating INT NOT NULL,
    comment TEXT
);