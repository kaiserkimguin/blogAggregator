-- +goose Up
CREATE TABLE feeds (
  id UUID,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR,
  url VARCHAR UNIQUE,
  user_id UUID,
  CONSTRAINT fk_feeds_users
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
