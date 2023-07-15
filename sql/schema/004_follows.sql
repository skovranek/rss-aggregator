-- +goose Up
CREATE TABLE follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    feed_id UUID NOT NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds(id),
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE follows;

