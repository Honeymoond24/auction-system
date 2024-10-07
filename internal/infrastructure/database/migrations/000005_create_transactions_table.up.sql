CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    auction_id INT NOT NULL,
    user_id INT NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL CHECK (transaction_type IN ('replenishment', 'bid_payment', 'refund', 'auction_payout')),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    FOREIGN KEY (auction_id) REFERENCES auctions (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
