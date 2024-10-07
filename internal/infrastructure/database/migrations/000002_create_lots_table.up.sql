CREATE TABLE IF NOT EXISTS lots (
    id SERIAL PRIMARY KEY,
    seller_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    starting_price NUMERIC(15, 2) NOT NULL,
    min_bid_step NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    FOREIGN KEY (seller_id) REFERENCES users (id) ON DELETE CASCADE
);
