CREATE TABLE IF NOT EXISTS auctions (
    id SERIAL PRIMARY KEY,
    lot_id INT NOT NULL,
    start_time TIMESTAMP DEFAULT NOW() NOT NULL,
    end_time TIMESTAMP NOT NULL,
    winner_id INT,
    is_closed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    FOREIGN KEY (lot_id) REFERENCES lots (id) ON DELETE CASCADE,
    FOREIGN KEY (winner_id) REFERENCES users (id)
);
