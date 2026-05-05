-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR NOT NULL,
  url VARCHAR NOT NULL UNIQUE,
  user_id UUID NOT NULL,
  CONSTRAINT fk_feeds_users
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
