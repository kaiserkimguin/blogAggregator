-- +goose Up
CREATE TABLE posts (
  id UUID PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title VARCHAR NOT NULL,
  url VARCHAR NOT NULL,
  description VARCHAR,
  published_at TIMESTAMP,
  feed_id UUID NOT NULL,
  CONSTRAINT fk_posts_feeds
    FOREIGN KEY (feed_id)
    REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;
